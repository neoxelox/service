package main

import (
	"net/http"

	"mst/config"
	own_middleware "mst/middleware"

	//"github.com/getsentry/sentry-go"
	//sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load Config
	cfg := config.Load()

	// Init App
	e := echo.New()

	// Middlewares
	e.Use(middleware.Recover())

	e.Use(own_middleware.Logrus())

	// Needs credentials
	// err := sentry.Init(sentry.ClientOptions{
	// 	Dsn: cfg.Sentry.Dsn,
	// })
	// if err != nil {
	// 	e.Logger.Fatalf("Sentry initialization failed: %v\n", err)
	// }
	// defer sentry.Flush(time.Second * 5)
	// // https://docs.sentry.io/platforms/go/echo/#usage
	// e.Use(sentryecho.New(sentryecho.Options{
	// 	Repanic: true,
	// }))

	p := prometheus.NewPrometheus(cfg.Prometheus.SubsystemName, nil)
	p.Use(e)

	// Needs credentials
	//e.Use(own_middleware.NewRelic("app name", "license_key"))

	c := jaegertracing.New(e, nil)
	defer c.Close()

	// ----------

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start App
	e.Logger.Fatal(e.Start(":8000"))
}
