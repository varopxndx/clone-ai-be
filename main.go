package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/varopxndx/clone-ai-be/config"
	"github.com/varopxndx/clone-ai-be/controller"
	"github.com/varopxndx/clone-ai-be/router"
	"github.com/varopxndx/clone-ai-be/service"
	"github.com/varopxndx/clone-ai-be/usecase"

	"github.com/rs/zerolog"
)

// Abnormal exit constants
const (
	ExitAbnormalErrorLoadingConfiguration = iota
)

// @title           Auth API
// @version         1.0
// @description     Clone AI

// @host     localhost:8080
// @BasePath  /

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Msg("starting clone-ai-api")

	// read config file
	cfg, err := config.Load()
	if err != nil {
		logger.Info().Msgf("failed to load config: %v", err)
		os.Exit(ExitAbnormalErrorLoadingConfiguration)
	}

	// create service layer
	service := service.New(cfg.OpenAIToken)

	// create usecase layer
	usecase := usecase.New(service, logger)

	// create controller layer
	controller := controller.New(usecase, logger)

	// create router layer
	router := router.New(controller)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("Fatal error starting server: %v \n", err))
		}
	}()

	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("Fatal error shutdown server: %v \n", err))
	}
	log.Println("Server has been stopped...")
}
