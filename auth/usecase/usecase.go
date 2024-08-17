package usecase

import (
	"github.com/desulaidovich/pretty-link/auth/repository"
	"github.com/jmoiron/sqlx"
)

type AuthUseCase struct {
	repo *repository.Postgres
}

func New(db *sqlx.DB) *AuthUseCase {
	useCase := new(AuthUseCase)
	postgres := new(repository.Postgres)
	useCase.repo = postgres.New(db)

	return useCase
}

func (auth *AuthUseCase) SignIn(email, password string) (string, error) {
	return "", nil
}

func (auth *AuthUseCase) SignUp(email, password string) error {
	if err := auth.repo.CreateAccount(email, password); err != nil {
		return err
	}

	return nil
}
