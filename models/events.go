package models

import (
	"time"

	"github.com/sarthak7509/event-management/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() error {
	//save it to database
	query := `
		INSERT INTO events(name,description,location,datetime,user_id)
		values (?, ?, ?, ?, ?)
	` //safe way of using query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId) // does to make impact on data
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT * FROM events
	`
	rows, err := db.DB.Query(query) // only fetch the data
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
