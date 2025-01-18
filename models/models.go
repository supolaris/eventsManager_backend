package models

import (
	"basicapis/db"
	"log"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Time        time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func GetAllEvents() ([]Event, error) {
	getEventsQuery := `SELECT * FROM events`
	rows, err := db.DB.Query(getEventsQuery)
	if err != nil {
		log.Fatalf("error in getting rows", err)
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Time, &event.UserID)
		if err != nil {
			log.Fatalf("Error in maping rows items", err)
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetSingleEvent(id int64) {
	getSingleEventQuery := "SELECT * FROM evnets WHERE id = ?"

	db.DB.Query(getSingleEventQuery)
}

func (e Event) SaveEvent() error {
	saveEventQuery := `
	INSERT INTO events (name, description, location, time, user_id) VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(saveEventQuery)
	if err != nil {
		log.Fatalf("error in saving event", err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Time, e.UserID)
	if err != nil {
		log.Fatalf("error in executing event query", err)
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}
