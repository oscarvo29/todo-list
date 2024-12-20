package main

import (
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/oscarvo29/todo-list/db"
	"github.com/oscarvo29/todo-list/models"
	"github.com/oscarvo29/todo-list/utils"
)

var commands = []string{"add", "delete", "list", "complete"}

// To do:
// The rest of the crud

func main() {
	db.InitDB()
	args := os.Args

	if len(args) <= 1 {
		return
	}

	for _, command := range commands {
		if slices.Contains(args, command) {
			switch command {
			case "add":
				task := models.Task{
					Description: "Create the rest of the app",
					CreatedAt:   time.Now(),
					Done:        false,
				}

				err := task.Save()
				if err != nil {
					fmt.Println("Error when trying to save.")
					panic(err)
				}
			case "delete":
				err := models.DeleteTask(int64(1))
				if err != nil {
					fmt.Println("Error when trying to delete task.")
					panic(err)
				}
			case "list":
				tasks, err := models.GetAllTasks()
				if err != nil {
					fmt.Println("Error when trying fetch all tasks.")
					panic(err)
				}
				fmt.Printf("tasks: %v\n", tasks)
			case "complete":
				id, err := utils.ConvertStringToInt64("2")
				if err != nil {
					fmt.Println("Error when trying convert the id. Make sure you type the right id.")
					panic(err)
				}
				task, err := models.GetTaskById(id)
				if err != nil {
					fmt.Println("Error when trying fetch the taskt. Make sure you type the right id.")
					panic(err)
				}

				task.TaskDone()
				err = task.UpdateTask()
				if err != nil {
					fmt.Println("Error when trying to update tasks completion.")
					panic(err)
				}
			}
		}
	}
}
