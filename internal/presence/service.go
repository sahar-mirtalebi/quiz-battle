package presence

import (
	"context"
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/internal/presence/dto"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

type Config struct {
	PresencePrefix string `koanf:"presence_prefix"`
}

type Repo interface {
	Upsert(ctx context.Context, key string, timestamp int64) error
}

type Service struct {
	Config Config
	Repo   Repo
}

func New(repo Repo) Service {
	return Service{Repo: repo}
}

func (s Service) UpsertPresence(ctx context.Context, req dto.UpsertPresenceRequest) (dto.UpsertPresenceResponse, error) {
	err := s.Repo.Upsert(ctx, fmt.Sprintf("%s:%d", s.Config.PresencePrefix, req.UserID), req.Timestamp)
	if err != nil {
		return dto.UpsertPresenceResponse{}, richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}
	return dto.UpsertPresenceResponse{}, nil
}
