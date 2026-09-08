// Package service contains user application use cases.
package service

import (
	"context"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type Repository interface {
	Register(ctx context.Context, user userdomain.User) (userdomain.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (userdomain.User, error)
	GetUserByID(ctx context.Context, userID uint) (userdomain.User, error)
}

type AuthService interface {
	CreateAccessToken(userID uint, role userdomain.Role) (string, error)
	CreateRefreshToken(userID uint, role userdomain.Role) (string, error)
}

type Service struct {
	Auth AuthService
	Repo Repository
}

func New(repo Repository, auth AuthService) Service {
	return Service{Repo: repo, Auth: auth}
}

func (s Service) GetUserRole(ctx context.Context, userID uint) (userdomain.Role, error) {
	user, err := s.Repo.GetUserByID(ctx, userID)
	if err != nil {
		return 0, richerror.New().WithError(err)
	}

	return user.Role, nil
}
