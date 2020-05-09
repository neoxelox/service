package entity

import (
	"context"

	"github.com/rs/xid"

	repository "github.com/neoxelox/microservice-template/internal/repository/cookie"
)

type Cookie struct {
	ID       xid.ID
	Title    string
	Category *string
	Fortune  string
}

func NewCookie(title string, category *string,
	fortune string) *Cookie {
	return &Cookie{
		ID:       xid.New(),
		Title:    title,
		Category: category,
		Fortune:  fortune,
	}
}

func NewCookieFromModel(model *repository.Model) *Cookie {
	return &Cookie{
		ID:       model.ID,
		Title:    model.Title,
		Category: model.Category,
		Fortune:  model.Fortune,
	}
}

func (e *Cookie) ToModel() *repository.Model {
	return &repository.Model{
		ID:       e.ID,
		Title:    e.Title,
		Category: e.Category,
		Fortune:  e.Fortune,
	}
}

type Methods interface {
	Create(ctx context.Context, title string, category *string) (*Cookie, error)
	ListAll(ctx context.Context) ([]*Cookie, error)
	Delete(ctx context.Context, cookieID xid.ID) error
	Modify(ctx context.Context, cookieID xid.ID, title string, category *string) (*Cookie, error)
}

type CookieMethods struct {
	repository repository.Repository
}

func NewCookieMethods(repository repository.Repository) Methods {
	return &CookieMethods{
		repository: repository,
	}
}
