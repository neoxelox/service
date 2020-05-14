package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/neoxelox/microservice-template/internal/server"
)

func main() {

	e := echo.New()
	app := server.NewServer(e)
	go app.Run()

	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Instance.Logger.Fatal(err)
	}

}
