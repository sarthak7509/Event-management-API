package models

import (
	"errors"

	"github.com/sarthak7509/event-management/db"
	"github.com/sarthak7509/event-management/utils"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	u.Id, err = result.LastInsertId()
	return err
}

func (u *User) Validate() error {
	query := `
		SELECT password FROM users WHERE email=?
	`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}

	if !utils.CompareHash(u.Password, retrievedPassword) {
		return errors.New("password didnt matched")
	}
	return nil
}
