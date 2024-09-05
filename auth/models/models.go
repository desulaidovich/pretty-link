package models

import (
	"time"
)

type Account struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}
