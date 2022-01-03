package opentsdbexporter

import (
	"github.com/pkg/errors"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
	"strings"
)

type Config struct {
	config.ExporterSettings `mapstructure:",squash"`

	// OpenTSDB Endpoint
	confighttp.HTTPClientSettings `mapstructure:",squash"`

	// The maximum number of datapoints to send to the OpenTSDB backend
	BatchSize int `mapstructure:"batch_size"`
	// The maximum number of tags per datapoint
	MaxTags int `mapstructure:"max_tags"`
	// Tags to skip
	SkipTags []string `mapstructure:"skip_tags"`
}

var _ config.Exporter = (*Config)(nil)

func (cfg *Config) Validate() error {
	if !(strings.HasPrefix(cfg.Endpoint, "http://") || strings.HasPrefix(cfg.Endpoint, "https")) {
		return errors.New("endpoint must start with https:// or http://")
	}
	if !(strings.HasSuffix(cfg.Endpoint, "/api/put")) {
		return errors.New("endpoint must end with /api/put")
	}
	return nil
}
