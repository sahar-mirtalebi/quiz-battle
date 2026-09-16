package dto

import "github.com/sahar-mirtalebi/quiz-battle/internal/matching/domain"

type MatchPlayersInput struct {
}

type MatchPlayersOutput struct {
	Pairs []domain.WaitingPair `json:"pairs"`
}
