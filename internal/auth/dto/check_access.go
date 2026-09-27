package dto

import (
	authdomain "github.com/sahar-mirtalebi/quiz-battle/internal/auth/domain"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type CheckAccessInput struct {
	UserID      uint
	Role        userdomain.Role
	Permissions []authdomain.PermissionTitle
}

type CheckAccessOutput struct {
	Allowed bool `json:"allowed"`
}
