// Command server starts the quiz game API.
package main

import (
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql/migrator"
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/repository/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/internal/config"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/repository/redis"
	"github.com/sahar-mirtalebi/quiz-battle/internal/server"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/repository/mysql"

	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/authorization"
	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/token"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/validator"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/service"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/service"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/validator"
)

func main() {
	cfg := config.Load("config.yml")

	// TODO - add command for migration
	migrator := migrator.New(cfg.Mysql)
	migrator.Up()

	userSvc, backofficeUserSvc, authzSvc, userValidator, matchingValidator, matchingService := setupServices(*cfg)

	server := httpserver.New(*cfg, userSvc, backofficeUserSvc, authzSvc, userValidator, matchingService, matchingValidator)

	server.Serve()

}

func setupServices(cfg config.Config) (userservice.Service, backofficeuserservice.Service, authorizationservice.Service, uservalidator.Validator, matchingvalidator.Validator, matchingservice.Service) {
	authSvc := authservice.New(cfg.JWT)

	db := mysql.New(cfg.Mysql)

	userRepo := userrepo.New(db)
	userValidator := uservalidator.New(userRepo)
	userSvc := userservice.New(userRepo, authSvc)

	accessControlRepo := accesscontrolrepo.New(db)
	authzSvc := authorizationservice.New(accessControlRepo)

	backofficeUserSvc := backofficeuserservice.New(userRepo)

	cache := redis.New(cfg.Redis)

	matchingRepo := matchingrepo.New(cache)
	matchingValidator := matchingvalidator.New()
	matchingService := matchingservice.New(matchingRepo, cfg.MatchingService)

	return userSvc, backofficeUserSvc, authzSvc, userValidator, matchingValidator, matchingService
}
