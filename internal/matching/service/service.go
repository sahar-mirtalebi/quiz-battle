// Package service contains matchmaking use cases.
package service

import (
	"context"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/domain"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
)

type Repo interface {
	AddToWaitingList(ctx context.Context, userID uint, category quizdomain.Category) error
	GetWaitingListByCategory(ctx context.Context, category quizdomain.Category) ([]domain.WaitingMember, error)
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
