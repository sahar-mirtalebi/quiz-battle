package mysql

import (
	"database/sql"
	"time"

	"github.com/sahar-mirtalebi/quiz-battle/entity"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/richerror"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	const op = "mysql.IsPhoneNumberUnique"
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number= ?`, phoneNumber)
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

func (d *MysqlDB) Register(user entity.User) (entity.User, error) {
	const op = "mysql.Register"

	result, err := d.db.Exec(`INSERT INTO users(name, phone_number, password) VALUES(?, ?, ?)`, user.Name, user.PhoneNumber, user.HashedPassword)
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

func (d *MysqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number= ?`, phoneNumber)
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

func (d *MysqlDB) GetUserByID(userID uint) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.db.QueryRow(`SELECT * FROM users WHERE id= ?`, userID)
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
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt, &user.HashedPassword)

	return user, err
}
