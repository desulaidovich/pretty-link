package application

import (
	"database/sql"
	"net/http"

	"github.com/desulaidovich/pretty-link/auth/api"
	"github.com/desulaidovich/pretty-link/auth/usecase"
	"github.com/desulaidovich/pretty-link/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Application struct {
	*config.Config
}

func New(options ...func(*Application)) *Application {
	app := new(Application)

	for _, call := range options {
		call(app)
	}
	return app
}

func WithConfig(cfg *config.Config) func(*Application) {
	return func(a *Application) {
		a.Config = cfg
	}
}

func (app *Application) Run() error {
	db, err := sql.Open("pgx", app.ConnectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	defer db.Close()

	mux := http.NewServeMux()

	authUseCase := usecase.NewAuthUseCase(db)
	api.RegisterAuthEndpoints(mux, authUseCase)

	server := &http.Server{
		Addr:    ":" + app.Port,
		Handler: mux,
	}

	return server.ListenAndServe()
}
