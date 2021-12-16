package serialization

import (
	"fmt"
	"go.opentelemetry.io/collector/model/pdata"
	"go.uber.org/zap"
)

func serializeGauge(logger *zap.Logger, metric pdata.Metric, resource pdata.Resource, instrumentationLibrary pdata.InstrumentationLibrary) (mSlice []*Metric, err error) {
	dps := metric.Gauge().DataPoints()
	for i := 0; i < dps.Len(); i++ {
		dp := dps.At(i)

		var value interface{}
		switch dp.Type() {
		case pdata.MetricValueTypeNone:
			return nil, fmt.Errorf("unsupported value type non")
		case pdata.MetricValueTypeInt:
			value = dp.IntVal()
		case pdata.MetricValueTypeDouble:
			value = dp.DoubleVal()
		}

		labels := createTags(resource, instrumentationLibrary, dp.Attributes())

		m := &Metric{
			Metric:    metric.Name(),
			Timestamp: uint64(dp.Timestamp()) / 1000000,
			Value:     value,
			Tags:      labels,
		}
		mSlice = append(mSlice, m)
	}
	return mSlice, nil
}
