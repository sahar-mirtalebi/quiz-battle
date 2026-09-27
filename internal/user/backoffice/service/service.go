// Package service contains administrative user use cases.
package service

import (
	"context"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/dto"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type Repository interface {
	ListAll(ctx context.Context) ([]userdomain.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) ListAllUsers(ctx context.Context, input dto.ListAllUsersInput) (dto.ListAllUsersOutput, error) {
	users, err := s.repo.ListAll(ctx)
	if err != nil {
		return dto.ListAllUsersOutput{}, richerror.New().WithError(err)
	}
	return dto.ListAllUsersOutput{Users: users}, nil
}
