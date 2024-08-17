package api

import (
	"net/http"

	"github.com/desulaidovich/pretty-link/auth"
)

func RegisterAuthEndpoints(mux *http.ServeMux, uc auth.UseCase) {
	handler := New(uc)

	mux.HandleFunc("POST /signin", handler.SignIn)
	mux.HandleFunc("POST /signup", handler.SignUp)
}
