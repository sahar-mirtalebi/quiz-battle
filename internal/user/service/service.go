// Package service contains user application use cases.
package service

import (
	"context"
	authdto "github.com/sahar-mirtalebi/quiz-battle/internal/auth/dto"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type Repository interface {
	Register(ctx context.Context, user userdomain.User) (userdomain.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (userdomain.User, error)
	GetUserByID(ctx context.Context, userID uint) (userdomain.User, error)
}

type AuthService interface {
	CreateAccessToken(input authdto.CreateAccessTokenInput) (authdto.CreateAccessTokenOutput, error)
	CreateRefreshToken(input authdto.CreateRefreshTokenInput) (authdto.CreateRefreshTokenOutput, error)
}

type Service struct {
	Auth AuthService
	Repo Repository
}

func New(repo Repository, auth AuthService) Service {
	return Service{Repo: repo, Auth: auth}
}

func (s Service) GetUserRole(ctx context.Context, input dto.GetUserRoleInput) (dto.GetUserRoleOutput, error) {
	user, err := s.Repo.GetUserByID(ctx, input.UserID)
	if err != nil {
		return dto.GetUserRoleOutput{}, richerror.New().WithError(err)
	}

	return dto.GetUserRoleOutput{Role: user.Role}, nil
}
