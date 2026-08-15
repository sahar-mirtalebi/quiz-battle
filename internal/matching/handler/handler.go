// Package handler exposes matchmaking HTTP endpoints.
package handler

import (
	matchingservice "github.com/sahar-mirtalebi/quiz-battle/internal/matching/service"
	matchingvalidator "github.com/sahar-mirtalebi/quiz-battle/internal/matching/validator"
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
