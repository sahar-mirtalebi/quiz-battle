// Package dto contains matchmaking transport DTOs.
package dto

import (
	"time"

	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
)

type AddToWaitingListInput struct {
	UserID   uint                `json:"-"`
	Category quizdomain.Category `json:"category"`
}

type AddToWaitingListOutput struct {
	Timeout time.Duration `json:"timeout"`
}
