package authservice

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sahar-mirtalebi/quiz-battle/entity"
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
	UserID uint        `json:"user_id"`
	Role   entity.Role `json:"role"`
}

func (s Service) CreateAccessToken(userID uint, role entity.Role) (string, error) {

	return s.CreateToken(userID, role, s.Config.AccessExpirationTime, s.Config.AccessSubject)
}

func (s Service) CreateRefreshToken(userID uint, role entity.Role) (string, error) {

	return s.CreateToken(userID, role, s.Config.RefreshExpirationTime, s.Config.RefreshSubject)
}

func (s Service) CreateToken(userID uint, role entity.Role, expirationTime time.Duration, subject string) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			Subject:   subject,
		},

		UserID: userID,
		Role:   role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.Config.SignKey))
	if err != nil {
		return "", fmt.Errorf("unexpected error %w", err)
	}

	return tokenString, nil
}
