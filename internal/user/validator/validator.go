// Package uservalidator validates user requests.
package uservalidator

import (
	"context"

	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

const (
	phoneNumberRegex = `^09[0-9]{9}$`
	passwordRegex    = `^[A-Za-z0-9!@#\$%\^&\*]{8,}$`
)

type Repository interface {
	IsPhoneNumberUnique(ctx context.Context, phoneNumber string) (bool, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (userdomain.User, error)
}

type Validator struct {
	Repo Repository
}

func New(repo Repository) Validator {
	return Validator{
		Repo: repo,
	}
}
