package main

import (
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/config"
	"github.com/sahar-mirtalebi/quiz-battle/delivery/httpserver"
	"github.com/sahar-mirtalebi/quiz-battle/repository/migrator"
	"github.com/sahar-mirtalebi/quiz-battle/repository/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/repository/mysql/accesscontrolrepo"
	"github.com/sahar-mirtalebi/quiz-battle/repository/mysql/userrepo"
	"github.com/sahar-mirtalebi/quiz-battle/service/authorizationservice"
	"github.com/sahar-mirtalebi/quiz-battle/service/authservice"
	"github.com/sahar-mirtalebi/quiz-battle/service/backofficeuserservice"
	"github.com/sahar-mirtalebi/quiz-battle/service/userservice"
	"github.com/sahar-mirtalebi/quiz-battle/validator/uservalidator"
)

func main() {
	fmt.Println("start echo server")

	cfg := config.Load("config.yml")
	fmt.Printf("cfg : %+v\n", cfg)

	// TODO - add command for migration
	migrator := migrator.New(cfg.Mysql)
	migrator.Up()

	authSvc, userSvc, backofficeUserSvc, authzSvc, userValidator := setupServices(*cfg)

	server := httpserver.New(*cfg, authSvc, userSvc, backofficeUserSvc, authzSvc, userValidator)

	server.Serve()

}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, backofficeuserservice.Service, authorizationservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.JWT)
	db := mysql.New(cfg.Mysql)
	userRepo := userrepo.New(db)
	accessControlRepo := accesscontrolrepo.New(db)
	userSvc := userservice.New(userRepo, authSvc)
	backofficeUserSvc := backofficeuserservice.New()
	authzSvc := authorizationservice.New(accessControlRepo)

	userValidator := uservalidator.New(userRepo)

	return authSvc, userSvc, backofficeUserSvc, authzSvc, userValidator
}
