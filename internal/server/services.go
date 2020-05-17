package server

import (
	"io"

	"github.com/friendsofgo/errors"
	"github.com/getsentry/sentry-go"
	esentry "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

// Services describes the application external services
type Services struct {
	Sentry     *sentry.Client
	Prometheus *prometheus.Prometheus
	Jaeger     io.Closer
}

// NewServices creates a new Services instance
func NewServices(instance *echo.Echo, configuration Configuration) (*Services, error) {
	prometheusClient := prometheus.NewPrometheus(configuration.App.Name, nil)
	prometheusClient.Use(instance)

	jaegerClient := jaegertracing.New(instance, nil)

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              configuration.Sentry.Dsn,
		Debug:            configuration.App.Debug,
		ServerName:       configuration.App.Name,
		Release:          configuration.App.Release,
		AttachStacktrace: true,
	})
	if err != nil {
		err = errors.Wrap(err, "Unable to initialize Sentry client!")
	}
	sentryMiddleware := esentry.New(esentry.Options{
		Repanic: true,
	})
	instance.Use(sentryMiddleware) // https://docs.sentry.io/platforms/go/echo/#usage

	return &Services{
		Sentry:     sentry.CurrentHub().Client(),
		Prometheus: prometheusClient,
		Jaeger:     jaegerClient,
	}, err
}
