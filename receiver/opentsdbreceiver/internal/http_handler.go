package internal

import (
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/pdata/pmetric"
	conventions "go.opentelemetry.io/collector/model/semconv/v1.5.0"
	"go.uber.org/zap"
	"net/http"
	"runtime"
)

const fallbackContentType = "application/json"

type HttpHandler struct {
	logger *zap.Logger

	serializer      *JsonSerializer
	metricsConsumer consumer.Metrics
}

func NewHttpHandler(logger *zap.Logger, consumer consumer.Metrics) *HttpHandler {
	return &HttpHandler{
		logger:          logger,
		serializer:      NewJsonSerializer(logger),
		metricsConsumer: consumer,
	}
}

func (h *HttpHandler) NewHttpRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/api/put", h.HandleWrite)
	return router
}

func (h *HttpHandler) HandleWrite(w http.ResponseWriter, req *http.Request) {
	defer func() {
		_ = req.Body.Close()
	}()
	w.Header().Add("Content-Type", fallbackContentType)

	h.logger.Debug("Request received", zap.String("Method", req.Method), zap.String("URI", req.RequestURI))

	if req.Method != "POST" {
		h.unhandledHttpMethod(w, req)
		return
	}

	opentsdbMetrics, serializationErrs := h.serializer.Serialize(req.Body)

	ms := pmetric.NewMetricSlice()
	if opentsdbMetrics != nil {
		for _, m := range opentsdbMetrics {
			mp := ms.AppendEmpty()
			m.ToOtel().CopyTo(mp)
		}
	}

	md := pmetric.NewMetrics()
	rs := md.ResourceMetrics().AppendEmpty()
	rs.SetSchemaUrl(conventions.SchemaURL)
	ils := rs.ScopeMetrics().AppendEmpty()
	ms.CopyTo(ils.Metrics())

	if err := h.metricsConsumer.ConsumeMetrics(req.Context(), md); err != nil {
		if consumererror.IsPermanent(err) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		h.logger.Debug(fmt.Sprintf("failed to pass metrics to next consumer: %s", err))
	}

	h.writeDetails(w, len(opentsdbMetrics), len(serializationErrs), serializationErrs)
	w.WriteHeader(http.StatusNoContent)
}

func (h *HttpHandler) unhandledHttpMethod(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
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
}

func (h *HttpHandler) writeDetails(w http.ResponseWriter, successCount int, errorsCount int, errors []error) {

	dpErrors := make([]*opentsdbDataPointError, len(errors))
	for i := 0; i < len(errors); i++ {
		dpErrors = append(dpErrors, &opentsdbDataPointError{
			Error: errors[i].Error(),
		})
	}
	buff, _ := json.Marshal(
		httpResponse{
			Success: successCount,
			Failed:  errorsCount,
			Errors:  dpErrors,
		},
	)
	_, _ = w.Write(buff)
}

type httpResponse struct {
	Failed  int                       `json:"failed"`
	Success int                       `json:"success"`
	Error   *opentsdbErrorResponse    `json:"error,omitempty"`
	Errors  []*opentsdbDataPointError `json:"errors,omitempty"`
}

type opentsdbErrorResponse struct {
	Code    uint16 `json:"code"`
	Details string `json:"details,omitempty"`
	Message string `json:"message,omitempty"`
	Trace   string `json:"trace,omitempty"`
}

type opentsdbDataPointError struct {
	Error string `json:"error"`
}
