package server

import (
	"github.com/friendsofgo/errors"

	cookiemethods "github.com/neoxelox/microservice-template/internal/entity/cookie"
	cookiehandler "github.com/neoxelox/microservice-template/internal/handler/cookie"
	cookierepository "github.com/neoxelox/microservice-template/internal/repository/cookie"
)

// Handlers describes the application handlers
type Handlers struct {
	Cookie cookiehandler.Handler
}

// NewHandlers creates a new Handlers instance
func NewHandlers(dependencies Dependencies) (*Handlers, error) {
	cookieRepository, err := cookierepository.NewSQLRepository(dependencies.Database)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create cookie Repository!")
	}
	cookieMethods := cookiemethods.NewCookieMethods(cookieRepository)
	cookieHandler := cookiehandler.NewCookieHandler(cookieMethods)

	return &Handlers{
		Cookie: cookieHandler,
	}, nil
}
