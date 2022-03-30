package opentsdbexporter

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.uber.org/zap"
)

func NewFactory() component.ExporterFactory {
	return component.NewExporterFactory(
		"opentsdb",
		createDefaultConfig,
		component.WithMetricsExporter(createMetricsExporter),
	)
}

func createMetricsExporter(_ context.Context, settings component.ExporterCreateSettings, config config.Exporter) (component.MetricsExporter, error) {
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

func createDefaultConfig() config.Exporter {
	return &Config{
		ExporterSettings: config.NewExporterSettings(config.NewComponentID("opentsdb")),
		BatchSize:        20,
		MaxTags:          8,
		SkipTags:         make([]string, 0),
	}
}
