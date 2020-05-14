package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
)

// Server describes the main application instance
type Server struct {
	Instance      *echo.Echo
	Configuration *Configuration
	Dependencies  *Dependencies
	Handlers      *Handlers
	Services      *Services
}

// NewServer creates a new Server instance
func NewServer(e *echo.Echo) *Server {
	configuration, err := NewConfiguration()
	if err != nil {
		panic(errors.Wrap(err, "Unable to initialize the server, configuration error!"))
	}

	dependencies, err := NewDependencies(*configuration)
	if err != nil {
		panic(errors.Wrap(err, "Unable to initialize the server, dependency error!"))
	}

	handlers, err := NewHandlers(*dependencies)
	if err != nil {
		panic(errors.Wrap(err, "Unable to initialize the server, handler error!"))
	}

	SetupRoutes(e, *handlers)

	services, err := NewServices(e, *configuration)
	if err != nil {
		log.Println("Unable to initialize external services, ignoring...")
	}

	return &Server{
		Instance:      e,
		Configuration: configuration,
		Dependencies:  dependencies,
		Handlers:      handlers,
		Services:      services,
	}
}

// Run starts the server
func (s *Server) Run() {
	defer s.Dependencies.Database.Close(context.Background())
	defer s.Services.Sentry.Flush(time.Second * 5)
	defer s.Services.Jaeger.Close()

	s.Instance.Logger.Fatal(s.Instance.Start(fmt.Sprintf(":%d", s.Configuration.App.Port)))
}

// Shutdown stops the server
func (s *Server) Shutdown(ctx context.Context) error {
	deadline := time.Second * 5
	if ctxDeadline, ok := ctx.Deadline(); ok {
		deadline = ctxDeadline.Sub(time.Now())
	}

	err := s.Dependencies.Database.Close(ctx)
	if err != nil {
		return errors.Wrap(err, "Unable to close the connection with the database!")
	}

	s.Services.Sentry.Flush(deadline)
	s.Services.Jaeger.Close()

	return s.Instance.Shutdown(ctx)
}
