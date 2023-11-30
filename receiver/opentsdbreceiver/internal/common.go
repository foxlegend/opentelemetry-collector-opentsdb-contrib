package internal

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type OpenTSDBMetric struct {
	Metric    string            `json:"metric"`
	Timestamp FlexUInt64        `json:"timestamp"`
	Value     json.Number       `json:"value"`
	Tags      map[string]string `json:"tags"`
}

func (o *OpenTSDBMetric) ToOtel() pmetric.Metric {
	md := pmetric.NewMetric()

	metricName := o.Metric
	md.SetName(metricName)
	md.SetEmptyGauge()
	dp := md.Gauge().DataPoints().AppendEmpty()

	// json.Number can be only of type int64 or float64
	if value, err := o.Value.Int64(); err == nil {
		dp.SetIntValue(value)
	} else if value, err := o.Value.Float64(); err == nil {
		dp.SetDoubleValue(value)
	}

	
	var ts pcommon.Timestamp
	if IsTimestampSeconds(o.Timestamp) {
		ts = pcommon.Timestamp(o.Timestamp * 1000000000)
	}  else {
		ts = pcommon.Timestamp(o.Timestamp * 1000000)
	}
	dp.SetTimestamp(ts)
	dp.SetStartTimestamp(ts)

	for key, value := range o.Tags {
		dp.Attributes().PutStr(key, value)
	}

	return md
}

func IsTimestampSeconds(timestamp FlexUInt64) bool {
	if t:=time.UnixMilli(int64(timestamp)); t.Year() > 1971 {
		return false
	}
	return true
}

// FlexInt64 and FlexFloat64 trick comes from
// https://engineering.bitnami.com/articles/dealing-with-json-with-non-homogeneous-types-in-go.html
// and https://github.com/gildas/go-core/blob/master/flexint.go

type FlexUInt64 uint64

// UnmarshalJSON decodes JSON
//   implements json.Unmarshaler interface
// from: https://github.com/gildas/go-core/blob/master/flexint.go
func (i *FlexUInt64) UnmarshalJSON(payload []byte) error {
	unquoted := strings.Replace(string(payload), `"`, ``, -1)
	value, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}
	*i = FlexUInt64(value)
	return nil
}
