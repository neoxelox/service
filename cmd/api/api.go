package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/neoxelox/kit"
	kitMiddleware "github.com/neoxelox/kit/middleware"
	kitUtil "github.com/neoxelox/kit/util"

	"service/pkg/config"
	"service/pkg/example"
	"service/pkg/util"
)

type API struct {
	Run   func(ctx context.Context) error
	Close func(ctx context.Context) error
}

// nolint: gocognit, maintidx
func NewAPI(ctx context.Context, config config.Config) (*API, error) {
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

	var observerGilkConfig *kit.ObserverGilkConfig
	if config.Service.Environment == kit.EnvDevelopment {
		observerGilkConfig = &kit.ObserverGilkConfig{
			Port: config.Gilk.Port,
		}
	}

	observer, err := kit.NewObserver(ctx, kit.ObserverConfig{
		Environment: config.Service.Environment,
		Release:     config.Service.Release,
		Service:     config.Service.Name,
		Level:       level,
		Sentry:      observerSentryConfig,
		Gilk:        observerGilkConfig,
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

	err = migrator.Apply(ctx, config.Database.SchemaVersion)
	if err != nil {
		return nil, err
	}

	err = migrator.Assert(ctx, config.Database.SchemaVersion)
	if err != nil {
		return nil, err
	}

	err = migrator.Close(ctx)
	if err != nil {
		return nil, err
	}

	errorHandler := kit.NewErrorHandler(observer, kit.ErrorHandlerConfig{
		Environment: config.Service.Environment,
	})

	serializer := kit.NewSerializer(observer, kit.SerializerConfig{})

	binder := kit.NewBinder(observer, kit.BinderConfig{})

	renderer, err := kit.NewRenderer(observer, kit.RendererConfig{
		TemplatesPath:       kitUtil.Pointer(config.Service.TemplatesPath),
		TemplateFilePattern: kitUtil.Pointer(config.Service.TemplateFilePattern),
	})
	if err != nil {
		return nil, err
	}

	localizer, err := kit.NewLocalizer(observer, kit.LocalizerConfig{
		DefaultLocale:     config.Service.DefaultLocale,
		LocalesPath:       kitUtil.Pointer(config.Service.LocalesPath),
		LocaleFilePattern: kitUtil.Pointer(config.Service.LocaleFilePattern),
	})
	if err != nil {
		return nil, err
	}

	server := kit.NewHTTPServer(observer, serializer, binder, renderer, errorHandler, kit.HTTPServerConfig{
		Environment:              config.Service.Environment,
		Port:                     config.Server.Port,
		RequestHeaderMaxSize:     kitUtil.Pointer(config.Server.RequestHeaderMaxSize),
		RequestBodyMaxSize:       kitUtil.Pointer(config.Server.RequestBodyMaxSize),
		RequestFileMaxSize:       kitUtil.Pointer(config.Server.RequestFileMaxSize),
		RequestFilePattern:       kitUtil.Pointer(config.Server.RequestFilePattern),
		RequestKeepAliveTimeout:  kitUtil.Pointer(config.Server.RequestKeepAliveTimeout),
		RequestReadTimeout:       kitUtil.Pointer(config.Server.RequestReadTimeout),
		RequestReadHeaderTimeout: kitUtil.Pointer(config.Server.RequestReadHeaderTimeout),
		RequestIPExtractor:       kitUtil.Pointer((func(*http.Request) string)(echo.ExtractIPFromXFFHeader())),
		ResponseWriteTimeout:     kitUtil.Pointer(config.Server.ResponseWriteTimeout),
	})

	observerMiddleware := kitMiddleware.NewObserver(observer, kitMiddleware.ObserverConfig{})
	recoverMiddleware := kitMiddleware.NewRecover(observer, kitMiddleware.RecoverConfig{})
	secureMiddleware := kitMiddleware.NewSecure(observer, kitMiddleware.SecureConfig{
		CORSAllowOrigins: kitUtil.Pointer(config.Server.Origins),
	})
	localizerMiddleware := kitMiddleware.NewLocalizer(observer, localizer, kitMiddleware.LocalizerConfig{})
	timeoutMiddleware := kitMiddleware.NewTimeout(observer, kitMiddleware.TimeoutConfig{
		Timeout: config.Service.GracefulTimeout,
	})
	errorMiddleware := kitMiddleware.NewError(observer, kitMiddleware.ErrorConfig{})

	server.Use(observerMiddleware.HandleRequest)
	server.Use(recoverMiddleware.HandleRequest)
	server.Use(secureMiddleware.Handle)
	server.Use(localizerMiddleware.Handle)
	server.Use(timeoutMiddleware.Handle)
	server.Use(errorMiddleware.Handle)

	api := server.Default()

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

	exampleRepository := example.NewExampleRepository(observer, database, config)

	/* SERVICES */

	exampleService := example.NewExampleService(observer, config)

	/* USECASES */

	exampleGetter := example.NewExampleGetter(observer, cache, exampleRepository, config)
	exampleCreator := example.NewExampleCreator(observer, exampleService, exampleRepository, enqueuer, config)
	exampleDeleter := example.NewExampleDeleter(observer, database, exampleRepository, config)

	/* ENDPOINTS */

	healthEndpoints := util.NewHealthEndpoints(observer, database, cache, config)
	fileEndpoints := util.NewFileEndpoints(observer, config)

	exampleEndpoints := example.NewExampleEndpoints(observer, localizer, exampleGetter, exampleCreator, exampleDeleter, config)

	/* MIDDLEWARES */

	exampleMiddlewares := example.NewExampleMiddlewares(observer, config)

	/* ROUTES */

	api.GET("/health", healthEndpoints.GetServerHealth)
	api.Static("/assets", config.Service.AssetsPath)
	api.GET("/file/:name", fileEndpoints.GetFile)
	api.POST("/file", fileEndpoints.PostFile)

	api = api.Group("", exampleMiddlewares.HandleRequest)

	api.GET("/example", exampleEndpoints.GetExample)
	api.GET("/example/:id", exampleEndpoints.GetExampleByID)
	api.POST("/example", exampleEndpoints.PostExample)
	api.DELETE("/example/:id", exampleEndpoints.DeleteExample)

	return &API{
		Run: func(ctx context.Context) error {
			observer.Infof(ctx, "Starting %s API", config.Service.Name)

			err := server.Run(ctx)
			if err != nil {
				return err
			}

			return nil
		},
		Close: func(ctx context.Context) error {
			err := kitUtil.Deadline(ctx, func(exceeded <-chan struct{}) error {
				observer.Infof(ctx, "Closing %s API", config.Service.Name)

				err := server.Close(ctx)
				if err != nil {
					observer.Error(ctx, err)
				}

				err = enqueuer.Close(ctx)
				if err != nil {
					observer.Error(ctx, err)
				}

				err = exampleService.Close(ctx)
				if err != nil {
					observer.Error(ctx, err)
				}

				err = cache.Close(ctx)
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
