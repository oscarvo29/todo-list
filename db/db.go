package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "tasks.db")
	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(1)
	DB.SetMaxIdleConns(1)

	createTable()
}

func createTable() {
	createTasksTable := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		createdAt DATETIME NOT NULL,
		done BOOLEAN NOT NULL
	)`

	_, err := DB.Exec(createTasksTable)
	if err != nil {
		panic("could not create tasks table")
	}
}
