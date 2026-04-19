package mysql

import (
	"database/sql"
	"fmt"

	"github.com/sahar-mirtalebi/quiz-battle/entity"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number= ?`, phoneNumber)
	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("can`t scan query result %w", err)
	}

	return false, nil
}

func (d *MysqlDB) Register(user entity.User) (entity.User, error) {
	result, err := d.db.Exec(`INSERT INTO users(name, phone_number, password) VALUES(?, ?, ?)`, user.Name, user.PhoneNumber, user.HashedPassword)
	if err != nil {
		return entity.User{}, fmt.Errorf("can`t execute command: %w", err)
	}

	// error is always nil
	id, _ := result.LastInsertId()

	user.ID = uint(id)

	return user, nil
}

func (d *MysqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE phone_number= ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, fmt.Errorf("can`t scan query result %w", err)
	}

	return user, true, nil

}

func (d *MysqlDB) GetUserByID(userID uint) (entity.User, error) {
	row := d.db.QueryRow(`SELECT * FROM users WHERE id= ?`, userID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("record not found")
		}

		return entity.User{}, fmt.Errorf("can`t scan query result %w", err)
	}

	return user, nil

}

func scanUser(row *sql.Row) (entity.User, error) {
	user := entity.User{}
	var createdAt []uint8
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.HashedPassword, &createdAt)

	return user, err
}
