package models

import (
	"time"

	"github.com/oscarvo29/todo-list/db"
)

type Task struct {
	Id          int64
	Description string    `binding:"required"`
	CreatedAt   time.Time `binding:"required"`
	Done        bool      `binding:"required"`
}

func (t *Task) TaskDone() {
	t.Done = true
}

func (t *Task) Save() error {
	query := `INSERT INTO tasks(description, createdAt, done) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(t.Description, t.CreatedAt, t.Done)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	t.Id = id
	return nil
}

func GetAllTasks() ([]Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Id, &task.Description, &task.CreatedAt, &task.Done)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
