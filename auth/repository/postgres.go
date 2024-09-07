package repository

import (
	"database/sql"

	"github.com/desulaidovich/pretty-link/auth/models"
	"github.com/desulaidovich/pretty-link/internal/fail"
	"github.com/desulaidovich/pretty-link/internal/sqlstate"
)

type Postgres struct {
	*sql.DB
}

func (sql *Postgres) New(db *sql.DB) *Postgres {
	postgres := new(Postgres)
	postgres.DB = db
	return postgres
}

func (p *Postgres) Create(account *models.Account) *fail.Fail {
	err := p.DB.QueryRow(`
		INSERT INTO public.account (email, password)
		VALUES ($1, $2)
		RETURNING id, created_at;
	`, account.Email, account.Password).Scan(&account.ID, &account.CreatedAt)

	if err != nil && err != sql.ErrNoRows {
		return sqlstate.ErrNo(err)
	}

	return nil
}

// TODO: handle error codes
func (p *Postgres) GetByEmail(account *models.Account) (*models.Account, error) {
	acc := new(models.Account)

	rows := p.DB.QueryRow("SELECT * FROM public.account WHERE email=$1;", account.Email)
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	if err := rows.Scan(&acc.ID, &acc.Email, &acc.Password, &acc.CreatedAt); err != nil {
		return nil, err
	}

	return acc, nil
}
