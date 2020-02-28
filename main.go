package main

import (
	"net/http"

	own_middleware "mst/middleware"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Recover())

	e.Use(own_middleware.Logrus())

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// Needs credentials
	//e.Use(own_middleware.NewRelic("app name", "license_key"))

	c := jaegertracing.New(e, nil)
	defer c.Close()

	// ----------

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
