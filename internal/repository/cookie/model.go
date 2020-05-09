package repository

import (
	"context"
	"time"

	"github.com/rs/xid"
)

type Model struct {
	ID         xid.ID    `db:"id"`
	Title      string    `db:"title"`
	Category   *string   `db:"category"`
	Fortune    string    `db:"fortune"`
	createdAt  time.Time `db:"created_at"`
	modifiedAt time.Time `db:"modified_at"`
}

type Repository interface {
	CreateOrUpdate(ctx context.Context, cookie Model) (*Model, error)
	GetByID(ctx context.Context, cookieID xid.ID) (*Model, error)
	ListAll(ctx context.Context) ([]*Model, error)
	Delete(ctx context.Context, cookieID xid.ID) error
}
