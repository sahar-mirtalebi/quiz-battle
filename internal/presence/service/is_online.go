package presence

import (
	"context"
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/internal/presence/dto"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func (s Service) presenceKey(userID uint) string {
	return fmt.Sprintf("%s:%d", s.Config.PresencePrefix, userID)
}

// IsOnline returns the last activity timestamps for unexpired presence entries.
func (s Service) IsOnline(ctx context.Context, input dto.IsOnlineInput) (dto.IsOnlineOutput, error) {
	keys := make([]string, len(input.UserIDs))
	for i, userID := range input.UserIDs {
		keys[i] = s.presenceKey(userID)
	}

	existing, err := s.Repo.IsOnline(ctx, keys)
	if err != nil {
		return dto.IsOnlineOutput{}, richerror.New().WithError(err)
	}
	presences := make([]dto.Presence, 0, len(input.UserIDs))
	for i, userID := range input.UserIDs {
		if lastSeen, exists := existing[keys[i]]; exists {
			presences = append(presences, dto.Presence{UserID: userID, Timestamp: lastSeen})
		}
	}
	return dto.IsOnlineOutput{Presences: presences}, nil
}
