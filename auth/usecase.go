package auth

type UseCase interface {
	SignIn() error
	SignUp() error
}
