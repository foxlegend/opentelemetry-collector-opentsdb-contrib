package opentsdbexporter

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap"
)

const (
		stability = component.StabilityLevelDevelopment
)

func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		"opentsdb",
		createDefaultConfig,
		exporter.WithMetrics(createMetricsExporter, stability),
	)
}

func createMetricsExporter(_ context.Context, settings exporter.CreateSettings, config component.Config) (exporter.Metrics, error) {
	cfg := config.(*Config)

	exporterLogger, err := createLogger(cfg)
	if err != nil {
		return nil, err
	}

	return NewMetricsExporter(config, exporterLogger, settings)
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

func createDefaultConfig() component.Config {
	return &Config{
		TimeoutSettings: exporterhelper.NewDefaultTimeoutSettings(),
		RetrySettings: exporterhelper.NewDefaultRetrySettings(),
		BatchSize:        20,
		MaxTags:          8,
		SkipTags:         make([]string, 0),
	}
}
