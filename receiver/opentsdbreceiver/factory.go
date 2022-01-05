package opentsdbreceiver

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"
)

const (
	typeStr = "opentsdb"
)

func NewFactory() component.ReceiverFactory {
	return receiverhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		receiverhelper.WithMetrics(createMetricsReceiver))
}

func createDefaultConfig() config.Receiver {
	return &Config{
		ReceiverSettings: config.NewReceiverSettings(config.NewComponentID(typeStr)),
		HTTPServerSettings: confighttp.HTTPServerSettings{
			Endpoint: "0.0.0.0:4242",
		},
	}
}

func createMetricsReceiver(_ context.Context, params component.ReceiverCreateSettings, config config.Receiver, nextConsumer consumer.Metrics) (component.MetricsReceiver, error) {
	cfg := config.(*Config)
	receiverLogger, err := createLogger(cfg)
	if err != nil {
		return nil, err
	}
	return newMetricsReceiver(cfg, receiverLogger, params.TelemetrySettings, nextConsumer)
}

func createLogger(cfg *Config) (*zap.Logger, error) {
	conf := zap.NewDevelopmentConfig()
	conf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	loggingLogger, err := conf.Build()
	if err != nil {
		return nil, err
	}
	return loggingLogger, nil
}
