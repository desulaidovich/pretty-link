package api

import (
	"net/http"

	"github.com/desulaidovich/pretty-link/auth"
)

func RegisterAuthEndpoints(mux *http.ServeMux, uc auth.UseCase) {
	h := New(uc)

	mux.HandleFunc("/signin", h.SignIn)
	mux.HandleFunc("/signup", h.SignUp)
}
