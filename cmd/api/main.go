package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	config := NewConfig()

	api, err := NewAPI(ctx, *config)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	done := make(chan bool, 1)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		// Wait until interrupt signal
		<-quit

		ctx, cancel := context.WithTimeout(ctx, config.Service.GracefulTimeout) // nolint:govet
		defer cancel()

		err := api.Close(ctx) // nolint:govet
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}

		close(done)
	}()

	err = api.Run(ctx)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	// Wait until everything has been gracefully closed
	<-done
}
