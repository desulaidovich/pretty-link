package request

import (
	"github.com/desulaidovich/pretty-link/auth"
	"github.com/go-chi/chi/v5"
)

func RegisterHTTPEndpoints(uc auth.UseCase) chi.Router {
	router := chi.NewRouter()
	h := New(uc)

	router.Get("/signin", h.SignIn)
	router.Get("/signup", h.SignUp)

	return router
}
