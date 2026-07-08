package httpserver

import (
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/config"
	backofficeuserhandler "github.com/sahar-mirtalebi/quiz-battle/delivery/httpserver/backofficeuserhandler"
	userhandler "github.com/sahar-mirtalebi/quiz-battle/delivery/httpserver/userhandler"
	"github.com/sahar-mirtalebi/quiz-battle/service/authorizationservice"
	"github.com/sahar-mirtalebi/quiz-battle/service/authservice"
	"github.com/sahar-mirtalebi/quiz-battle/service/backofficeuserservice"
	"github.com/sahar-mirtalebi/quiz-battle/service/userservice"
	"github.com/sahar-mirtalebi/quiz-battle/validator/uservalidator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service, backofficeUserSvc backofficeuserservice.Service, authzSvc authorizationservice.Service, userValidator uservalidator.Validator) Server {
	return Server{
		config:                config,
		userHandler:           userhandler.New(authSvc, userSvc, userValidator, config.JWT.SignKey),
		backofficeUserHandler: backofficeuserhandler.New(backofficeUserSvc, authzSvc, config.JWT.SignKey),
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLogger())

	e.GET("/health-check", s.healthCheck)
	s.userHandler.SetUpUserRoute(e)
	s.backofficeUserHandler.SetUpBackofficeUserRoute(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPConfig.Port)))
}
