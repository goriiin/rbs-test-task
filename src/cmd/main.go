package main

import (
	"context"
	"errors"
	"github.com/goriiin/rbs-test-task/src/internal/configs/postgres_pool"
	wd "github.com/goriiin/rbs-test-task/src/internal/delivery/weather"
	wr "github.com/goriiin/rbs-test-task/src/internal/repository/weather"
	"github.com/goriiin/rbs-test-task/src/internal/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	envAppServiceName   = "APP_SERVICE_NAME"
	weatherTemplatePath = "/templates/weather.html"
	healthyTemplatePath = "/templates/healthy.html"
	addPath             = "/templates/add.html"
	needDB              = "main_db"
	serverAddress       = ":8050"
	shutdownTimeout     = 30 * time.Second
)

func main() {
	requiredDbConfigs, err := postgres_pool.LoadForCurrentService(envAppServiceName)
	if err != nil {
		log.Fatalf("[ main ] FATAL: Could not load configuration: %v", err)
	}
	log.Println("[ main ] INFO: Configuration loaded.")

	dbPools, err := postgres_pool.InitConnections(requiredDbConfigs)
	if err != nil {
		log.Fatalf("[ main ] FATAL: Could not initialize required db conns: %v", err)
	}
	log.Println("[ main ] INFO: Database connections initialized.")

	mainDbPool, ok := dbPools[needDB]
	if !ok {
		log.Fatal("[ main ] FATAL: main_db not found in cache")
	}

	weatherRepo := wr.NewWeatherRepository(mainDbPool)

	weatherHandler, err := wd.NewWeatherDelivery(weatherRepo, weatherTemplatePath, healthyTemplatePath, addPath)
	if err != nil {
		log.Fatalf("[ main ] FATAL: Could not initialize weather handler: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", weatherHandler.Ping)
	mux.HandleFunc("/list", weatherHandler.List)
	mux.HandleFunc("/health", weatherHandler.Health)
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			weatherHandler.Show(w, r)
		} else if r.Method == http.MethodPost {
			weatherHandler.Add(w, r)
		} else {
			utils.WriteJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})

	server := &http.Server{
		Addr:    serverAddress,
		Handler: mux,
	}

	serverErr := make(chan error, 1)
	go func() {
		log.Printf("[ main ] INFO: Starting web server on %s", serverAddress)

		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("[ main ] INFO: Shutdown signal received.")
	case err = <-serverErr:
		log.Fatalf("[ main ] FATAL: Server failed to start: %v", err)
	}

	log.Println("[ main ] INFO: Starting graceful shutdown.")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatalf("[ main ] FATAL: Server shutdown failed: %v", err)
	}

	log.Println("[ main ] INFO: Closing database connections.")
	for name, pool := range dbPools {
		pool.Close()
		log.Printf("[ main ] INFO: Connection pool '%s' closed.", name)
	}

	log.Println("[ main ] INFO: Server gracefully stopped.")
}
