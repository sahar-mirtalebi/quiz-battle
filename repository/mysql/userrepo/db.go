package userrepo

import (
	"database/sql"
)

type UserDB struct {
	db *sql.DB
}

func New(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}
