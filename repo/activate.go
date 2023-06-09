package repo

import (
	_ "Thanh17b4/practice/model"
	"database/sql"
	"github.com/pkg/errors"
)

type ActivateRepo struct {
	db *sql.DB
}

func NewActivate(db *sql.DB) *ActivateRepo {
	return &ActivateRepo{db: db}
}

func (o *ActivateRepo) Activate(code int, email string) (int, error) {
	newUser, err := o.db.Exec(" UPDATE users SET activated = 1 WHERE email = ? ", email)
	if err != nil {
		return 0, errors.Wrap(err, "could not activate user")
	}
	rowAffect, err := newUser.RowsAffected()
	if err != nil || rowAffect == 0 {
		return 0, err
	}

	return int(rowAffect), nil
}
