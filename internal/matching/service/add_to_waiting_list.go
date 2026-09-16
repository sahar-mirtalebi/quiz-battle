package service

import (
	"context"

	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func (s Service) AddToWaitingList(ctx context.Context, req dto.AddToWaitingListInput) (dto.AddToWaitingListOutput, error) {
	err := s.repo.AddToWaitingList(ctx, req.UserID, req.Category)
	if err != nil {
		return dto.AddToWaitingListOutput{}, richerror.New().WithError(err)
	}

	return dto.AddToWaitingListOutput{Timeout: s.config.WaitingTimeout}, nil
}
