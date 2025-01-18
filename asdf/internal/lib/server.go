package lib

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	// "reflect"
	// "runtime"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

const (
	defaultIdleTimeout  = time.Minute
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 10 * time.Second
)

var (
	idleTimeout           = os.Getenv("SERVER_IDLE_TIMEOUT")
	readTimeout           = os.Getenv("SERVER_READ_TIMEOUT")
	writeTimeout          = os.Getenv("SERVER_WRITE_TIMEOUT")
	shutdownPeriod        = os.Getenv("SERVER_GRACEFUL_SHUTDOWN_PERIOD")
	defaultShutdownPeriod = 30 * time.Second
)

func newServer(port int, router http.Handler, logger ILogger) (server *http.Server, err error) {
	server = &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", port),
		Handler:      router,
		IdleTimeout:  defaultIdleTimeout,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelWarn),
	}

	if idleTimeout != "" {
		if i, err := strconv.Atoi(idleTimeout); err == nil {
			server.IdleTimeout = time.Duration(i * int(time.Minute))
		} else {
			return nil, err
		}
	}

	if readTimeout != "" {
		if i, err := strconv.Atoi(readTimeout); err == nil {
			server.ReadTimeout = time.Duration(i * int(time.Second))
		} else {
			return nil, err
		}
	}

	if writeTimeout != "" {
		if i, err := strconv.Atoi(writeTimeout); err == nil {
			server.WriteTimeout = time.Duration(i * int(time.Second))
		} else {
			return nil, err
		}
	}

	if shutdownPeriod != "" {
		if i, err := strconv.Atoi(shutdownPeriod); err == nil {
			defaultShutdownPeriod = time.Duration(i * int(time.Second))
		} else {
			return nil, err
		}
	}

	return server, nil
}

func (i *idi) ServeHTTP() error {
	// graceful shutdown
	done := make(chan bool, 1)
	go i.gracefulShutdown(done)

	i.logger.Info(fmt.Sprintf("server started at port %d", i.port))

	return i.server.ListenAndServe()
}

func (i *idi) gracefulShutdown(done chan bool) {
	// Create context that listens for the interrupt signal from the Oi.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	i.logger.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownPeriod)
	defer cancel()
	if err := i.server.Shutdown(ctx); err != nil {
		i.logger.Warn("Server forced to shutdown with error: %v", err)
	}

	i.logger.Warn("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}
