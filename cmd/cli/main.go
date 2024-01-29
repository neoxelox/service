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

	cli, err := NewCLI(ctx, *config)
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

		err := cli.Close(ctx) // nolint:govet
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}

		close(done)
	}()

	// CLI run will exit directly by any error returned by the command
	// so don't panic and allow the CLI to close itself gracefully
	err = cli.Run(ctx)
	if err != nil {
		// Because of this print, errors may be printed twice
		// (but only processed once by the observer)
		fmt.Printf("%+v", err) // nolint:forbidigo
	}

	// Check whether everything has already been gracefully closed
	select {
	case <-done:
	default:
		err := cli.Close(ctx) // nolint:govet
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}

		close(done)
	}

	// Wait until everything has been gracefully closed
	<-done

	// Exit CLI following POSIX standard
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
