package serialization

import (
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
	"strings"
)

type Serializer interface {
	Marshal() ([]byte, error)
	SetMetrics(metrics pdata.Metrics)
}

type Metric struct {
	Metric    string            `json:"metric"`
	Timestamp uint64            `json:"timestamp"`
	Value     interface{}       `json:"value"`
	Tags      map[string]string `json:"tags"`
}

type HttpSerializer struct {
	logger *zap.Logger
	init   bool
}

func NewHttpSerializer(logger *zap.Logger) *HttpSerializer {
	return &HttpSerializer{
		logger: logger,
		init:   false,
	}
}

func (h HttpSerializer) Marshal(metrics pdata.Metrics) ([]*Metric, error) {
	h.logger.Debug("HttpSerializer#Marshal", zap.Int("#metrics", metrics.MetricCount()), zap.Int("#datapoints", metrics.DataPointCount()))
	var sMetrics []*Metric

	rms := metrics.ResourceMetrics()
	for i := 0; i < rms.Len(); i++ {
		h.logger.Debug("ResourceMetrics", zap.Int("#id", i))
		rm := rms.At(i)
		resource := rm.Resource()

		ilms := rm.InstrumentationLibraryMetrics()
		for j := 0; j < ilms.Len(); j++ {
			h.logger.Debug("InstrumentationLibraryMetric", zap.Int("#id", j))
			ilm := ilms.At(j)

			il := ilm.InstrumentationLibrary()
			h.logger.Debug("InstrumentationLibrary", zap.String("#name", il.Name()), zap.String("#version", il.Version()))

			ms := ilm.Metrics()
			for k := 0; k < ms.Len(); k++ {
				h.logger.Debug("Metric", zap.Int("#id", k))
				metric := ms.At(k)

				switch metric.DataType() {
				case pdata.MetricDataTypeNone:
					h.logger.Debug("    -> DataType: None")
				case pdata.MetricDataTypeGauge:
					h.logger.Debug("    -> DataType: Gauge")
					s, err := serializeGauge(h.logger, metric, resource, il)
					if err != nil {
						h.logger.Sugar().Errorf("failed to serialize %s %s: %s", metric.DataType().String(), metric.Name(), err.Error())
					} else {
						sMetrics = append(sMetrics, s...)
					}
				default:
					h.logger.Sugar().Errorf("unhandled DataType: %s", metric.DataType())
					// return nil, fmt.Errorf("unhandled DataType: %s", metric.DataType())
				}

			}
		}
	}

	return sMetrics, nil
	//return json.Marshal(sMetrics)
}

func resourceToTags(resource pdata.Resource, tags map[string]string) map[string]string {
	resource.Attributes().Range(func(key string, value pdata.AttributeValue) bool {
		tags[key] = strings.Replace(value.AsString(), ":", "_", -1)
		return true
	})
	return tags
}

func instrumentationLibraryToTags(instrumentationLibrary pdata.InstrumentationLibrary, tags map[string]string) map[string]string {
	if instrumentationLibrary.Name() != "" {
		tags["otel.library.name"] = instrumentationLibrary.Name()
	}
	if instrumentationLibrary.Version() != "" {
		tags["otel.library.version"] = instrumentationLibrary.Version()
	}
	return tags
}

func createTags(resource pdata.Resource, instrumentationLibrary pdata.InstrumentationLibrary, attributes pdata.AttributeMap) map[string]string {
	tags := make(map[string]string)

	attributes.Range(func(key string, value pdata.AttributeValue) bool {
		if key != "" {
			tags[key] = strings.Replace(value.AsString(), ":", "_", -1)
		}
		return true
	})

	tags = resourceToTags(resource, tags)
	tags = instrumentationLibraryToTags(instrumentationLibrary, tags)

	return tags
}
