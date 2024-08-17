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

type (
	reqeust struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	response struct {
		Message string `json:"message"`
	}
)

func (auth *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		asd, _ := json.Marshal(response{
			Message: err.Error(),
		})
		w.Write(asd)
		return
	}

	var req reqeust
	if err := json.Unmarshal(body, &req); err != nil {
		asd, _ := json.Marshal(response{
			Message: err.Error(),
		})
		w.Write(asd)
		return
	}

	if err := auth.usecase.SignUp(req.Email, req.Password); err != nil {
		asd, _ := json.Marshal(response{
			Message: err.Error(),
		})
		w.Write(asd)
		return
	}

	asd, _ := json.Marshal(response{
		Message: "Ok",
	})
	w.Write(asd)
}

func (auth *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// body, err := io.ReadAll(r.Body)
	// defer r.Body.Close()

	// if err != nil {
	// 	return
	// }

	// var req reqeust
	// if err := json.Unmarshal(body, &req); err != nil {
	// 	return
	// }
}
