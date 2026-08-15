// Package service contains matchmaking use cases.
package service

import (
	"context"
	"time"

	dto "github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

type Repo interface {
	AddToWaitingList(ctx context.Context, userID uint, category quizdomain.Category) error
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

func (s Service) AddToWaitingList(ctx context.Context, req dto.AddToWaitingListInput) (dto.AddToWaitingListOutput, error) {
	err := s.repo.AddToWaitingList(ctx, req.UserID, req.Category)
	if err != nil {
		return dto.AddToWaitingListOutput{}, richerror.New().WithError(err)
	}

	return dto.AddToWaitingListOutput{Timeout: s.config.WaitingTimeout}, nil
}
