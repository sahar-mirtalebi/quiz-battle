// Package mysql stores access-control rules in MySQL.
package mysql

import "database/sql"

type AccessControlDB struct {
	db *sql.DB
}

func New(db *sql.DB) *AccessControlDB {
	return &AccessControlDB{db: db}
}
