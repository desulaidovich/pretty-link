package usecase

import (
	"github.com/desulaidovich/pretty-link/auth/repository"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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

func (auth *AuthUseCase) SignIn(email, password string) error {
	_, err := auth.repo.GetByEmail(email)

	if err != nil {
		return err
	}

	return nil
}

func (auth *AuthUseCase) SignUp(email, password string) error {

	passwd := []byte(password)
	passwd, err := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	if err := auth.repo.Create(email, string(passwd)); err != nil {
		return err
	}

	return nil
}
