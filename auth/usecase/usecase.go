package usecase

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/desulaidovich/pretty-link/auth/models"
	"github.com/desulaidovich/pretty-link/auth/repository"
)

var (
	errIncorrectUserPassword = errors.New("incorrect user password")
)

type AuthUseCase struct {
	repo *repository.Postgres
	*slog.Logger
}

func NewAuthUseCase(db *sql.DB) func(logger *slog.Logger) *AuthUseCase {
	useCase := new(AuthUseCase)
	postgres := new(repository.Postgres)
	useCase.repo = postgres.New(db)

	return func(logger *slog.Logger) *AuthUseCase {
		useCase.Logger = logger

		return useCase
	}
}

func (auth *AuthUseCase) SignIn(email, password string) error {
	account := new(models.Account)
	account.Email = email

	account, err := auth.repo.GetByEmail(account)

	if err != nil {
		auth.Logger.Error("Failed from SignIn",
			slog.String("Error", err.Error()),
		)
		return err
	}

	if !account.CheckPasswordHash(password) {
		auth.Logger.Error("Error from SignIn",
			slog.String("Message", "incorrect password for "+account.Email),
		)
		return errIncorrectUserPassword
	}

	return nil
}

func (auth *AuthUseCase) SignUp(email, password string) error {
	account := new(models.Account)
	account.Email = email

	if err := account.HashPasswordFromString(password); err != nil {
		auth.Logger.Error("Failed from SignUp",
			slog.String("Message", err.Error()),
		)
		return err
	}

	if err := auth.repo.Create(account); err != nil {
		auth.Logger.Error("Failed from SignUp",
			slog.String("Message", err.Error()),
		)
		return err
	}

	return nil
}
