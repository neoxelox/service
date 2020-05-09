package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/rs/xid"
)

type SQLRepository struct {
	db *pgx.Conn
}

func NewSQLRepository(db *pgx.Conn) (Repository, error) {
	// Make migrations if needed
	return &SQLRepository{
		db: db,
	}, nil
}

func (r *SQLRepository) CreateOrUpdate(ctx context.Context, cookie Model) (*Model, error) {
	return &Model{}, nil
}

func (r *SQLRepository) GetByID(ctx context.Context, cookieID xid.ID) (*Model, error) {
	return &Model{}, nil
}

func (r *SQLRepository) ListAll(ctx context.Context) ([]*Model, error) {
	return nil, nil
}

func (r *SQLRepository) Delete(ctx context.Context, cookieID xid.ID) error {
	return nil
}
