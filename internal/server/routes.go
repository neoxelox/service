package server

import (
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/neoxelox/microservice-template/internal/middleware"
)

// SetupRoutes assigns middlewares and handlers to the application routes
func SetupRoutes(instance *echo.Echo, handlers Handlers) {
	instance.Use(emiddleware.Recover())
	instance.Use(middleware.Logrus())

	v1 := instance.Group("/v1")
	/*-*/ cookie := v1.Group("/cookie")
	/*-------*/ cookie.GET("", handlers.Cookie.List)
	/*-------*/ cookie.POST("", handlers.Cookie.Create)
	/*-------*/ cookie.DELETE("", handlers.Cookie.Delete)
	/*-------*/ cookie.PUT("", handlers.Cookie.Modify)
}
