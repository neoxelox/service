package dependencies

import (
	"context"

	"github.com/jackc/pgx/v4"

	"mst/config"
	"mst/database"
	cookierepository "mst/database/repository/cookie"
	cookiemethods "mst/logic/entity/cookie"
	cookiehandler "mst/server/handler/cookie"
)

// Dependencies contains the internal App Dependencies
type Dependencies struct {
	DB            *pgx.Conn
	CookieHandler cookiehandler.Handler
}

// Initialize setups all App Dependencies
func Initialize(cfg *config.Config) (*Dependencies, error) {
	db, err := database.Connect(context.Background(), cfg.Database.Dsn, 5)
	if err != nil {
		panic(err)
	}

	cookieRepository, err := cookierepository.NewSQLRepository(db)
	if err != nil {
		panic(err)
	}
	cookieMethods := cookiemethods.NewCookieMethods(cookieRepository)
	cookieHandler := cookiehandler.NewCookieHandler(cookieMethods)

	return &Dependencies{
		DB:            db,
		CookieHandler: cookieHandler,
	}, nil
}
