package opentsdbexporter

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry"
	"github.com/pkg/errors"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"strings"
)

type Config struct {
	exporterhelper.TimeoutSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.

	// OpenTSDB Endpoint
	confighttp.HTTPClientSettings `mapstructure:",squash"`

	// The maximum number of datapoints to send to the OpenTSDB backend
	BatchSize int `mapstructure:"batch_size"`
	// The maximum number of tags per datapoint
	MaxTags int `mapstructure:"max_tags"`
	// Tags to skip
	SkipTags []string `mapstructure:"skip_tags"`

	// ResourceToTelemetrySettings is the option for converting resource attributes to telemetry attributes.
	// "Enabled" - A boolean field to enable/disable this option. Default is `false`.
	// If enabled, all the resource attributes will be converted to metric labels by default.
	ResourceToTelemetrySettings resourcetotelemetry.Settings `mapstructure:"resource_to_telemetry_conversion"`
}

var _ component.Config = (*Config)(nil)

func (cfg *Config) Validate() error {
	if !(strings.HasPrefix(cfg.Endpoint, "http://") || strings.HasPrefix(cfg.Endpoint, "https")) {
		return errors.New("endpoint must start with https:// or http://")
	}
	if !(strings.HasSuffix(cfg.Endpoint, "/api/put")) {
		return errors.New("endpoint must end with /api/put")
	}
	return nil
}
