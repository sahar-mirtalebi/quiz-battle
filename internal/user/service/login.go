package service

import (
	"context"

	authdto "github.com/sahar-mirtalebi/quiz-battle/internal/auth/dto"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Login(ctx context.Context, req dto.LoginInput) (dto.LoginOutput, error) {
	user, err := s.Repo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return dto.LoginOutput{}, richerror.New().WithError(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return dto.LoginOutput{}, richerror.New().WithError(err).WithMessage("phone number or password is not correct").WithKind(richerror.KindInvalid)
	}

	accessToken, err := s.Auth.CreateAccessToken(authdto.CreateAccessTokenInput{UserID: user.ID, Role: user.Role})
	if err != nil {
		return dto.LoginOutput{}, richerror.New().WithError(err)
	}

	refreshToken, err := s.Auth.CreateRefreshToken(authdto.CreateRefreshTokenInput{UserID: user.ID, Role: user.Role})
	if err != nil {
		return dto.LoginOutput{}, richerror.New().WithError(err)
	}

	return dto.LoginOutput{
		Tokens: dto.Tokens{
			AccessToken:  accessToken.Token,
			RefreshToken: refreshToken.Token,
		},
		User: dto.UserInfo{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}
