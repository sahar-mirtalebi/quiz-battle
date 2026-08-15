// Package handler exposes administrative user endpoints.
package handler

import (
	authorizationservice "github.com/sahar-mirtalebi/quiz-battle/internal/auth/authorization"
	backofficeuserservice "github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/service"
)

type Handler struct {
	authzSvc          authorizationservice.Service
	backofficeUserSvc backofficeuserservice.Service
	signKey           []byte
}

func New(backofficeUserSvc backofficeuserservice.Service, authzSvc authorizationservice.Service, signKey string) Handler {
	return Handler{
		authzSvc:          authzSvc,
		backofficeUserSvc: backofficeUserSvc,
		signKey:           []byte(signKey),
	}
}
