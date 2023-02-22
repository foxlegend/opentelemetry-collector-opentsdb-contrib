package opentsdbreceiver

import (
	"go.opentelemetry.io/collector/config/confignet"
)

type Config struct {
	confignet.TCPAddr       `mapstructure:",squash"`
}
