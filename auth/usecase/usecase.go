package usecase

import (
	"database/sql"

	"github.com/desulaidovich/pretty-link/auth/models"
	"github.com/desulaidovich/pretty-link/auth/repository"
	"github.com/desulaidovich/pretty-link/internal/fail"
	"golang.org/x/crypto/bcrypt"
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

func (auth *AuthUseCase) SignIn(email, password string) *fail.Fail {
	account := new(models.Account)
	account.Email = email

	account, err := auth.repo.GetByEmail(account)
	if err != nil {
		return fail.New(fail.AccountIsNotExists)
	}

	err = bcrypt.CompareHashAndPassword(account.Password, []byte(password))
	if err != nil {
		return fail.New(fail.IncorrectUserPassword)
	}

	return nil
}

func (auth *AuthUseCase) SignUp(email, password string) *fail.Fail {
	account := new(models.Account)
	account.Email = email

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fail.New(fail.InvalidUserPassword)
	}

	account.Password = hash

	if failed := auth.repo.Create(account); failed != nil {
		return failed
	}

	return nil
}
