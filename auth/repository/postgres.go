package repository

import (
	"github.com/desulaidovich/pretty-link/auth/models"
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

func (sql *Postgres) Create(account *models.Account) error {
	rows, err := sql.DB.NamedQuery(`INSERT INTO public.account
		(email, password) VALUES (:email, :password) RETURNING *;`, account)

	if err != nil {
		return err
	}

	for rows.Next() {
		if err = rows.StructScan(account); err != nil {
			return err
		}
	}

	return nil
}

func (sql *Postgres) GetByEmail(account *models.Account) error {
	rows, err := sql.DB.NamedQuery(`SELECT * FROM public.account
		WHERE email=:email;`, &account)

	if err != nil {
		return err
	}

	for rows.Next() {
		if err = rows.StructScan(account); err != nil {
			return err
		}
	}

	return nil
}
