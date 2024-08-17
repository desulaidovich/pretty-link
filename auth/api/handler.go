package api

import (
	"net/http"

	"github.com/desulaidovich/pretty-link/auth"
)

type AuthHandler struct {
	usecase auth.UseCase
}

func New(usecase auth.UseCase) *AuthHandler {
	handler := new(AuthHandler)
	handler.usecase = usecase

	return handler
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if err := h.usecase.SignUp(); err != nil {
		// ok
	}
	w.Write([]byte("signup"))
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if err := h.usecase.SignIn(); err != nil {
		// ok
	}
	w.Write([]byte("signin"))
}
