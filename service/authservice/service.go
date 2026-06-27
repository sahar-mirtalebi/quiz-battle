package authservice

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	SignKey               string        `koanf:"sign_key"`
	AccessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject         string        `koanf:"access_subject"`
	RefreshSubject        string        `koanf:"refresh_subject"`
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}

type Claims struct {
	jwt.RegisteredClaims
	UserID  uint   `json:"user_id"`
	Subject string `json:"subject"`
}

func (s Service) CreateAccessToken(userID uint) (string, error) {

	return s.CreateToken(userID, s.config.AccessExpirationTime, s.config.AccessSubject)
}

func (s Service) CreateRefreshToken(userID uint) (string, error) {

	return s.CreateToken(userID, s.config.RefreshExpirationTime, s.config.RefreshSubject)
}

func (s Service) ParseJWT(bearerToken []string) (*Claims, error) {
	tokenstr := strings.Split(bearerToken[0], " ")[1]

	token, err := jwt.ParseWithClaims(tokenstr, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("userID : %v, expiration : %v", claims.UserID, claims.RegisteredClaims.ExpiresAt)
		return claims, nil
	} else {
		return nil, err
	}

}

func (s Service) CreateToken(userID uint, expirationTime time.Duration, subject string) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		},
		Subject: subject,
		UserID:  userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", fmt.Errorf("unexpected error %w", err)
	}

	return tokenString, nil
}
