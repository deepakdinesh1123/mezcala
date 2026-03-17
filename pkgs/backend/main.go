package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/deepakdinesh1123/mezcala/pkgs/backend/api"
	"github.com/deepakdinesh1123/mezcala/pkgs/backend/config"
	"github.com/deepakdinesh1123/mezcala/pkgs/backend/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	envConfig := config.GetEnvConfig()

	logger := log.GetLogger(log.NewLogConfig())

	srv, err := api.NewServer(ctx, envConfig, logger)
	if err != nil {
		logger.Fatal().Err(err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Info().Msg("Shutdown signal received, gracefully shutting down...")
		cancel()
		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Error().Err(err).Msg("Error during server shutdown")
		}
	}()

	err = srv.ListenAndServe()
	if err != nil {
		logger.Error().Err(err).Msg("Error starting server")
	}
}
