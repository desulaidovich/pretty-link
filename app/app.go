package app

import (
	"database/sql"
	"net/http"

	"github.com/desulaidovich/pretty-link/auth/api"
	"github.com/desulaidovich/pretty-link/auth/usecase"
	"github.com/desulaidovich/pretty-link/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Run() error {

	cfg, err := config.LoadFromEnv()
	if err != nil {
		return err
	}

	db, err := databaseInit(cfg)

	if err != nil {
		return err
	}

	defer db.Close()

	mux := http.NewServeMux()

	authUseCase := usecase.NewAuthUseCase(db)
	api.RegisterAuthEndpoints(mux, authUseCase)

	server := new(http.Server)
	server.Addr = ":" + cfg.Port
	server.Handler = mux

	if err = server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func databaseInit(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.ConnectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
