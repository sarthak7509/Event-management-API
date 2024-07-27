package models

import (
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
