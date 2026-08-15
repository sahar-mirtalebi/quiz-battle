package service

import (
	"context"

	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
)

func (s Service) MatchPlayers(ctx context.Context, req dto.MatchPlayersInput) (dto.MatchPlayersOutput, error) {
	// TODO: implement matchmaking logic.
	return dto.MatchPlayersOutput{}, nil
}
