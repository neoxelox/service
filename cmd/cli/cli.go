package main

import (
	"context"
	"fmt"
	"time"

	"github.com/neoxelox/kit"
	kitMiddleware "github.com/neoxelox/kit/middleware"
	kitUtil "github.com/neoxelox/kit/util"

	"service/pkg/config"
	"service/pkg/example"
	"service/pkg/util"
)

type CLI struct {
	Run   func(ctx context.Context) error
	Close func(ctx context.Context) error
}

// nolint: gocognit, maintidx
func NewCLI(ctx context.Context, config config.Config) (*CLI, error) {
	retry := kit.RetryConfig{
		Attempts:     5,
		InitialDelay: 1 * time.Second,
		LimitDelay:   5 * time.Second,
	}

	level := kit.LvlInfo
	if config.Service.Environment == kit.EnvDevelopment {
		level = kit.LvlDebug
	}

	var observerSentryConfig *kit.ObserverSentryConfig
	if config.Service.Environment == kit.EnvProduction {
		observerSentryConfig = &kit.ObserverSentryConfig{
			Dsn: config.Sentry.DSN,
		}
	}

	observer, err := kit.NewObserver(ctx, kit.ObserverConfig{
		Environment: config.Service.Environment,
		Release:     config.Service.Release,
		Service:     config.Service.Name,
		Level:       level,
		Sentry:      observerSentryConfig,
	}, retry)
	if err != nil {
		return nil, err
	}

	migrator, err := kit.NewMigrator(ctx, observer, kit.MigratorConfig{
		DatabaseHost:     config.Database.Host,
		DatabasePort:     config.Database.Port,
		DatabaseSSLMode:  config.Database.SSLMode,
		DatabaseUser:     config.Database.User,
		DatabasePassword: config.Database.Password,
		DatabaseName:     config.Database.Name,
		MigrationsPath:   kitUtil.Pointer(config.Service.MigrationsPath),
	}, retry)
	if err != nil {
		return nil, err
	}

	errorHandler := kit.NewErrorHandler(observer, kit.ErrorHandlerConfig{
		Environment: config.Service.Environment,
	})

	// NOTE: You can use this too
	_, err = kit.NewRenderer(observer, kit.RendererConfig{
		TemplatesPath:       kitUtil.Pointer(config.Service.TemplatesPath),
		TemplateFilePattern: kitUtil.Pointer(config.Service.TemplateFilePattern),
	})
	if err != nil {
		return nil, err
	}

	// NOTE: You can use this too
	_, err = kit.NewLocalizer(observer, kit.LocalizerConfig{
		DefaultLocale:     config.Service.DefaultLocale,
		LocalesPath:       kitUtil.Pointer(config.Service.LocalesPath),
		LocaleFilePattern: kitUtil.Pointer(config.Service.LocaleFilePattern),
	})
	if err != nil {
		return nil, err
	}

	runner := kit.NewRunner(observer, errorHandler, kit.RunnerConfig{
		Service: config.Service.Name,
		Release: config.Service.Release,
	})

	observerMiddleware := kitMiddleware.NewObserver(observer, kitMiddleware.ObserverConfig{})
	recoverMiddleware := kitMiddleware.NewRecover(observer, kitMiddleware.RecoverConfig{})

	runner.Use(observerMiddleware.HandleCommand)
	runner.Use(recoverMiddleware.HandleCommand)

	database, err := kit.NewDatabase(ctx, observer, kit.DatabaseConfig{
		Host:                  config.Database.Host,
		Port:                  config.Database.Port,
		SSLMode:               config.Database.SSLMode,
		User:                  config.Database.User,
		Password:              config.Database.Password,
		Database:              config.Database.Name,
		Service:               config.Service.Name,
		MinConns:              kitUtil.Pointer(config.Database.MinConns),
		MaxConns:              kitUtil.Pointer(config.Database.MaxConns),
		MaxConnIdleTime:       kitUtil.Pointer(config.Database.MaxConnIdleTime),
		MaxConnLifeTime:       kitUtil.Pointer(config.Database.MaxConnLifeTime),
		DialTimeout:           kitUtil.Pointer(config.Database.DialTimeout),
		StatementTimeout:      kitUtil.Pointer(config.Database.StatementTimeout),
		DefaultIsolationLevel: kitUtil.Pointer(config.Database.DefaultIsolationLevel),
	}, retry)
	if err != nil {
		return nil, err
	}

	cache, err := kit.NewCache(ctx, observer, kit.CacheConfig{
		Host:            config.Cache.Host,
		Port:            config.Cache.Port,
		SSLMode:         config.Cache.SSLMode,
		Password:        config.Cache.Password,
		MinConns:        kitUtil.Pointer(config.Cache.MinConns),
		MaxConns:        kitUtil.Pointer(config.Cache.MaxConns),
		MaxConnIdleTime: kitUtil.Pointer(config.Cache.MaxConnIdleTime),
		MaxConnLifeTime: kitUtil.Pointer(config.Cache.MaxConnLifeTime),
		ReadTimeout:     kitUtil.Pointer(config.Cache.ReadTimeout),
		WriteTimeout:    kitUtil.Pointer(config.Cache.WriteTimeout),
		DialTimeout:     kitUtil.Pointer(config.Cache.DialTimeout),
	}, retry)
	if err != nil {
		return nil, err
	}

	enqueuer := kit.NewEnqueuer(observer, kit.EnqueuerConfig{
		CacheHost:         config.Cache.Host,
		CachePort:         config.Cache.Port,
		CacheSSLMode:      config.Cache.SSLMode,
		CachePassword:     config.Cache.Password,
		CacheMaxConns:     kitUtil.Pointer(config.Cache.MaxConns),
		CacheReadTimeout:  kitUtil.Pointer(config.Cache.ReadTimeout),
		CacheWriteTimeout: kitUtil.Pointer(config.Cache.WriteTimeout),
		CacheDialTimeout:  kitUtil.Pointer(config.Cache.DialTimeout),
	})

	/* REPOSITORIES  */

	/* SERVICES */

	/* USECASES */

	/* COMMANDS */

	databaseCommands := util.NewDatabaseCommands(observer, migrator, config)

	exampleCommands := example.NewExampleCommands(observer, enqueuer, config)

	/* MIDDLEWARES */

	exampleMiddlewares := example.NewExampleMiddlewares(observer, config)

	/* REGISTRATIONS */

	runner.Register(util.DatabaseCommandsRollback, databaseCommands.Rollback,
		util.DatabaseCommandsRollbackArgs{}, "rollback migrations")

	runner.Use(exampleMiddlewares.HandleCommand)

	runner.Register(example.ExampleCommandsForceOnboarding, exampleCommands.ForceOnboarding,
		example.ExampleCommandsForceOnboardingArgs{}, "force onboarding to an example")

	return &CLI{
		Run: func(ctx context.Context) error {
			observer.Infof(ctx, "Starting %s CLI", config.Service.Name)

			err := runner.Run(ctx)
			if err != nil {
				return err
			}

			return nil
		},
		Close: func(ctx context.Context) error {
			err := kitUtil.Deadline(ctx, func(exceeded <-chan struct{}) error {
				observer.Infof(ctx, "Closing %s CLI", config.Service.Name)

				err := runner.Close(ctx)
				if err != nil {
					observer.Error(ctx, err)
				}

				err = enqueuer.Close(ctx)
				if err != nil {
					observer.Error(ctx, err)
				}

				err = cache.Close(ctx)
				if err != nil {
					observer.Error(ctx, err)
				}

				err = migrator.Close(ctx)
				if err != nil {
					observer.Error(ctx, err)
				}

				err = database.Close(ctx)
				if err != nil {
					observer.Error(ctx, err)
				}

				err = observer.Close(ctx)
				if err != nil {
					fmt.Printf("%+v", err) // nolint:forbidigo
				}

				return nil
			})
			if err != nil {
				return err
			}

			return nil
		},
	}, nil
}
