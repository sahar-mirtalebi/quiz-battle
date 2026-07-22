// Package authorizationservice evaluates user permissions.
package authorizationservice

import (
	authdomain "github.com/sahar-mirtalebi/quiz-battle/internal/auth/domain"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type AccessControlRepository interface {
	GetUserPermissions(userID uint, role userdomain.Role) ([]authdomain.PermissionTitle, error)
}

type Service struct {
	accessControlRepo AccessControlRepository
}

func New(accessControlRepo AccessControlRepository) Service {
	return Service{
		accessControlRepo: accessControlRepo,
	}
}

func (s Service) CheckAccess(userID uint, role userdomain.Role, permissions []authdomain.PermissionTitle) (bool, error) {
	userPermissions, err := s.accessControlRepo.GetUserPermissions(userID, role)
	if err != nil {
		return false, err
	}

	granted := make(map[authdomain.PermissionTitle]bool, len(userPermissions))
	for _, permission := range userPermissions {
		granted[permission] = true
	}

	for _, permission := range permissions {
		if !granted[permission] {
			return false, nil
		}
	}

	return true, nil
}
