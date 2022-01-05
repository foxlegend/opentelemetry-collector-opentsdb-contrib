FROM alpine:latest

COPY dist/otelcol-custom /
COPY config/otelcol.yaml /config/otelcol.yaml

ENTRYPOINT ["/otelcol-custom", "--config", "/config/otelcol.yaml"]
