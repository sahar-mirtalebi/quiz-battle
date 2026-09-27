package dto

import userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"

type GetUserRoleInput struct {
	UserID uint
}

type GetUserRoleOutput struct {
	Role userdomain.Role `json:"role"`
}
