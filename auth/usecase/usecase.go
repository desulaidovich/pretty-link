package usecase

import (
	"database/sql"
	"errors"

	"github.com/desulaidovich/pretty-link/auth/models"
	"github.com/desulaidovich/pretty-link/auth/repository"
	"golang.org/x/crypto/bcrypt"
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

	err = bcrypt.CompareHashAndPassword(account.Password, []byte(password))
	if err != nil {
		return errIncorrectUserPassword
	}

	return nil
}

func (auth *AuthUseCase) SignUp(email, password string) error {
	account := new(models.Account)
	account.Email = email

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	account.Password = hash

	if err := auth.repo.Create(account); err != nil {
		return err
	}

	return nil
}
