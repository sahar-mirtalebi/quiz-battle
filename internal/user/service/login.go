package service

import (
	"context"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Login(ctx context.Context, req dto.LoginInput) (dto.LoginOutput, error) {

	// TODO - it would be better to use two seperate method for existance check and GetUserByPhoneNumber
	user, err := s.Repo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return dto.LoginOutput{}, richerror.New().WithError(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return dto.LoginOutput{}, richerror.New().WithError(err).WithMessage("phone number or password is not correct").WithKind(richerror.KindInvalid)
	}

	accessToken, err := s.Auth.CreateAccessToken(user.ID, user.Role)
	if err != nil {
		return dto.LoginOutput{}, richerror.New().WithError(err)
	}

	refreshToken, err := s.Auth.CreateRefreshToken(user.ID, user.Role)
	if err != nil {
		return dto.LoginOutput{}, richerror.New().WithError(err)
	}

	return dto.LoginOutput{
		Tokens: dto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		User: dto.UserInfo{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}
