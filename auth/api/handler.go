package api

import (
	"encoding/json"
	"io"
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

type reqeust struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (auth *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var req reqeust
	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := auth.usecase.SignUp(req.Email, req.Password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (auth *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	var req reqeust
	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := auth.usecase.SignIn(req.Email, req.Password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
