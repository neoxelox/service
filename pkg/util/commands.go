package util

import (
	"context"

	"github.com/mkideal/cli"
	"github.com/neoxelox/kit"

	"service/pkg/config"
)

const (
	DatabaseCommandsRollback = "rollback"
)

type DatabaseCommands struct {
	config   config.Config
	observer *kit.Observer
	migrator *kit.Migrator
}

func NewDatabaseCommands(observer *kit.Observer, migrator *kit.Migrator, config config.Config) *DatabaseCommands {
	return &DatabaseCommands{
		config:   config,
		observer: observer,
		migrator: migrator,
	}
}

type DatabaseCommandsRollbackArgs struct {
	cli.Helper
	Version int  `cli:"*version" usage:"target migration version"`
	DryRun  bool `cli:"dry-run" dft:"true" usage:"whether it is a dry run"`
}

func (self *DatabaseCommands) Rollback(ctx context.Context, command *cli.Context) error {
	args, ok := command.Argv().(*DatabaseCommandsRollbackArgs)
	if !ok {
		return kit.ErrRunnerGeneric.Raise().With("cannot get command arguments")
	}

	version, dirty, err := self.migrator.Version(ctx)
	if err != nil {
		return err
	}

	self.observer.Infof(ctx, "current schema: version=%d, dirty=%v", version, dirty)
	self.observer.Infof(ctx, "desired schema: version=%d", args.Version)

	if !args.DryRun {
		err = self.migrator.Rollback(ctx, args.Version)
		if err != nil {
			return err
		}
	} else {
		self.observer.Info(ctx, "rollback command runned dry")
	}

	return nil
}
