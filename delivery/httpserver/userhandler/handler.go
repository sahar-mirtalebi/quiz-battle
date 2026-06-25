package userhandler

import (
	"github.com/sahar-mirtalebi/quiz-battle/service/authservice"
	"github.com/sahar-mirtalebi/quiz-battle/service/userservice"
	"github.com/sahar-mirtalebi/quiz-battle/validator/uservalidator"
)

type Handler struct {
	authSvc       authservice.Service
	userSvc       userservice.Service
	userValidator uservalidator.Validator
	signKey       []byte
}

func New(authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator, signKey string) Handler {
	return Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		userValidator: userValidator,
		signKey:       []byte(signKey),
	}
}
