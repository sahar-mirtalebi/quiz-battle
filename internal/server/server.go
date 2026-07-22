// Package httpserver wires and runs the HTTP transport.
package httpserver

import (
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/authorization"
	"github.com/sahar-mirtalebi/quiz-battle/internal/config"
	matchinghandler "github.com/sahar-mirtalebi/quiz-battle/internal/matching/handler"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/validator"
	backofficeuserhandler "github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/handler"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/service"
	userhandler "github.com/sahar-mirtalebi/quiz-battle/internal/user/handler"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/service"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler
}

func New(
	config config.Config,
	userSvc userservice.Service,
	backofficeUserSvc backofficeuserservice.Service,
	authzSvc authorizationservice.Service,
	userValidator uservalidator.Validator,
	matchingSvc matchingservice.Service,
	matchingValidator matchingvalidator.Validator,
) Server {
	return Server{
		config:                config,
		userHandler:           userhandler.New(userSvc, userValidator, config.JWT.SignKey),
		backofficeUserHandler: backofficeuserhandler.New(backofficeUserSvc, authzSvc, config.JWT.SignKey),
		matchingHandler:       matchinghandler.New(matchingSvc, matchingValidator, config.JWT.SignKey),
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLogger())

	e.GET("/health-check", s.healthCheck)
	s.userHandler.SetupRoutes(e)
	s.backofficeUserHandler.SetupRoutes(e)
	s.matchingHandler.SetupRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPConfig.Port)))
}
