package domain

import (
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/category"
)

type WaitingMember struct {
	UserID uint
	JoinAt time.Time
}

type WaitingPair struct {
	Category category.Category `json:"category"`
	Members  []WaitingMember   `json:"members"`
}
