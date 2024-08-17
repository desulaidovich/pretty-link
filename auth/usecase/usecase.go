package usecase

import (
	"errors"

	"github.com/desulaidovich/pretty-link/auth/models"
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

func (auth *AuthUseCase) SignIn(email, password string) error {
	account := new(models.Account)
	account.Email = email

	if err := auth.repo.GetByEmail(account); err != nil {
		return err
	}

	if !account.CheckPasswordHash(password) {
		return errors.New("разных хэш крч))")
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
