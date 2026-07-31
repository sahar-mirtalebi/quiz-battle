package userrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

func (d *UserDB) IsPhoneNumberUnique(ctx context.Context, phoneNumber string) (bool, error) {

	row := d.db.QueryRowContext(ctx, `
	SELECT id, name, phone_number, created_at, password, role
	FROM users
	WHERE phone_number= ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New().
			WithError(err).
			WithMessage(errormessage.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)
	}

	return false, nil
}

func (d *UserDB) Register(ctx context.Context, user userdomain.User) (userdomain.User, error) {

	result, err := d.db.ExecContext(ctx, `INSERT INTO users(name, phone_number, password, role) VALUES(?, ?, ?, ?)`, user.Name, user.PhoneNumber, user.HashedPassword, user.Role.String())
	if err != nil {

		return userdomain.User{}, richerror.New().
			WithError(err).
			WithMessage("can`t execute command").
			WithKind(richerror.KindUnexpected)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return userdomain.User{}, richerror.New().
			WithError(err).
			WithMessage("can't read created user id").
			WithKind(richerror.KindUnexpected)
	}

	user.ID = uint(id)

	return user, nil
}

func (d *UserDB) ListAll(ctx context.Context) ([]userdomain.User, error) {
	rows, err := d.db.QueryContext(ctx, `
		SELECT id, name, phone_number, created_at, password, role
		FROM users
		ORDER BY id DESC`)
	if err != nil {
		return nil, richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	users := make([]userdomain.User, 0)
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, richerror.New().
				WithError(err).
				WithMessage(errormessage.ErrorMsgCantScanQuery).
				WithKind(richerror.KindUnexpected)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New().WithError(err).WithKind(richerror.KindUnexpected)
	}

	return users, nil
}

func (d *UserDB) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (userdomain.User, error) {

	row := d.db.QueryRowContext(ctx, `
	SELECT id, name, phone_number, created_at, password, role
	FROM users
	WHERE phone_number= ?`, phoneNumber)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return userdomain.User{}, richerror.New().
				WithError(err).
				WithMessage(errormessage.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)

		}

		//TODO - log unexpected error for better observability
		return userdomain.User{}, richerror.New().
			WithError(err).
			WithMessage(errormessage.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)
	}

	return user, nil

}

func (d *UserDB) GetUserByID(ctx context.Context, userID uint) (userdomain.User, error) {

	row := d.db.QueryRowContext(ctx, `
	SELECT id, name, phone_number, created_at, password, role
	FROM users
	WHERE id = ?
`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return userdomain.User{}, richerror.New().
				WithError(err).
				WithMessage(errormessage.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)

		}

		return userdomain.User{}, richerror.New().
			WithError(err).
			WithMessage(errormessage.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)

	}

	return user, nil

}

type scanner interface {
	Scan(dest ...any) error
}

func scanUser(row scanner) (userdomain.User, error) {
	user := userdomain.User{}
	var createdAt time.Time
	var role string

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNumber,
		&createdAt,
		&user.HashedPassword,
		&role,
	)
	if err != nil {
		return user, err
	}

	user.Role, err = userdomain.MapStringToRole(role)

	return user, err
}
