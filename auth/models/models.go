package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

func (account *Account) HashPasswordFromString(password string) error {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	account.Password = p

	return nil
}

func (account *Account) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword(account.Password, []byte(password))

	return err == nil
}
