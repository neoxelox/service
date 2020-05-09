package database

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
)

// Connect tries to connect to an specified database via the dsn connection string
func Connect(ctx context.Context, dsn string, retries int) (*pgx.Conn, error) {
	delay := time.NewTicker(1 * time.Second)
	timeout := (time.Duration(retries) * time.Second)

	defer delay.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return nil, errors.New("Unable to connect to database!")

		case <-delay.C:
			db, err := pgx.Connect(ctx, dsn)
			if err == nil {
				return db, nil
			}
		}
	}
}
