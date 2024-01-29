package example

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/mkideal/cli"
	"github.com/neoxelox/kit"

	"service/pkg/config"
)

type ExampleMiddlewares struct {
	config   config.Config
	observer *kit.Observer
}

func NewExampleMiddlewares(observer *kit.Observer, config config.Config) *ExampleMiddlewares {
	return &ExampleMiddlewares{
		config:   config,
		observer: observer,
	}
}

func (self *ExampleMiddlewares) HandleRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		self.observer.Info(ctx.Request().Context(), "this middleware does nothing!")

		return next(ctx)
	}
}

func (self *ExampleMiddlewares) HandleTask(next asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
		self.observer.Info(ctx, "this middleware does nothing!")

		return next.ProcessTask(ctx, task)
	})
}

func (self *ExampleMiddlewares) HandleCommand(next kit.RunnerHandler) kit.RunnerHandler {
	return func(ctx context.Context, command *cli.Context) error {
		self.observer.Info(ctx, "this middleware does nothing!")

		return next(ctx, command)
	}
}
