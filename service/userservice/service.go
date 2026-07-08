package userservice

import (
	"github.com/sahar-mirtalebi/quiz-battle/entity"
)

type Repository interface {
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(userID uint) (entity.User, error)
}

type AuthService interface {
	CreateAccessToken(userID uint, role entity.Role) (string, error)
	CreateRefreshToken(userID uint, role entity.Role) (string, error)
}

type Service struct {
	Auth AuthService
	Repo Repository
}

func New(repo Repository, auth AuthService) Service {
	return Service{Repo: repo, Auth: auth}
}

func (s Service) GetUserRole(userID uint) (entity.Role, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return 0, err
	}

	return user.Role, nil
}
