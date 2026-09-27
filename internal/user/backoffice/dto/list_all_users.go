package dto

import userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"

type ListAllUsersInput struct{}

type ListAllUsersOutput struct {
	Users []userdomain.User `json:"users"`
}
