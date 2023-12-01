package internal

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"strings"
)

type JsonSerializer struct {
	logger *zap.Logger
}

func NewJsonSerializer(logger *zap.Logger) *JsonSerializer {
	return &JsonSerializer{
		logger: logger,
	}
}

func (j *JsonSerializer) Serialize(input io.ReadCloser) (metrics []*OpenTSDBMetric, errs []error) {
	decoder := json.NewDecoder(input)

	token, err := decoder.Token()
	if err != nil {
		j.logger.Error(fmt.Sprintf("Unable to read token: %s", err))
	}

	delim, ok := token.(json.Delim)
	if !ok {
		err := errors.Errorf("Expected an object or an array had '%s'", delim)
		errs = append(errs, err)
		j.logger.Error(err.Error())
		return nil, errs
	}

	if delim == '{' {
		// Construct a new decoder with stream reset with initial token
		newDecoder := json.NewDecoder(io.MultiReader(strings.NewReader("{"), decoder.Buffered(), input))
		metrics, errs = j.serializeSingle(newDecoder)
	} else if delim == '[' {
		metrics, errs = j.serializeMultiple(decoder)
	} else {
		err := errors.Errorf("Unexpected delimiter '%s', expected '{' (object) or '[' (array)", delim)
		errs = append(errs, err)
		j.logger.Error(err.Error())
		return nil, errs
	}

	return metrics, errs
}

func (j *JsonSerializer) serializeSingle(decoder *json.Decoder) (metrics []*OpenTSDBMetric, errs []error) {
	metric, err := j.serializeMetric(decoder)
	if err != nil {
		errs = append(errs, err)
	}
	if metric != nil {
		metrics = append(metrics, metric)
	}
	return metrics, errs
}

func (j *JsonSerializer) serializeMultiple(decoder *json.Decoder) (metrics []*OpenTSDBMetric, errs []error) {
	for decoder.More() {
		metric, err := j.serializeMetric(decoder)
		if err != nil {
			errs = append(errs, err)
		}
		if metric != nil {
			metrics = append(metrics, metric)
		}
	}
	return metrics, errs
}

func (j *JsonSerializer) serializeMetric(decoder *json.Decoder) (*OpenTSDBMetric, error) {
	metric := OpenTSDBMetric{}
	if err := decoder.Decode(&metric); err != nil {
		j.logger.Warn(fmt.Sprintf("Unable to decode Metric: %s", err))
		return nil, err
	}
	return &metric, nil
}
