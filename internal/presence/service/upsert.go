package presence

import (
	"context"
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/internal/presence/dto"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func (s Service) Upsert(ctx context.Context, input dto.UpsertPresenceInput) (dto.UpsertPresenceOutput, error) {
	err := s.Repo.Upsert(ctx, fmt.Sprintf("%s:%d", s.Config.PresencePrefix, input.UserID), input.Timestamp, s.Config.ExpirationTime)
	if err != nil {
		return dto.UpsertPresenceOutput{}, richerror.New().WithError(err)
	}
	return dto.UpsertPresenceOutput{}, nil
}
