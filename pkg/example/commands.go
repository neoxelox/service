package example

import (
	"context"

	"github.com/mkideal/cli"
	"github.com/neoxelox/kit"

	"service/pkg/config"
)

const (
	ExampleCommandsForceOnboarding = "force-onboarding"
)

type ExampleCommands struct {
	config   config.Config
	observer *kit.Observer
	enqueuer *kit.Enqueuer
}

func NewExampleCommands(observer *kit.Observer, enqueuer *kit.Enqueuer, config config.Config) *ExampleCommands {
	return &ExampleCommands{
		config:   config,
		observer: observer,
		enqueuer: enqueuer,
	}
}

type ExampleCommandsForceOnboardingArgs struct {
	cli.Helper
	ID string `cli:"*id" usage:"example id"`
}

func (self *ExampleCommands) ForceOnboarding(ctx context.Context, command *cli.Context) error {
	args, ok := command.Argv().(*ExampleCommandsForceOnboardingArgs)
	if !ok {
		return kit.ErrRunnerGeneric.Raise().With("cannot get command arguments")
	}

	self.observer.Infof(ctx, "forcing onboarding manually to %s", args.ID)

	err := self.enqueuer.Enqueue(ctx, ExampleTasksMakeOnboarding, ExampleTasksMakeOnboardingParams{ID: args.ID})
	if err != nil {
		return err
	}

	return nil
}
