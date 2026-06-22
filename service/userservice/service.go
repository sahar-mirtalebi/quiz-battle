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
	CreateAccessToken(userID uint) (string, error)
	CreateRefreshToken(userID uint) (string, error)
}

type Service struct {
	Auth AuthService
	Repo Repository
}

func New(repo Repository, auth AuthService) Service {
	return Service{Repo: repo, Auth: auth}
}
