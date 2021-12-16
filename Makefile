include ./Makefile.Common

BUILD_OTELCOL=builder
OTELCOL=./dist/otelcol-custom

TOOLS_MOD_DIR := ./internal/tools
.PHONY: install-tools
install-tools:
	cd $(TOOLS_MOD_DIR) && go install go.opentelemetry.io/collector/cmd/builder

.PHONY: build-otelcol
build-otelcol:
	$(BUILD_OTELCOL) --output-path=dist --config=config/otelcol-builder.yaml

.PHONY: run-otelcol
run-otelcol:
	$(OTELCOL) --config config/otelcol.yaml
