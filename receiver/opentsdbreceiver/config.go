package opentsdbreceiver

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confignet"
)

type Config struct {
	config.ReceiverSettings `mapstructure:"-"`
	confignet.TCPAddr       `mapstructure:",squash"`
}
