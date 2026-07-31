package userservice

import (
	"context"
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Login(ctx context.Context, req param.LoginRequest) (param.LoginResponse, error) {

	// TODO - it would be better to use two seperate method for existance check and GetUserByPhoneNumber
	user, err := s.Repo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, richerror.New().WithError(err)
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
