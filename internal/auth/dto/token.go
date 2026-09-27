package dto

import (
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
	"time"
)

type CreateAccessTokenInput struct {
	UserID uint
	Role   userdomain.Role
}
type CreateRefreshTokenInput struct {
	UserID uint
	Role   userdomain.Role
}
type CreateTokenInput struct {
	UserID         uint
	Role           userdomain.Role
	ExpirationTime time.Duration
	Subject        string
}

type CreateAccessTokenOutput struct {
	Token string `json:"token"`
}

type CreateRefreshTokenOutput struct {
	Token string `json:"token"`
}

type CreateTokenOutput struct {
	Token string `json:"token"`
}
