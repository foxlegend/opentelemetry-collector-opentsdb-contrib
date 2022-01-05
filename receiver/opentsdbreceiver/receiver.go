package opentsdbreceiver

import (
	"context"
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
	"io"
	"net/http"
	"runtime"
	"strings"
	"sync"

	conventions "go.opentelemetry.io/collector/model/semconv/v1.5.0"
)

type metricsReceiver struct {
	logger             *zap.Logger
	nextConsumer       consumer.Metrics
	httpServerSettings *confighttp.HTTPServerSettings

	server *http.Server
	wg     sync.WaitGroup

	settings component.TelemetrySettings
}

func newMetricsReceiver(config *Config, logger *zap.Logger, settings component.TelemetrySettings, nextConsumer consumer.Metrics) (*metricsReceiver, error) {
	receiver := &metricsReceiver{
		logger:             logger,
		nextConsumer:       nextConsumer,
		httpServerSettings: &config.HTTPServerSettings,
		settings:           settings,
	}
	return receiver, nil
}

func (r *metricsReceiver) Start(_ context.Context, host component.Host) error {
	ln, err := r.httpServerSettings.ToListener()
	if err != nil {
		return fmt.Errorf("failed to bind to address %s: %w", r.httpServerSettings.Endpoint, err)
	}

	router := http.NewServeMux()
	router.HandleFunc("/api/put", r.handleWrite)
	router.HandleFunc("/", r.unhandledRoute)

	r.wg.Add(1)
	r.server, err = r.httpServerSettings.ToServer(host, r.settings, router)
	if err != nil {
		return fmt.Errorf("failed to instantiate HTTP server: %w", err)
	}
	go func() {
		defer r.wg.Done()
		if err := r.server.Serve(ln); err != nil && err != http.ErrServerClosed {
			host.ReportFatalError(err)
		}
	}()
	return nil
}

func (r *metricsReceiver) Shutdown(ctx context.Context) error {
	if err := r.server.Close(); err != nil {
		return err
	}
	r.wg.Wait()
	return nil
}

func (r *metricsReceiver) handleWrite(w http.ResponseWriter, req *http.Request) {
	defer func() {
		_ = req.Body.Close()
	}()
	w.Header().Add("Content-Type", "application/json")

	r.logger.Debug("Request received", zap.String("Method", req.Method))

	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Add("Content-Type", "application/json")
		trace := make([]byte, 4096)
		_ = runtime.Stack(trace, true)
		buff, _ := json.Marshal(
			httpResponse{
				Error: &opentsdbErrorResponse{
					Code:    http.StatusMethodNotAllowed,
					Details: fmt.Sprintf("The HTTP method [%s] is not permitted for this endpoint", req.Method),
					Message: "Method not allowed",
					Trace:   string(trace),
				},
			})
		_, _ = w.Write(buff)
		return
	}

	md := pdata.NewMetrics()
	rs := md.ResourceMetrics().AppendEmpty()
	rs.SetSchemaUrl(conventions.SchemaURL)
	ils := rs.InstrumentationLibraryMetrics().AppendEmpty()

	dec := json.NewDecoder(req.Body)
	t, err := dec.Token()
	if err != nil {
		// TODO: error
		r.logger.Error(fmt.Sprintf("Unable to read token: %s", err))
	}

	delim, ok := t.(json.Delim)
	if !ok {
		r.logger.Error(fmt.Sprintf("Expected an object or an array"))
	}

	if delim == '{' {
		singleDec := json.NewDecoder(io.MultiReader(strings.NewReader("{"), dec.Buffered(), req.Body))
		metric := opentsdbMetric{}
		err = singleDec.Decode(&metric)
		if err != nil {
			r.logger.Error(fmt.Sprintf("Unable to decode Metric: %s", err))
		}

		r.logger.Info("Fetched Metric!", zap.String("metric", metric.Metric), zap.Uint64("timestamp", metric.Timestamp))
		appendMetric(ils.Metrics(), &metric)
	} else if delim == '[' {
		for dec.More() {
			metric := opentsdbMetric{}
			err = dec.Decode(&metric)
			if err != nil {
				r.logger.Error(fmt.Sprintf("Unable to decode Metric: %s", err))
			}
			r.logger.Info("Fetched array Metric!", zap.String("metric", metric.Metric), zap.Uint64("timestamp", metric.Timestamp))
			appendMetric(ils.Metrics(), &metric)
		}
	}

	if err := r.nextConsumer.ConsumeMetrics(req.Context(), md); err != nil {
		if consumererror.IsPermanent(err) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		r.logger.Debug(fmt.Sprintf("failed to pass metrics to next consumer: %s", err))
	}

	w.WriteHeader(http.StatusNoContent)
}

func (r *metricsReceiver) unhandledRoute(w http.ResponseWriter, req *http.Request) {
	defer func() {
		_ = req.Body.Close()
	}()

	r.logger.Info("Unhandled API route", zap.String("Method", req.Method), zap.String("URL", req.RequestURI))
	w.WriteHeader(http.StatusNotFound)
	w.Header().Add("Content-Type", "application/json")
	buff, _ := json.Marshal(
		httpResponse{&opentsdbErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Endpoint not found",
		}})
	_, _ = w.Write(buff)
}

type opentsdbMetric struct {
	Metric    string            `json:"metric"`
	Timestamp uint64            `json:"timestamp"`
	Value     interface{}       `json:"value"`
	Tags      map[string]string `json:"tags"`
}

type httpResponse struct {
	Error *opentsdbErrorResponse `json:"error,omitempty"`
}

type opentsdbErrorResponse struct {
	Code    uint16 `json:"code"`
	Details string `json:"details,omitempty"`
	Message string `json:"message,omitempty"`
	Trace   string `json:"trace,omitempty"`
}

func appendMetric(dest pdata.MetricSlice, metric *opentsdbMetric) {
	md := dest.AppendEmpty()
	metricName := metric.Metric
	md.SetName(metricName)
	md.SetDataType(pdata.MetricDataTypeGauge)
	dp := md.Gauge().DataPoints().AppendEmpty()

	switch result := metric.Value.(type) {
	case float64:
		dp.SetDoubleVal(result)
	case int64:
		dp.SetIntVal(result)
	default:
		fmt.Printf("Unhandeld type %s, %t", result, result)
	}
	ts := pdata.Timestamp(metric.Timestamp * 1000000000)
	dp.SetTimestamp(ts)
	dp.SetStartTimestamp(ts)

	for key, value := range metric.Tags {
		dp.Attributes().InsertString(key, value)
	}
}
