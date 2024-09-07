package auth

import "github.com/desulaidovich/pretty-link/internal/fail"

type UseCase interface {
	SignIn(email, password string) *fail.Fail
	SignUp(email, password string) *fail.Fail
}
