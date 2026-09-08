// Package server wires and runs the HTTP transport.
package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authorizationservice "github.com/sahar-mirtalebi/quiz-battle/internal/auth/authorization"
	"github.com/sahar-mirtalebi/quiz-battle/internal/config"
	matchinghandler "github.com/sahar-mirtalebi/quiz-battle/internal/matching/handler"
	matchingservice "github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
	matchingvalidator "github.com/sahar-mirtalebi/quiz-battle/internal/matching/validator"
	presencemiddleware "github.com/sahar-mirtalebi/quiz-battle/internal/presence/middleware"
	presenceservice "github.com/sahar-mirtalebi/quiz-battle/internal/presence/service"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	backofficeuserhandler "github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/handler"
	backofficeuserservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/service"
	userhandler "github.com/sahar-mirtalebi/quiz-battle/internal/user/handler"
	userservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/service"
	uservalidator "github.com/sahar-mirtalebi/quiz-battle/internal/user/validator"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler
	presenceService       presenceservice.Service

	router *echo.Echo
}

func New(
	config config.Config,
	userSvc userservice.Service,
	backofficeUserSvc backofficeuserservice.Service,
	authzSvc authorizationservice.Service,
	userValidator uservalidator.Validator,
	matchingSvc matchingservice.Service,
	matchingValidator matchingvalidator.Validator,
	presenceService presenceservice.Service,
) *Server {
	return &Server{
		router:                echo.New(),
		config:                config,
		userHandler:           userhandler.New(userSvc, userValidator, config.JWT.SignKey),
		backofficeUserHandler: backofficeuserhandler.New(backofficeUserSvc, authzSvc, config.JWT.SignKey),
		matchingHandler:       matchinghandler.New(matchingSvc, matchingValidator, config.JWT.SignKey),
		presenceService:       presenceService,
	}
}

func (s *Server) Serve() {

	s.router.Use(middleware.Recover())
	s.router.Use(middleware.RequestLogger())

	s.router.GET("/health-check", s.healthCheck)
	upsertPresence := presencemiddleware.Upsert(s.presenceService)
	s.userHandler.SetupRoutes(s.router, upsertPresence)
	s.backofficeUserHandler.SetupRoutes(s.router)
	s.matchingHandler.SetupRoutes(s.router, upsertPresence)

	address := fmt.Sprintf(":%d", s.config.HTTPConfig.Port)

	fmt.Printf("start echo server on %s\n", address)

	go func() {
		if err := s.router.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.router.Logger.Errorf("http server failed: %v", err)
		}
	}()

}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.router.Shutdown(ctx); err != nil {
		return richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
