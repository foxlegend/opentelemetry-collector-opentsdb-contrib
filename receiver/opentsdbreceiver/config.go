package opentsdbreceiver

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
)

type Config struct {
	config.ReceiverSettings       `mapstructure:"-"`
	confighttp.HTTPServerSettings `mapstructure:",squash"`
}
