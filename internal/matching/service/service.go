// Package matchingservice contains matchmaking use cases.
package matchingservice

import (
	"context"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

type Repo interface {
	AddToWaitingList(ctx context.Context, userID uint, category quizdomain.Category, joinedAt time.Time) error
}

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}
type Service struct {
	repo   Repo
	config Config
}

func New(repo Repo, config Config) Service {
	return Service{
		repo:   repo,
		config: config,
	}
}

func (s Service) AddToWaitingList(ctx context.Context, req param.AddToWaitingListRequest) (param.AddToWaitingListResponse, error) {
	err := s.repo.AddToWaitingList(ctx, req.UserID, req.Category, time.Now())
	if err != nil {
		return param.AddToWaitingListResponse{}, richerror.New().WithError(err)
	}

	return param.AddToWaitingListResponse{Timeout: s.config.WaitingTimeout}, nil
}
