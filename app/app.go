package app

import (
	"net/http"

	"github.com/desulaidovich/pretty-link/auth/api"
	auth "github.com/desulaidovich/pretty-link/auth/usecase"
	"github.com/desulaidovich/pretty-link/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Run() error {

	cfg, err := config.LoadFromEnv()
	if err != nil {
		return err
	}

	db, err := connectDB(cfg)

	if err != nil {
		return err
	}

	defer db.Close()

	mux := http.NewServeMux()

	authUseCase := auth.New(db)
	api.RegisterAuthEndpoints(mux, authUseCase)

	server := new(http.Server)
	server.Addr = ":" + cfg.Port
	server.Handler = mux

	if err = server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func connectDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.ConnectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
