package opentsdbreceiver

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	typeStr   = "opentsdb"
	stability = component.StabilityLevelDevelopment
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, stability))
}

func createDefaultConfig() component.Config {
	return &Config{
		TCPAddr: confignet.TCPAddr{
			Endpoint: "0.0.0.0:4242",
		},
	}
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.CreateSettings,
	config component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	cfg := config.(*Config)
	return newMetricsReceiver(cfg, params.Logger, params.TelemetrySettings, nextConsumer)
}
