package service

import (
	"context"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
)

func (s Service) Profile(ctx context.Context, req dto.ProfileInput) (dto.ProfileOutput, error) {

	user, err := s.Repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return dto.ProfileOutput{}, richerror.New().WithError(err).WithMeta(map[string]interface{}{"req": req})
	}
	return dto.ProfileOutput{Name: user.Name}, nil
}
