package main

import (
	"fmt"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/config"
	"github.com/sahar-mirtalebi/quiz-battle/delivery/httpserver"
	"github.com/sahar-mirtalebi/quiz-battle/repository/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/service/authservice"
	"github.com/sahar-mirtalebi/quiz-battle/service/userservice"
	"github.com/sahar-mirtalebi/quiz-battle/validator/uservalidator"
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

	cfg := config.Load()
	fmt.Printf("cfg : %+v\n", cfg)

	// TODO - add command for migration
	// migrator := migrator.New(cfg.Mysql)
	// migrator.Up()

	authSvc, userSvc, userValidator := setupServices(*cfg)

	server := httpserver.New(*cfg, authSvc, userSvc, userValidator)

	server.Serve()

}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.JWT)
	repo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(repo, authSvc)

	userValidator := uservalidator.New(repo)

	return authSvc, userSvc, userValidator
}
