package userservice

import (
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/param"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	const op = "userservice.Login"
	// TODO - it would be better to use two seperate method for existance check and GetUserByPhoneNumber
	user, err := s.Repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, richerror.New(op).WithError(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("phone number or password is not correct")
	}

	accessToken, err := s.Auth.CreateAccessToken(user.ID, user.Role)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshToken, err := s.Auth.CreateRefreshToken(user.ID, user.Role)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.LoginResponse{
		Tokens: param.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		User: param.UserInfo{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}
