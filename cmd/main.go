package main

import (
	"context"
	"errors"
	"goth/internal/config"
	"goth/internal/handlers"
	database "goth/internal/store/db"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

)

/*
* Set to production at build time
* used to determine what assets to load
 */
var Environment = "development"

func init() {
	os.Setenv("env", Environment)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg := config.MustLoadConfig()

	db, err := database.MustOpen(cfg.DatabaseName)
  if err != nil {
    logger.Error("failed to start rutrow",
      "Error", err)
    // return 
  }

  router, err := handlers.Router(cfg, db)
  if err != nil {
    logger.Error("failed to start router rutrow",
      "Error", err)
    // return 
  }

  start(cfg, router, logger)

	logger.Info("Server shutdown complete")
}

func start(cfg *config.Config, router http.Handler, logger *slog.Logger) {
	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Server shutdown complete")
		} else if err != nil {
			logger.Error("Server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	logger.Info("Server started", slog.String("port", cfg.Port), slog.String("env", Environment))
	<-killSig

	logger.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}
}
