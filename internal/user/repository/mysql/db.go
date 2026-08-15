// Package mysql stores users in MySQL.
package mysql

import (
	"database/sql"
)

type UserDB struct {
	db *sql.DB
}

func New(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}
