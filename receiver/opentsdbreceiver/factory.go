package opentsdbreceiver

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer"
)

const (
	typeStr   = "opentsdb"
	stability = component.StabilityLevelInDevelopment
)

func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createMetricsReceiver, stability))
}

func createDefaultConfig() config.Receiver {
	return &Config{
		ReceiverSettings: config.NewReceiverSettings(config.NewComponentID(typeStr)),
		TCPAddr: confignet.TCPAddr{
			Endpoint: "0.0.0.0:4242",
		},
	}
}

func createMetricsReceiver(
	_ context.Context,
	params component.ReceiverCreateSettings,
	config config.Receiver,
	nextConsumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	cfg := config.(*Config)
	return newMetricsReceiver(cfg, params.Logger, params.TelemetrySettings, nextConsumer)
}
