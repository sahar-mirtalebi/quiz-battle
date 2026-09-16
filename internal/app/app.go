package app

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sahar-mirtalebi/quiz-battle/internal/adapter/mysql"
	redisadapter "github.com/sahar-mirtalebi/quiz-battle/internal/adapter/redis"
	authorizationservice "github.com/sahar-mirtalebi/quiz-battle/internal/auth/authorization"
	accesscontrolrepo "github.com/sahar-mirtalebi/quiz-battle/internal/auth/repository/mysql"
	authservice "github.com/sahar-mirtalebi/quiz-battle/internal/auth/token"
	"github.com/sahar-mirtalebi/quiz-battle/internal/config"
	matchingrepo "github.com/sahar-mirtalebi/quiz-battle/internal/matching/repository/redis"
	matchingservice "github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
	matchingvalidator "github.com/sahar-mirtalebi/quiz-battle/internal/matching/validator"
	presencerepo "github.com/sahar-mirtalebi/quiz-battle/internal/presence/repository/redis"
	presenceservice "github.com/sahar-mirtalebi/quiz-battle/internal/presence/service"
	httpserver "github.com/sahar-mirtalebi/quiz-battle/internal/server"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	backofficeuserservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/service"
	userrepo "github.com/sahar-mirtalebi/quiz-battle/internal/user/repository/mysql"
	userservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/service"
	uservalidator "github.com/sahar-mirtalebi/quiz-battle/internal/user/validator"
)

type App struct {
	Server *httpserver.Server
	MySQL  *sql.DB
	Redis  redisadapter.Adapter
}

func New(
	ctx context.Context,
	cfg config.Config,
) (*App, error) {
	db := mysql.New(ctx, cfg.Mysql)
	cache := redisadapter.New(ctx, cfg.Redis)

	userRepo := userrepo.New(db)
	accessControlRepo := accesscontrolrepo.New(db)
	matchingRepo := matchingrepo.New(cache, cfg.MatchingRepository)
	presenceRepo := presencerepo.New(cache)

	authService := authservice.New(cfg.JWT)
	userService := userservice.New(userRepo, authService)
	backofficeUserService := backofficeuserservice.New(userRepo)
	authorizationService := authorizationservice.New(accessControlRepo)
	matchingService := matchingservice.New(matchingRepo, cfg.MatchingService)
	presenceService := presenceservice.New(cfg.PresenceService, presenceRepo)

	userValidator := uservalidator.New(userRepo)
	matchingValidator := matchingvalidator.New()

	server := httpserver.New(
		cfg,
		userService,
		backofficeUserService,
		authorizationService,
		userValidator,
		matchingService,
		matchingValidator,
		presenceService,
	)

	return &App{
		Server: server,
		MySQL:  db,
		Redis:  cache,
	}, nil
}

func (a *App) Close() error {
	err := errors.Join(
		a.Redis.Close(),
		a.MySQL.Close(),
	)
	if err != nil {
		return richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
