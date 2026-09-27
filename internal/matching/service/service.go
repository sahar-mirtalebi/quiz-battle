// Package service contains matchmaking use cases.
package service

import (
	"context"
	"time"

	presencedto "github.com/sahar-mirtalebi/quiz-battle/internal/presence/dto"

	"github.com/sahar-mirtalebi/quiz-battle/internal/category"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/domain"
)

type PresenceClient interface {
	IsOnline(ctx context.Context, input presencedto.IsOnlineInput) (presencedto.IsOnlineOutput, error)
}

type Repo interface {
	AddToWaitingList(ctx context.Context, userID uint, category category.Category) error
	GetWaitingListByCategory(ctx context.Context, category category.Category) ([]domain.WaitingMember, error)
}

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}
type Service struct {
	repo           Repo
	config         Config
	presenceClient PresenceClient
}

func New(repo Repo, presenceClient PresenceClient, config Config) Service {
	return Service{
		repo:           repo,
		config:         config,
		presenceClient: presenceClient,
	}
}
