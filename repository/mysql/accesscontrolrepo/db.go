package accesscontrolrepo

import "database/sql"

type AccessControlDB struct {
	db *sql.DB
}

func New(db *sql.DB) *AccessControlDB {
	return &AccessControlDB{db: db}
}
