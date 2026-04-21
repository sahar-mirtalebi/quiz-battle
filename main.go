package main

import (
	"fmt"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/config"
	"github.com/sahar-mirtalebi/quiz-battle/delivery/httpserver"
	"github.com/sahar-mirtalebi/quiz-battle/repository/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/service/authservice"
	"github.com/sahar-mirtalebi/quiz-battle/service/userservice"
)

const (
	JwtSignKey            = "jwt_secret"
	AccessSubject         = "AT"
	RefreshSubject        = "RT"
	AccessExpirationTime  = time.Hour * 24
	RefreshExpirationTime = time.Hour * 24 * 7
)

func main() {
	fmt.Println("start echo server")

	cfg := config.Config{
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessExpirationTime,
			RefreshExpirationTime: RefreshExpirationTime,
			AccessSubject:         AccessSubject,
			RefreshSubject:        RefreshSubject,
		},
		HTTPConfig: config.HTTPConfig{
			Port: 8080,
		},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "gameappt0lk2o20",
			Host:     "localhost",
			Port:     3306,
			DBName:   "gameapp_db",
		},
	}

	// TODO - add command for migration
	// migrator := migrator.New(cfg.Mysql)
	// migrator.Up()

	authSvc, userSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)

	server.Serve()

}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)
	repo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(repo, authSvc)

	return authSvc, userSvc
}
