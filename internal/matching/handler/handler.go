// Package matchinghandler exposes matchmaking HTTP endpoints.
package matchinghandler

import (
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"

	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/validator"
)

type Handler struct {
	matchingSvc       matchingservice.Service
	matchingValidator matchingvalidator.Validator
	signKey           []byte
}

func New(matchingSvc matchingservice.Service, matchingValidator matchingvalidator.Validator, signKey string) Handler {
	return Handler{
		matchingSvc:       matchingSvc,
		matchingValidator: matchingValidator,
		signKey:           []byte(signKey),
	}
}
