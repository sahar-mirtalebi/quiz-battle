// Package userhandler exposes the user HTTP endpoints.
package userhandler

import (
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/service"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/validator"
)

type Handler struct {
	userSvc       userservice.Service
	userValidator uservalidator.Validator
	signKey       []byte
}

func New(userSvc userservice.Service, userValidator uservalidator.Validator, signKey string) Handler {
	return Handler{
		userSvc:       userSvc,
		userValidator: userValidator,
		signKey:       []byte(signKey),
	}
}
