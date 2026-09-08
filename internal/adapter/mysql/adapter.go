// Package mysql provides the application's MySQL adapter.
package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	DBName   string `koanf:"db_name"`
}

func New(ctx context.Context, cfg Config) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		panic(richerror.New().WithError(err).WithKind(richerror.KindUnexpected).WithMessage("mysql open err"))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		panic(richerror.New().WithError(err).WithKind(richerror.KindUnexpected).WithMessage("mysql connection failed"))
	}
	return db
}
