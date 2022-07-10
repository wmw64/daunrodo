package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	v1 "daunrodo/internal/delivery/http/v1"
	"daunrodo/internal/usecase"
	"daunrodo/pkg/config"
	httpclient "daunrodo/pkg/http/client"
	httpserver "daunrodo/pkg/http/server"
	"daunrodo/pkg/logger"
	"daunrodo/pkg/service/instagram"

	"github.com/gin-gonic/gin"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	var err error

	l := logger.New(cfg.Level)
	l.Trace("daunrodo starting")

	httpClient := httpclient.New(cfg.HTTP.Proxy)

	// Crawlers
	ig := instagram.New(httpClient.Client)

	// Use cases
	igUseCase := usecase.New(ig)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, igUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
