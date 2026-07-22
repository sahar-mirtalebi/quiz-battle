package userservice

import (
	"fmt"

	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("error hashing the password")
	}

	user := userdomain.User{
		Name:           req.Name,
		PhoneNumber:    req.PhoneNumber,
		HashedPassword: string(hashedPassword),
		Role:           userdomain.UserRole,
	}

	createdUser, err := s.Repo.Register(user)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	resp := param.RegisterResponse{
		User: param.UserInfo{
			ID:          createdUser.ID,
			Name:        createdUser.Name,
			PhoneNumber: createdUser.PhoneNumber,
		},
	}

	return resp, nil
}
