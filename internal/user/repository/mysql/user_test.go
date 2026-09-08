package mysql

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type userScannerStub struct {
	err  error
	role string
}

func (s userScannerStub) Scan(dest ...any) error {
	if s.err != nil {
		return s.err
	}
	*dest[5].(*string) = s.role
	return nil
}

func TestScanUserPreservesNoRows(t *testing.T) {
	_, err := scanUser(userScannerStub{err: sql.ErrNoRows})
	var re richerror.RichError
	if !errors.As(err, &re) || !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected RichError preserving sql.ErrNoRows, got %v", err)
	}
}

func TestScanUserRole(t *testing.T) {
	user, err := scanUser(userScannerStub{role: userdomain.UserRoleStr})
	if err != nil || user.Role != userdomain.UserRole {
		t.Fatalf("expected valid user and nil error, got %+v, %v", user, err)
	}
	_, err = scanUser(userScannerStub{role: "unknown"})
	var re richerror.RichError
	if !errors.As(err, &re) || re.Kind() != richerror.KindUnexpected {
		t.Fatalf("invalid stored role must be an unexpected RichError, got %v", err)
	}
}
