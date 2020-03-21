package main

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"

	"mst/config"
	"mst/dependencies"
	"mst/server"
)

func main() {
	// Load Config
	cfg := config.Load()

	// Init App Dependencies
	dependencies, err := dependencies.Initialize(cfg)
	if err != nil {
		panic(err)
	}
	defer dependencies.DB.Close(context.Background())

	// Create App
	app := echo.New()

	// Setup Routes with Middlewares
	services, err := server.SetupRoutes(cfg, app, dependencies)
	if err != nil {
		panic(err)
	}
	defer services.Sentry.Flush(time.Second * 5)
	defer services.Jaeger.Close()

	// Start extra Asynchronous Services with goroutines
	// -

	// Start App
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", cfg.App.Port)))
}
