package opentsdbreceiver

import (
	"context"
	"fmt"
	"github.com/foxlegend/opentelemetry-collector-opentsdb-contrib/receiver/opentsdbreceiver/internal"
	"github.com/reiver/go-telnet"
	"github.com/soheilhy/cmux"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
	"net"
	"net/http"
	"strings"
	"sync"
)

type metricsReceiver struct {
	logger        *zap.Logger
	tcpServerAddr *confignet.TCPAddr

	metricsConsumer consumer.Metrics

	ln             net.Listener
	httpServer     *http.Server
	telnetServer   *telnet.Server
	telnetServerLn net.Listener
	wg             sync.WaitGroup

	settings component.TelemetrySettings
}

func newMetricsReceiver(config *Config, logger *zap.Logger, settings component.TelemetrySettings, nextConsumer consumer.Metrics) (*metricsReceiver, error) {
	if nextConsumer == nil {
		return nil, component.ErrNilNextConsumer
	}

	receiver := &metricsReceiver{
		logger:          logger,
		metricsConsumer: nextConsumer,
		tcpServerAddr:   &config.TCPAddr,
		settings:        settings,
	}
	return receiver, nil
}

func (r *metricsReceiver) Start(_ context.Context, host component.Host) error {
	ln, err := r.tcpServerAddr.Listen()
	if err != nil {
		return fmt.Errorf("failed to bind to address %s: %w", r.tcpServerAddr.Endpoint, err)
	}
	r.ln = ln

	m := cmux.New(r.ln)
	httpl := m.Match(
		cmux.HTTP1(),
		cmux.HTTP2(),
	)
	// If not matched by HTTP, we assume it is a telnet connection
	telnetl := m.Match(
		cmux.Any(),
	)

	r.httpServer, err = r.newHttpServer(host)
	if err != nil {
		return fmt.Errorf("failed to instantiate HTTP httpServer: %w", err)
	}

	r.telnetServer = &telnet.Server{Handler: OpenTSDBTelnetHandler{}}
	r.telnetServerLn = telnetl

	r.wg.Add(3)
	go func() {
		defer r.wg.Done()
		if err := r.httpServer.Serve(httpl); err != nil && err != http.ErrServerClosed {
			host.ReportFatalError(err)
		}
	}()
	go func() {
		defer r.wg.Done()
		if err := r.telnetServer.Serve(telnetl); err != nil && err != cmux.ErrServerClosed {
			host.ReportFatalError(err)
		}
	}()
	go func() {
		defer r.wg.Done()
		if err := m.Serve(); err != nil && err != cmux.ErrServerClosed && !strings.Contains(err.Error(), "use of closed network connection") {
			host.ReportFatalError(err)
		}
	}()
	return nil
}

func (r *metricsReceiver) Shutdown(ctx context.Context) error {
	var err error
	if r.httpServer != nil {
		e := r.httpServer.Close()
		if e != nil && e != http.ErrServerClosed && !strings.Contains(e.Error(), "use of closed network connection") {
			err = e
		}
	}
	if r.telnetServerLn != nil {
		e := r.telnetServerLn.Close()
		if e != nil && e != cmux.ErrServerClosed && !strings.Contains(e.Error(), "use of closed network connection") {
			err = e
		}
	}
	if r.ln != nil {
		e := r.ln.Close()
		if e != nil && e != cmux.ErrServerClosed && !strings.Contains(e.Error(), "use of closed network connection") {
			err = e
		}
	}
	return err
}

func (r *metricsReceiver) newHttpServer(host component.Host, opts ...confighttp.ToServerOption) (*http.Server, error) {
	// initialize somme dummy config to take advantage of OTEL observability
	dummyConfig := &confighttp.HTTPServerSettings{Endpoint: r.tcpServerAddr.Endpoint}
	httpHandler := internal.NewHttpHandler(r.logger, r.metricsConsumer)
	return dummyConfig.ToServer(host, r.settings, httpHandler.NewHttpRouter(), opts...)
}
