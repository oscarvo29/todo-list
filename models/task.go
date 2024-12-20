package models

import (
	"time"
)

type Task struct {
	Id          int
	Description string
	CreatedAt   time.Time
	Done        bool
}

func NewTask(description string) *Task {
	return &Task{
		Id:          1,
		Description: description,
		CreatedAt:   time.Now(),
		Done:        false,
	}
}

func (t *Task) TaskDone() {
	t.Done = true
}
