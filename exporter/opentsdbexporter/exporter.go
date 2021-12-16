package opentsdbexporter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/foxlegend/opentelemetry-collector-opentsdb-contrib/exporter/opentsdbexporter/serialization"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type OpenTSDBExporter struct {
	logger     *zap.Logger
	cfg        *Config
	client     *http.Client
	serializer *serialization.HttpSerializer
}

func (e *OpenTSDBExporter) PushMetrics(ctx context.Context, md pdata.Metrics) error {
	e.logger.Info("MetricsExporter", zap.Int("#metrics", md.MetricCount()), zap.Int("#datapoints", md.DataPointCount()))
	buf, err := e.serializer.Marshal(md)
	if err != nil {
		return err
	}

	for i := 0; i < len(buf); i += 10 {
		end := i + 10

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
		serializer: serialization.NewHttpSerializer(logger),
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
		exporterhelper.WithShutdown(loggerSync(logger)),
		exporterhelper.WithStart(t.start),
	)
}

func loggerSync(logger *zap.Logger) func(context.Context) error {
	return func(context.Context) error {
		err := logger.Sync()
		return err
	}
}

func (e *OpenTSDBExporter) start(_ context.Context, host component.Host) (err error) {
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

	e.logger.Info("Sending Request")
	resp, err := e.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	e.logger.Info("Response", zap.Int("#statuscode", resp.StatusCode), zap.String("#status", resp.Status))

	// At least some metrics were not accepted
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// if the response cannot be read, do not retry the batch as it may have been successful
		e.logger.Error(fmt.Sprintf("failed to read response: %s", err.Error()))
		return nil
	}

	e.logger.Info("Body", zap.String("#bodxy", string(bodyBytes)))
	return nil
}
