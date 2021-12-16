# OpenTSDB support for OpenTelemetry Collector.

This project aims to implement OpenTSDB receiver and exporter for OpenTelemetry Collector.

**This is currently WIP!**

## Development

Go 1.17.* is expected.

### Install required tools

```shell
# Build OpenTelemetry collector
make install-tools
```

Especially, it will install the OpenTelemetry Collector builder.

### Build OpenTelemetry Collector

```shell
make build-otelcol
```

OpenTelemetry Collector builder descriptor is located in `config/otelcol-builder.yaml`.

### Run OpenTelemetry Collector

You can start some dependencies with `docker-compose`:

* Telegraf scraped by the collector,
* OpenTSDB for backend storage,
* Grafana for visualization.

```shell
docker-compose up -d telegraf opentsdb grafana
```

And start the built collector:

```shell
make run-otelcol
```

#### Access

* OpenTSDB: http://localhost:4242
* Grafana: http://localhost:3000

## What is implemented

* Exporter
  * Gauge support
