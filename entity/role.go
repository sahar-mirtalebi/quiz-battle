package entity

import "fmt"

type Role uint8

const (
	UserRoleStr  = "user"
	AdminRoleStr = "admin"
)

const (
	UserRole Role = iota + 1
	AdminRole
)

func (r Role) String() string {
	switch r {
	case UserRole:
		return UserRoleStr
	case AdminRole:
		return AdminRoleStr
	}

	return ""
}

func MapStringToRole(role string) (Role, error) {
	switch role {
	case UserRoleStr:
		return UserRole, nil
	case AdminRoleStr:
		return AdminRole, nil
	}

	return 0, fmt.Errorf("invalid role: %s", role)
}
