package app

import (
	"net/http"

	"github.com/desulaidovich/pretty-link/auth/api"
	"github.com/desulaidovich/pretty-link/auth/usecase"
	"github.com/desulaidovich/pretty-link/config"
)

func Run() error {

	cfg, err := config.LoadFromEnv()

	if err != nil {
		return err
	}
	mux := http.NewServeMux()

	auth := usecase.New()
	api.RegisterAuthEndpoints(mux, auth)

	server := new(http.Server)
	server.Addr = ":" + cfg.Port
	server.Handler = mux

	if err = server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
