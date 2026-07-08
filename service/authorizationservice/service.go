package authorizationservice

import "github.com/sahar-mirtalebi/quiz-battle/entity"

type AccessControlRepository interface {
	GetUserPermissions(userID uint, role entity.Role) ([]entity.PermissionTitle, error)
}

type Service struct {
	accessControlRepo AccessControlRepository
}

func New(accessControlRepo AccessControlRepository) Service {
	return Service{
		accessControlRepo: accessControlRepo,
	}
}

func (s Service) CheckAccess(userID uint, role entity.Role, permissions []entity.PermissionTitle) (bool, error) {
	userPermissions, err := s.accessControlRepo.GetUserPermissions(userID, role)
	if err != nil {
		return false, err
	}

	granted := make(map[entity.PermissionTitle]bool, len(userPermissions))
	for _, permission := range userPermissions {
		granted[permission] = true
	}

	for _, permission := range permissions {
		if granted[permission] {
			return true, nil
		}
	}

	return false, nil
}
