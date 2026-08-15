// Package service contains administrative user use cases.
package service

import (
	"context"

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

func (s Service) ListAllUsers(ctx context.Context) ([]userdomain.User, error) {
	return s.repo.ListAll(ctx)
}
