// Package handler exposes the user HTTP endpoints.
package handler

import (
	userservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/service"
	uservalidator "github.com/sahar-mirtalebi/quiz-battle/internal/user/validator"
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
