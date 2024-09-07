package api

import (
	"net/http"

	"github.com/desulaidovich/pretty-link/auth"
	"github.com/desulaidovich/pretty-link/internal/fail"
	"github.com/desulaidovich/pretty-link/internal/render"
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

type response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}

func (auth *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	req, err := render.Bind[reqeust](r)

	if err != nil {
		failed := fail.New(fail.InvalidRequestJSON)

		render.Render(&response{
			Code:    failed.Code(),
			Message: failed.Message(),
		}, failed.HTTPStatusCode(), w)
		return
	}

	if failed := auth.usecase.SignUp(req.Email, req.Password); failed != nil {
		render.Render(&response{
			Code:    failed.Code(),
			Message: failed.Message(),
		}, failed.HTTPStatusCode(), w)
		return
	}

	render.Render(&response{
		Message: "Welcome to the club, buddy",
	}, http.StatusOK, w)
}

func (auth *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	req, err := render.Bind[reqeust](r)

	if err != nil {
		failed := fail.New(fail.InvalidRequestJSON)

		render.Render(&response{
			Code:    failed.Code(),
			Message: failed.Message(),
		}, failed.HTTPStatusCode(), w)
		return
	}

	if failed := auth.usecase.SignIn(req.Email, req.Password); failed != nil {
		render.Render(&response{
			Code:    failed.Code(),
			Message: failed.Message(),
		}, failed.HTTPStatusCode(), w)
		return
	}

	render.Render(&response{
		Token: "Alohomora",
	}, http.StatusOK, w)
}
