// Package token creates authentication tokens.
package token

import (
	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/dto"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type Config struct {
	SignKey               string        `koanf:"sign_key"`
	AccessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject         string        `koanf:"access_subject"`
	RefreshSubject        string        `koanf:"refresh_subject"`
}

type Service struct {
	Config Config
}

func New(cfg Config) Service {
	return Service{
		Config: cfg,
	}
}

type Claims struct {
	jwt.RegisteredClaims
	UserID uint            `json:"user_id"`
	Role   userdomain.Role `json:"role"`
}

func (s Service) CreateAccessToken(input dto.CreateAccessTokenInput) (dto.CreateAccessTokenOutput, error) {

	output, err := s.CreateToken(dto.CreateTokenInput{UserID: input.UserID, Role: input.Role, ExpirationTime: s.Config.AccessExpirationTime, Subject: s.Config.AccessSubject})
	if err != nil {
		return dto.CreateAccessTokenOutput{}, err
	}
	return dto.CreateAccessTokenOutput{Token: output.Token}, nil
}

func (s Service) CreateRefreshToken(input dto.CreateRefreshTokenInput) (dto.CreateRefreshTokenOutput, error) {

	output, err := s.CreateToken(dto.CreateTokenInput{UserID: input.UserID, Role: input.Role, ExpirationTime: s.Config.RefreshExpirationTime, Subject: s.Config.RefreshSubject})
	if err != nil {
		return dto.CreateRefreshTokenOutput{}, err
	}
	return dto.CreateRefreshTokenOutput{Token: output.Token}, nil
}

func (s Service) CreateToken(input dto.CreateTokenInput) (dto.CreateTokenOutput, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(input.ExpirationTime)),
			Subject:   input.Subject,
		},

		UserID: input.UserID,
		Role:   input.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.Config.SignKey))
	if err != nil {
		return dto.CreateTokenOutput{}, richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}

	return dto.CreateTokenOutput{Token: tokenString}, nil
}
