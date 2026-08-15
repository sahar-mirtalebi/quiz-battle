// Command server starts the quiz game API.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql"
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql/migrator"
	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
	authorizationservice "github.com/sahar-mirtalebi/quiz-battle/internal/auth/authorization"
	accesscontrolrepo "github.com/sahar-mirtalebi/quiz-battle/internal/auth/repository/mysql"
	authservice "github.com/sahar-mirtalebi/quiz-battle/internal/auth/token"
	"github.com/sahar-mirtalebi/quiz-battle/internal/config"
	matchingrepo "github.com/sahar-mirtalebi/quiz-battle/internal/matching/repository/redis"
	matchingservice "github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
	matchingvalidator "github.com/sahar-mirtalebi/quiz-battle/internal/matching/validator"
	httpserver "github.com/sahar-mirtalebi/quiz-battle/internal/server"
	backofficeuserservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/service"
	userrepo "github.com/sahar-mirtalebi/quiz-battle/internal/user/repository/mysql"
	userservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/service"
	uservalidator "github.com/sahar-mirtalebi/quiz-battle/internal/user/validator"
)

func main() {
	cfg := config.Load("config.yml")
	startupCtx, cancelStartup := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelStartup()

	// TODO - add command for migration
	migrator := migrator.New(cfg.Mysql)
	migrator.Up(startupCtx)

	userSvc, backofficeUserSvc, authzSvc, userValidator, matchingValidator, matchingService := setupServices(startupCtx, *cfg)

	server := httpserver.New(*cfg, userSvc, backofficeUserSvc, authzSvc, userValidator, matchingService, matchingValidator)

	server.Serve()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan
	fmt.Println("recieved interrupt signal, sutting down gracefully...")

	shutDownCtx, cancel := context.WithTimeout(
		context.Background(),
		cfg.Application.GracefullShutdownTimeout,
	)
	defer cancel()

	if err := server.Router.Shutdown(shutDownCtx); err != nil {
		fmt.Println("http server shutdown error", err)
	}
}

func setupServices(ctx context.Context, cfg config.Config) (
	userservice.Service,
	backofficeuserservice.Service,
	authorizationservice.Service,
	uservalidator.Validator,
	matchingvalidator.Validator,
	matchingservice.Service) {
	authSvc := authservice.New(cfg.JWT)

	db := mysql.New(ctx, cfg.Mysql)

	userRepo := userrepo.New(db)
	userValidator := uservalidator.New(userRepo)
	userSvc := userservice.New(userRepo, authSvc)

	accessControlRepo := accesscontrolrepo.New(db)
	authzSvc := authorizationservice.New(accessControlRepo)

	backofficeUserSvc := backofficeuserservice.New(userRepo)

	cache := redis.New(ctx, cfg.Redis)

	matchingRepo := matchingrepo.New(cache)
	matchingValidator := matchingvalidator.New()
	matchingService := matchingservice.New(matchingRepo, cfg.MatchingService)

	return userSvc, backofficeUserSvc, authzSvc, userValidator, matchingValidator, matchingService
}
