package usecase

import (
	"database/sql"
	"errors"

	"github.com/desulaidovich/pretty-link/auth/models"
	"github.com/desulaidovich/pretty-link/auth/repository"
)

var (
	errIncorrectUserPassword = errors.New("incorrect user password")
)

type AuthUseCase struct {
	repo *repository.Postgres
}

func NewAuthUseCase(db *sql.DB) *AuthUseCase {
	useCase := new(AuthUseCase)
	postgres := new(repository.Postgres)
	useCase.repo = postgres.New(db)

	return useCase
}

func (auth *AuthUseCase) SignIn(email, password string) error {
	account := new(models.Account)
	account.Email = email

	account, err := auth.repo.GetByEmail(account)

	if err != nil {
		return err
	}

	if !account.CheckPasswordHash(password) {
		return errIncorrectUserPassword
	}

	return nil
}

func (auth *AuthUseCase) SignUp(email, password string) error {
	account := new(models.Account)
	account.Email = email

	if err := account.HashPasswordFromString(password); err != nil {
		return err
	}

	if err := auth.repo.Create(account); err != nil {
		return err
	}

	return nil
}
