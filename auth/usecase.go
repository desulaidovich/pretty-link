package auth

type UseCase interface {
	SignIn(email, password string) (string, error)
	SignUp(email, password string) error
}
