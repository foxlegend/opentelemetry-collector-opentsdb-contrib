package opentsdbexporter

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
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
	exporterLogger := settings.Logger
	return NewMetricsExporter(config, exporterLogger, settings)
}

func createDefaultConfig() component.Config {
	return &Config{
		TimeoutSettings: exporterhelper.NewDefaultTimeoutSettings(),
		BatchSize:        20,
		MaxTags:          8,
		SkipTags:         make([]string, 0),
	}
}
