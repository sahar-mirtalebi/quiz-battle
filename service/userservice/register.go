package userservice

import (
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/entity"
	"github.com/sahar-mirtalebi/quiz-battle/param"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("error hashing the password")
	}

	user := entity.User{
		Name:           req.Name,
		PhoneNumber:    req.PhoneNumber,
		HashedPassword: string(hashedPassword),
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
