// Package userservice contains user application use cases.
package userservice

import (
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type Repository interface {
	Register(user userdomain.User) (userdomain.User, error)
	GetUserByPhoneNumber(phoneNumber string) (userdomain.User, error)
	GetUserByID(userID uint) (userdomain.User, error)
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

func (s Service) GetUserRole(userID uint) (userdomain.Role, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return 0, err
	}

	return user.Role, nil
}
