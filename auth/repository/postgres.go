package repository

import (
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	*sqlx.DB
}

func (sql *Postgres) New(db *sqlx.DB) *Postgres {
	postgres := new(Postgres)
	postgres.DB = db

	return postgres
}

func (sql *Postgres) CreateAccount(email, password string) error {
	_, err := sql.DB.Exec(`INSERT INTO public.account (email, password)
		VALUES ($1, $2);`, email, password)

	if err != nil {
		return err
	}

	return nil
}

// func (sql *Postgres) exist(email string) bool {

// }
