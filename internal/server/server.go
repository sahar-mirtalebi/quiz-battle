// Package server wires and runs the HTTP transport.
package server

import (
	"fmt"

	authorizationservice "github.com/sahar-mirtalebi/quiz-battle/internal/auth/authorization"
	"github.com/sahar-mirtalebi/quiz-battle/internal/config"
	matchinghandler "github.com/sahar-mirtalebi/quiz-battle/internal/matching/handler"
	matchingservice "github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
	matchingvalidator "github.com/sahar-mirtalebi/quiz-battle/internal/matching/validator"
	backofficeuserhandler "github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/handler"
	backofficeuserservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/service"
	userhandler "github.com/sahar-mirtalebi/quiz-battle/internal/user/handler"
	userservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/service"
	uservalidator "github.com/sahar-mirtalebi/quiz-battle/internal/user/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler

	Router *echo.Echo
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
		Router:                echo.New(),
		config:                config,
		userHandler:           userhandler.New(userSvc, userValidator, config.JWT.SignKey),
		backofficeUserHandler: backofficeuserhandler.New(backofficeUserSvc, authzSvc, config.JWT.SignKey),
		matchingHandler:       matchinghandler.New(matchingSvc, matchingValidator, config.JWT.SignKey),
	}
}

func (s Server) Serve() {

	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.RequestLogger())

	s.Router.GET("/health-check", s.healthCheck)
	s.userHandler.SetupRoutes(s.Router)
	s.backofficeUserHandler.SetupRoutes(s.Router)
	s.matchingHandler.SetupRoutes(s.Router)

	address := fmt.Sprintf(":%d", s.config.HTTPConfig.Port)

	fmt.Printf("start echo server on %s\n", address)

	go s.Router.Logger.Fatal(s.Router.Start(address))

}
