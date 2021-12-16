package opentsdbexporter

import (
	"github.com/pkg/errors"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
	"strings"
)

type Config struct {
	config.ExporterSettings       `mapstructure:",squash"`
	confighttp.HTTPClientSettings `mapstructure:",squash"`
}

var _ config.Exporter = (*Config)(nil)

func (cfg *Config) Validate() error {
	if !(strings.HasPrefix(cfg.Endpoint, "http://") || strings.HasPrefix(cfg.Endpoint, "https")) {
		return errors.New("endpoint must start with https:// or http://")
	}
	return nil
}
