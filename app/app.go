package app

import (
	"net/http"

	"github.com/desulaidovich/pretty-link/auth/request"
	"github.com/desulaidovich/pretty-link/auth/usecase"
	"github.com/desulaidovich/pretty-link/config"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Run() error {

	cfg, err := config.LoadFromEnv()

	if err != nil {
		return err
	}

	chi := chi.NewRouter()

	chi.Use(middleware.RequestID)
	chi.Use(middleware.RealIP)
	chi.Use(middleware.Logger)
	chi.Use(middleware.Recoverer)

	auth := usecase.New()
	chi.Mount("/auth", request.RegisterHTTPEndpoints(auth))

	server := new(http.Server)
	server.Addr = ":" + cfg.Port
	server.Handler = chi

	if err = server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
