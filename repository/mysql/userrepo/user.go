package userrepo

import (
	"database/sql"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/entity"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/richerror"
)

func (d *UserDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"
	row := d.db.QueryRow(`
	SELECT id, name, phone_number, created_at, password, role
	FROM users
	WHERE phone_number= ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).
			WithError(err).
			WithMessage(errormessage.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)
	}

	return false, nil
}

func (d *UserDB) Register(user entity.User) (entity.User, error) {
	const op = "mysql.Register"

	result, err := d.db.Exec(`INSERT INTO users(name, phone_number, password, role) VALUES(?, ?, ?)`, user.Name, user.PhoneNumber, user.HashedPassword, user.Role.String())
	if err != nil {

		return entity.User{}, richerror.New(op).
			WithError(err).
			WithMessage("can`t execute command").
			WithKind(richerror.KindUnexpected)
	}

	// error is always nil
	id, _ := result.LastInsertId()

	user.ID = uint(id)

	return user, nil
}

func (d *UserDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.db.QueryRow(`
	SELECT id, name, phone_number, created_at, password, role
	FROM users
	WHERE phone_number= ?`, phoneNumber)

	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).
				WithError(err).
				WithMessage(errormessage.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)

		}

		//TODO - log unexpected error for better observability
		return entity.User{}, richerror.New(op).
			WithError(err).
			WithMessage(errormessage.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)
	}

	return user, nil

}

func (d *UserDB) GetUserByID(userID uint) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.db.QueryRow(`
	SELECT id, name, phone_number, created_at, password, role
	FROM users
	WHERE id = ?
`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).
				WithError(err).
				WithMessage(errormessage.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)

		}

		return entity.User{}, richerror.New(op).
			WithError(err).
			WithMessage(errormessage.ErrorMsgCantScanQuery).
			WithKind(richerror.KindUnexpected)

	}

	return user, nil

}

func scanUser(row *sql.Row) (entity.User, error) {
	user := entity.User{}
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

	user.Role, err = entity.MapStringToRole(role)

	return user, err
}
