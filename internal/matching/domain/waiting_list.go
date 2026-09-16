package domain

import (
	"time"

	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
)

type WaitingMember struct {
	UserID uint
	JoinAt time.Time
}

type WaitingPair struct {
	Category quizdomain.Category `json:"category"`
	Members  []WaitingMember     `json:"members"`
}
