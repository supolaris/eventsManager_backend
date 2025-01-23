package models

import (
	"basicapis/db"
	"fmt"
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

func GetSingleEvent(id int64) (*Event, error) {
	getSingleEventQuery := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(getSingleEventQuery, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Time, &event.UserID)
	if err != nil {
		// log.Fatalf("Error in maping rows items", err)
		fmt.Println("error in getting event please check id", err)
		return nil, err
	}
	return &event, nil
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

func (e Event) UpdateEvent() error {
	updateEventQuery := `UPDATE events SET name = ?, description = ?, location = ?, time = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(updateEventQuery)
	if err != nil {
		log.Fatalf("error in updating event", err)
		return err
	}
	defer stmt.Close()
	stmt.Exec(e.Name, e.Description, e.Location, e.Time, e.ID)
	return err
}

func (e Event) DeleteEvent() error {
	deleteQuery := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(deleteQuery)

	if err != nil {
		log.Fatalf("Error in preparing delete query", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}
