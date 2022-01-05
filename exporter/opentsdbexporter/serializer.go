package opentsdbexporter

import (
	"bytes"
	"fmt"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

type Metric struct {
	Metric    string            `json:"metric"`
	Timestamp uint64            `json:"timestamp"`
	Value     interface{}       `json:"value"`
	Tags      map[string]string `json:"tags"`
}

func (m Metric) String() string {
	tags := new(bytes.Buffer)
	for key, value := range m.Tags {
		fmt.Fprintf(tags, "%s='%s' ", key, value)
	}
	return fmt.Sprintf("[%d] %s = %s { %s }", m.Timestamp, m.Metric, m.Value, tags)
}

type HttpSerializer struct {
	logger   *zap.Logger
	maxTags  int
	skipTags []string
}

func NewHttpSerializer(logger *zap.Logger, maxTags int, skipTags []string) *HttpSerializer {
	return &HttpSerializer{
		logger:   logger,
		maxTags:  maxTags,
		skipTags: skipTags,
	}
}

func (h HttpSerializer) Marshal(metrics pdata.Metrics) (sMetrics []*Metric, errs []error) {
	h.logger.Debug("HttpSerializer#Marshal", zap.Int("#metrics", metrics.MetricCount()), zap.Int("#datapoints", metrics.DataPointCount()))

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
				case pdata.MetricDataTypeGauge:
					s, sErrs := h.serializeGauge(metric, resource, il)
					if sErrs != nil {
						errs = append(errs, sErrs...)
					}
					sMetrics = append(sMetrics, s...)
				case pdata.MetricDataTypeSum:
					s, sErrs := h.serializeSum(metric, resource, il)
					if sErrs != nil {
						errs = append(errs, sErrs...)
					}
					sMetrics = append(sMetrics, s...)
				default:
					errs = append(errs, fmt.Errorf("unhandled DataType: %s", metric.DataType()))
				}

			}
		}
	}

	return sMetrics, errs
}

func (h *HttpSerializer) serializeGauge(metric pdata.Metric, resource pdata.Resource, instrumentationLibrary pdata.InstrumentationLibrary) (mSlice []*Metric, errs []error) {
	dps := metric.Gauge().DataPoints()
	for i := 0; i < dps.Len(); i++ {
		dp := dps.At(i)

		var value interface{}
		switch dp.Type() {
		case pdata.MetricValueTypeNone:
			continue
		case pdata.MetricValueTypeInt:
			value = dp.IntVal()
		case pdata.MetricValueTypeDouble:
			value = dp.DoubleVal()
		default:
			errs = append(errs, fmt.Errorf("unsupported gauge data point type %d", dp.Type()))
		}

		tags, _ := h.createTags(resource, instrumentationLibrary, dp.Attributes(), metric.Name())

		m := &Metric{
			Metric:    metric.Name(),
			Timestamp: uint64(dp.Timestamp()) / 1000000,
			Value:     value,
			Tags:      tags,
		}
		mSlice = append(mSlice, m)
	}
	return mSlice, errs
}

func (h *HttpSerializer) serializeSum(metric pdata.Metric, resource pdata.Resource, instrumentationLibrary pdata.InstrumentationLibrary) (mSlice []*Metric, errs []error) {
	if metric.Sum().AggregationTemporality() != pdata.MetricAggregationTemporalityCumulative {
		return nil, append(errs, fmt.Errorf("unsupported sum aggregation temporality %q", metric.Sum().AggregationTemporality()))
	}
	if !metric.Sum().IsMonotonic() {
		return nil, append(errs, fmt.Errorf("unsupported non-monotonic sum '%s'", metric.Name()))
	}

	dps := metric.Sum().DataPoints()
	for i := 0; i < dps.Len(); i++ {
		dp := dps.At(i)

		var value interface{}
		switch dp.Type() {
		case pdata.MetricValueTypeNone:
			continue
		case pdata.MetricValueTypeInt:
			value = dp.IntVal()
		case pdata.MetricValueTypeDouble:
			value = dp.DoubleVal()
		default:
			errs = append(errs, fmt.Errorf("unsupported sum data point type %d", dp.Type()))
		}

		tags, _ := h.createTags(resource, instrumentationLibrary, dp.Attributes(), metric.Name())

		m := &Metric{
			Metric:    metric.Name(),
			Timestamp: uint64(dp.Timestamp()) / 1000000,
			Value:     value,
			Tags:      tags,
		}
		mSlice = append(mSlice, m)
	}
	return mSlice, errs
}

func (h *HttpSerializer) createTags(resource pdata.Resource, instrumentationLibrary pdata.InstrumentationLibrary, attributes pdata.AttributeMap, metricName string) (map[string]string, map[string]string) {
	tags := make(map[string]string)
	skipped := make(map[string]string)

	attributes.Range(func(key string, value pdata.AttributeValue) bool {
		if key != "" {
			if h.shouldIncludeTag(key) {
				if len(tags) < h.maxTags {
					// tags[key] = strings.Replace(value.AsString(), ":", "/", -1)
					tags[key] = sanitizeForOpenTSDB(value.AsString())
				} else {
					skipped[key] = sanitizeForOpenTSDB(value.AsString())
				}
			}
		}
		return true
	})

	tags, skipped = h.resourceToTags(resource, tags, skipped)
	tags, skipped = h.instrumentationLibraryToTags(instrumentationLibrary, tags, skipped)

	if skipped != nil && len(skipped) > 0 {
		h.logger.Warn("tags skipped during serialization", zap.String("#name", metricName), zap.Int("#skippedCount", len(skipped)))
		if ce := h.logger.Check(zap.DebugLevel, "skipped tags"); ce != nil {
			buffer := new(bytes.Buffer)
			for key, value := range skipped {
				fmt.Fprintf(buffer, "%s=\"%s\" ", key, value)
			}
			ce.Write(zap.String("tags", strings.Trim(buffer.String(), " ")))
		}
	}

	return tags, skipped
}

func (h *HttpSerializer) resourceToTags(resource pdata.Resource, tags map[string]string, skipped map[string]string) (map[string]string, map[string]string) {
	resource.Attributes().Range(func(key string, value pdata.AttributeValue) bool {
		if key != "" {
			if h.shouldIncludeTag(key) {
				if len(tags) < h.maxTags {
					// tags[key] = strings.Replace(value.AsString(), ":", "/", -1)
					tags[key] = sanitizeForOpenTSDB(value.AsString())
				} else {
					skipped[key] = sanitizeForOpenTSDB(value.AsString())
				}
			}
		}
		return true
	})
	return tags, skipped
}

func (h *HttpSerializer) instrumentationLibraryToTags(instrumentationLibrary pdata.InstrumentationLibrary, tags map[string]string, skipped map[string]string) (map[string]string, map[string]string) {
	if instrumentationLibrary.Name() != "" {
		if len(tags) < h.maxTags {
			tags["otel.library.name"] = sanitizeForOpenTSDB(instrumentationLibrary.Name())
		} else {
			skipped["otel.library.name"] = sanitizeForOpenTSDB(instrumentationLibrary.Name())
		}
	}
	if instrumentationLibrary.Version() != "" {
		if len(tags) < h.maxTags {
			tags["otel.library.version"] = sanitizeForOpenTSDB(instrumentationLibrary.Version())
		} else {
			skipped["otel.library.version"] = sanitizeForOpenTSDB(instrumentationLibrary.Version())
		}
	}
	return tags, skipped
}

func (h *HttpSerializer) shouldIncludeTag(tag string) bool {
	for i := 0; i < len(h.skipTags); i++ {
		if strings.ToLower(tag) == strings.ToLower(h.skipTags[i]) {
			return false
		}
	}
	return true
}

func sanitizeForOpenTSDB(value string) string {
	if value == "" {
		return value
	}

	// Try to keep some meanings around bucket limits
	if value == "+Inf" {
		return "pInf"
	}
	if value == "-Inf" {
		return "mInf"
	}

	reg, _ := regexp.Compile(`[^a-zA-Z0-9\-._/]+`)
	return reg.ReplaceAllString(value, "_")
}