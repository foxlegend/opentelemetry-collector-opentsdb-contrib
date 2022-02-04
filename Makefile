include ./Makefile.Common

BUILD_OTELCOL=builder
OTELCOL=./dist/otelcol-custom

TOOLS_MOD_DIR := ./internal/tools
.PHONY: install-tools
install-tools:
	cd $(TOOLS_MOD_DIR) && go install go.opentelemetry.io/collector/cmd/builder

.PHONY: build-otelcol
build-otelcol:
	CGO_ENABLED=0 $(BUILD_OTELCOL) --output-path=dist --config=config/otelcol-builder.yaml

.PHONY: run-otelcol
run-otelcol:
	$(OTELCOL) --config config/otelcol.yaml

.PHONY: build-otelcol-docker
build-otelcol-docker: build-otelcol
	docker-compose build collector

.PHONY: run-otelcol-docker
run-otelcol-docker: build-otelcol-docker
	docker-compose up collector

.PHONY: run-3rdparty
run-3rdparty:
	docker-compose up -d telegraf opentsdb grafana


## TCollector
#############
start-tcollector:
	python tcollector/tcollector.py -H localhost -p 4242 --http --log-stdout -P tcollector.pid

stop-collector:
	if [ -e tcollector.pid ]; then \
  		kill -TERM $$(cat tcollector.pid) || true; \
	fi;
	rm tcollector.pid
