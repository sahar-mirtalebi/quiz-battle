// Package param contains matchmaking transport DTOs.
package param

import (
	"time"

	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
)

type AddToWaitingListRequest struct {
	UserID   uint                `json:"-"`
	Category quizdomain.Category `json:"category"`
}

type AddToWaitingListResponse struct {
	Timeout time.Duration `json:"timeout"`
}
