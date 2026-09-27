// Package dto contains matchmaking transport DTOs.
package dto

import (
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/category"
)

type AddToWaitingListInput struct {
	UserID   uint              `json:"-"`
	Category category.Category `json:"category"`
}

type AddToWaitingListOutput struct {
	Timeout time.Duration `json:"timeout"`
}
