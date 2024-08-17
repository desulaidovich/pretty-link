package auth

type UseCase interface {
	SignIn(email, password string) error
	SignUp(email, password string) error
}
