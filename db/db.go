package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Error in opening sqlite")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	time DATETIME NOT NULL,
	user_id INTEGER NOT NULL
	)
	`
	_, err := DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error in creating table", err)
	}
}
