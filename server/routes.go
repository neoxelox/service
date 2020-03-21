package server

import (
	"io"

	"github.com/friendsofgo/errors"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"mst/config"
	"mst/dependencies"
	own_middleware "mst/server/middleware"
)

// Services contains the internal App Services
type Services struct {
	Sentry     *sentry.Client
	Prometheus *prometheus.Prometheus
	Jaeger     io.Closer
}

// SetupRoutes initializes routes and it's middlewares
func SetupRoutes(cfg *config.Config, app *echo.Echo, dependencies *dependencies.Dependencies) (*Services, error) {

	// Middlewares

	app.Use(middleware.Recover())

	app.Use(own_middleware.Logrus())

	err := sentry.Init(sentry.ClientOptions{
		Dsn:        cfg.Sentry.Dsn,
		Debug:      cfg.App.Debug,
		ServerName: cfg.App.Name,
		Release:    cfg.App.Release,
	})
	if err != nil {
		app.Logger.Fatalf("Sentry initialization failed: %v\n", err)
		return &Services{}, errors.Wrap(err, "Sentry initialization failed")
	}
	sentryMiddleware := sentryecho.New(sentryecho.Options{
		Repanic: true,
	})
	app.Use(sentryMiddleware) // https://docs.sentry.io/platforms/go/echo/#usage

	prometheusClient := prometheus.NewPrometheus(cfg.App.Name, nil)
	prometheusClient.Use(app)

	// Needs credentials
	//app.Use(own_middleware.NewRelic("app name", "license_key"))

	jaegerClient := jaegertracing.New(app, nil)

	// ----------

	// Routes

	v1 := app.Group("/v1")
	/*-*/ cookie := v1.Group("/cookie")
	/*-------*/ cookie.GET("", dependencies.CookieHandler.List)
	/*-------*/ cookie.POST("", dependencies.CookieHandler.Create)
	/*-------*/ cookie.DELETE("", dependencies.CookieHandler.Delete)
	/*-------*/ cookie.PUT("", dependencies.CookieHandler.Modify)

	// ------

	return &Services{
		Sentry:     sentry.CurrentHub().Client(),
		Prometheus: prometheusClient,
		Jaeger:     jaegerClient,
	}, nil
}
