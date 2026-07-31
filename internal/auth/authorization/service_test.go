package authorizationservice

import (
	"context"
	"testing"

	authdomain "github.com/sahar-mirtalebi/quiz-battle/internal/auth/domain"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type accessRepoStub struct {
	permissions []authdomain.PermissionTitle
}

func (r accessRepoStub) GetUserPermissions(context.Context, uint, userdomain.Role) ([]authdomain.PermissionTitle, error) {
	return r.permissions, nil
}

func TestCheckAccessRequiresAllPermissions(t *testing.T) {
	service := New(accessRepoStub{permissions: []authdomain.PermissionTitle{authdomain.UserListPermission}})

	allowed, err := service.CheckAccess(context.Background(), 1, userdomain.AdminRole, []authdomain.PermissionTitle{
		authdomain.UserListPermission,
		authdomain.UserDeletePermission,
	})
	if err != nil {
		t.Fatal(err)
	}
	if allowed {
		t.Fatal("expected access to be denied when one permission is missing")
	}
}
