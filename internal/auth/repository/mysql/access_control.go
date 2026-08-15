package mysql

import (
	"context"

	authdomain "github.com/sahar-mirtalebi/quiz-battle/internal/auth/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

func (d *AccessControlDB) GetUserPermissions(ctx context.Context, userID uint, role userdomain.Role) ([]authdomain.PermissionTitle, error) {
	rows, err := d.db.QueryContext(ctx, `
	SELECT DISTINCT p.title
	FROM access_controls ac
	JOIN permissions p ON p.id = ac.permission_id
	WHERE (ac.actor_type = ? AND ac.actor_id = ?)
	   OR (ac.actor_type = ? AND ac.actor_id = ?)
`,
		authdomain.RoleActorType, role,
		authdomain.UserActorType, userID,
	)
	if err != nil {
		return nil, richerror.New().
			WithError(err).
			WithMessage(errormessage.ErrorMsgUnExpected).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	permissionTitles := []authdomain.PermissionTitle{}

	for rows.Next() {
		var permissionTitle authdomain.PermissionTitle
		err := rows.Scan(&permissionTitle)

		if err != nil {
			return nil, richerror.New().
				WithError(err).
				WithMessage(errormessage.ErrorMsgCantScanQuery).
				WithKind(richerror.KindUnexpected)
		}

		permissionTitles = append(permissionTitles, permissionTitle)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New().
			WithError(err).
			WithMessage(errormessage.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)
	}

	return permissionTitles, nil
}
