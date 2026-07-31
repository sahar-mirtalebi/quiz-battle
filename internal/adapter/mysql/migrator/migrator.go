// Package migrator applies the application's MySQL migrations.
package migrator

import (
	"context"
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql"
)

type Migrator struct {
	Dialect    string
	Config     mysql.Config
	Migrations *migrate.FileMigrationSource
}

func New(conf mysql.Config) Migrator {
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	return Migrator{
		Dialect:    "mysql",
		Migrations: migrations,
		Config:     conf,
	}
}

// TODO - set migration table name
// TODO - add limit to up and down method

func (m Migrator) Up(ctx context.Context) {
	db, err := sql.Open(m.Dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", m.Config.Username, m.Config.Password, m.Config.Host, m.Config.Port, m.Config.DBName))
	if err != nil {
		panic(fmt.Errorf("mysql open err: %v", err))
	}
	defer db.Close()

	n, err := migrate.ExecContext(ctx, db, m.Dialect, m.Migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("cant apply migration: %v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down(ctx context.Context) {
	db, err := sql.Open(m.Dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", m.Config.Username, m.Config.Password, m.Config.Host, m.Config.Port, m.Config.DBName))
	if err != nil {
		panic(fmt.Errorf("mysql open err: %v", err))
	}
	defer db.Close()

	n, err := migrate.ExecContext(ctx, db, m.Dialect, m.Migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("cant rollback migration: %v", err))
	}
	fmt.Printf("Rollbacked %d migrations!\n", n)
}
