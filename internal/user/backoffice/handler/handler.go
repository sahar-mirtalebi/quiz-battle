// Package backofficeuserhandler exposes administrative user endpoints.
package backofficeuserhandler

import (
	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/authorization"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/service"
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
