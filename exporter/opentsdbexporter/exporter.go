package opentsdbexporter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type OpenTSDBExporter struct {
	logger     *zap.Logger
	cfg        *Config
	client     *http.Client
	serializer *HttpSerializer
}

func (e *OpenTSDBExporter) PushMetrics(ctx context.Context, md pdata.Metrics) error {
	e.logger.Info("MetricsExporter", zap.Int("#metrics", md.MetricCount()), zap.Int("#datapoints", md.DataPointCount()))
	buf, err := e.serializer.Marshal(md)
	e.logger.Debug("serialization results", zap.Int("#serialized", len(buf)), zap.Int("#errors", md.DataPointCount()-len(buf)))
	if err != nil {
		if ce := e.logger.Check(zap.DebugLevel, "serialization errors"); ce != nil {
			for i := 0; i < len(err); i++ {
				ce.Write(zap.String(fmt.Sprintf("%d", i), err[i].Error()))
			}
		}
	}

	for i := 0; i < len(buf); i += e.cfg.BatchSize {
		end := i + e.cfg.BatchSize

		if end > len(buf) {
			end = len(buf)
		}
		j, err := json.Marshal(buf[i:end])
		if err != nil {
			e.logger.Sugar().Errorf("Error serializing: %s", err.Error())
		}
		if err := e.send(ctx, j); err != nil {
			return err
		}
	}
	return nil
}

func NewOpenTSDBExporter(config *Config, logger *zap.Logger) *OpenTSDBExporter {
	return &OpenTSDBExporter{
		cfg:        config,
		logger:     logger,
		serializer: NewHttpSerializer(logger, config.MaxTags, config.SkipTags),
	}
}

func NewMetricsExporter(config config.Exporter, logger *zap.Logger, set component.ExporterCreateSettings) (component.MetricsExporter, error) {
	cfg := config.(*Config)
	t := NewOpenTSDBExporter(cfg, logger)
	return exporterhelper.NewMetricsExporter(
		config,
		set,
		t.PushMetrics,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
		exporterhelper.WithRetry(exporterhelper.RetrySettings{Enabled: false}),
		exporterhelper.WithQueue(exporterhelper.QueueSettings{Enabled: false}),
		exporterhelper.WithStart(t.start),
	)
}

func (e *OpenTSDBExporter) start(_ context.Context, host component.Host) (err error) {
	u, err := url.Parse(e.cfg.Endpoint)
	if err != nil {
		return err
	}
	q := u.Query()
	// Add details for better error handling
	q.Set("details", "true")
	u.RawQuery = q.Encode()
	e.cfg.Endpoint = u.String()

	client, err := e.cfg.HTTPClientSettings.ToClient(host.GetExtensions())
	if err != nil {
		return err
	}

	e.client = client
	return nil
}

func (e *OpenTSDBExporter) send(ctx context.Context, buffer []byte) error {
	req, err := http.NewRequestWithContext(ctx, "POST", e.cfg.Endpoint, bytes.NewReader(buffer))
	if err != nil {
		return consumererror.NewPermanent(err)
	}

	e.logger.Sugar().Debugf("Sending Request (%d bytes)", len(buffer))
	resp, err := e.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	e.logger.Debug("Response", zap.Int("#statuscode", resp.StatusCode), zap.String("#status", resp.Status))
	if resp.StatusCode == http.StatusBadRequest {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// if the response cannot be read, do not retry the batch as it may have been successful
			e.logger.Error(fmt.Sprintf("failed to read response: %s", err.Error()))
			return nil
		}

		responseBody := metricsResponse{}
		if err := json.Unmarshal(bodyBytes, &responseBody); err != nil {
			if strings.Contains(strings.ToLower(string(bodyBytes)), "chunked request not supported.") {
				e.logger.Sugar().Errorf("Request too large (%d bytes). OpenTSDB does not support chunked request. Please decrease batch size.", len(buffer))
			} else {
				e.logger.Sugar().Errorf("failed to unmarshal response: %s (%s)", bodyBytes, err.Error())
			}
		}

		e.logger.Info("Ingestion status", zap.Int("#success", responseBody.Ok), zap.Int("#failed", responseBody.Invalid))
		for i := 0; i < len(responseBody.Errors); i++ {
			e.logger.Debug("Ingestion error", zap.String("#metric", responseBody.Errors[i].Metric.String()), zap.String("#Error", responseBody.Errors[i].Error))
		}
	}
	return nil
}

type metricsResponse struct {
	Ok      int                    `json:"success"`
	Invalid int                    `json:"failed"`
	Errors  []metricsResponseError `json:"errors"`
}

type metricsResponseError struct {
	Metric Metric `json:"datapoint"`
	Error  string `json:"error"`
}
