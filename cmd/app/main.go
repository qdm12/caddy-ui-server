package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/qdm12/golibs/healthcheck"
	"github.com/qdm12/golibs/logging"
	"github.com/qdm12/golibs/server"

	"github.com/qdm12/caddy-ui-server/internal/handlers"
	"github.com/qdm12/caddy-ui-server/internal/params"
	"github.com/qdm12/caddy-ui-server/internal/processor"
	"github.com/qdm12/caddy-ui-server/internal/splash"
)

func main() {
	ctx := context.Background()
	os.Exit(_main(ctx))
}

func _main(ctx context.Context) int {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if healthcheck.Mode(os.Args) {
		// Running the program in a separate instance through the Docker
		// built-in healthcheck, in an ephemeral fashion to query the
		// long running instance of the program about its status
		if err := healthcheck.Query(); err != nil {
			fmt.Println(err)
			return 1
		}
		return 0
	}
	paramsReader := params.NewReader()
	fmt.Println(splash.Splash(
		paramsReader.GetVersion(),
		paramsReader.GetVcsRef(),
		paramsReader.GetBuildDate()))
	logger, err := createLogger(paramsReader)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	listeningPort, warning, err := paramsReader.GetListeningPort()
	if len(warning) > 0 {
		logger.Warn(warning)
	}
	if err != nil {
		logger.Error(err)
		return 1
	}
	rootURL, err := paramsReader.GetRootURL()
	if err != nil {
		logger.Error(err)
		return 1
	}
	caddyAPIEndpoint, err := paramsReader.GetCaddyAPIEndpoint()
	if err != nil {
		logger.Error(err)
		return 1
	}
	corsWhitelist, err := paramsReader.GetCorsWhitelist()
	if err != nil {
		logger.Error(err)
		return 1
	}

	proc := processor.NewProcessor(caddyAPIEndpoint, logger)
	productionHandlerFunc := handlers.NewHandler(rootURL, proc, logger, corsWhitelist)
	healthcheckHandlerFunc := healthcheck.GetHandler(func() error { return nil })
	logger.Info("Server listening at address 0.0.0.0:%s with root URL /%s", listeningPort, rootURL)
	serverErrors := make(chan []error)
	go func() {
		serverErrors <- server.RunServers(ctx,
			server.Settings{Name: "production", Addr: "0.0.0.0:" + listeningPort, Handler: productionHandlerFunc},
			server.Settings{Name: "healthcheck", Addr: "127.0.0.1:9999", Handler: healthcheckHandlerFunc},
		)
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)
	select {
	case errors := <-serverErrors:
		for _, err := range errors {
			logger.Error(err)
		}
		return 1
	case signal := <-osSignals:
		message := fmt.Sprintf("Stopping program: caught OS signal %q", signal)
		logger.Warn(message)
		return 2
	case <-ctx.Done():
		message := fmt.Sprintf("Stopping program: %s", ctx.Err())
		logger.Warn(message)
		return 1
	}
}

func createLogger(paramsReader params.Reader) (logger logging.Logger, err error) {
	encoding, level, nodeID, err := paramsReader.GetLoggerConfig()
	if err != nil {
		return nil, err
	}
	return logging.NewLogger(encoding, level, nodeID)
}
