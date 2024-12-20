package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "task.db")
	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(1)
	DB.SetMaxIdleConns(1)

	createTable()
}

func createTable() {
	createTaskTable := `
	CREATE TABLE IF NOT EXISTS task (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		createdAt DATETIME NOT NULL,
		done BOOLEAN NOT NULL
	)`

	_, err := DB.Exec(createTaskTable)
	if err != nil {
		panic("could not create task table")
	}
}
