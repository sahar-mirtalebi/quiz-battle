package accesscontrolrepo

import (
	"github.com/sahar-mirtalebi/quiz-battle/entity"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/richerror"
)

func (d *AccessControlDB) GetUserPermissions(userID uint, role entity.Role) ([]entity.PermissionTitle, error) {
	const op = "mysql.GetUserACL"

	rows, err := d.db.Query(`
	SELECT DISTINCT p.title
	FROM access_controls ac
	JOIN permissions p ON p.id = ac.permission_id
	WHERE (ac.actor_type = ? AND ac.actor_id = ?)
	   OR (ac.actor_type = ? AND ac.actor_id = ?)
`,
		entity.RoleActorType, role,
		entity.UserActorType, userID,
	)
	if err != nil {
		return nil, richerror.New(op).
			WithError(err).
			WithMessage(errormessage.ErrorMsgUnExpected).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	permissionTitles := []entity.PermissionTitle{}

	for rows.Next() {
		var permissionTitle entity.PermissionTitle
		err := rows.Scan(&permissionTitle)

		if err != nil {
			return nil, richerror.New(op).
				WithError(err).
				WithMessage(errormessage.ErrorMsgCantScanQuery).
				WithKind(richerror.KindUnexpected)
		}

		permissionTitles = append(permissionTitles, permissionTitle)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).
			WithError(err).
			WithMessage(errormessage.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)
	}

	return permissionTitles, nil
}
