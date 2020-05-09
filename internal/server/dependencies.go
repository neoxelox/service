package server

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/jackc/pgx/v4"
	"github.com/neoxelox/microservice-template/internal/database"
)

// Dependencies describes de application internal services
type Dependencies struct {
	Database *pgx.Conn
}

// NewDependencies creates a new Dependencies instance
func NewDependencies(configuration Configuration) (*Dependencies, error) {
	database, err := database.Connect(context.Background(), configuration.Database.Dsn, 5)
	if err != nil {
		return nil, errors.Wrap(err, "Database dependency failed!")
	}

	return &Dependencies{
		Database: database,
	}, nil
}
