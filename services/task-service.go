package services

import (
	"fmt"
	"time"

	"github.com/oscarvo29/todo-list/models"
	"github.com/oscarvo29/todo-list/utils"
)

func AddTask(description string) {
	task := models.Task{
		Description: description,
		CreatedAt:   time.Now(),
		Done:        false,
	}

	err := task.Save()
	if err != nil {
		fmt.Println("Error when trying to save.")
		panic(err)
	}
}

func DeleteTask(taskId string) {
	id, err := utils.ConvertStringToInt64(taskId)
	if err != nil {
		fmt.Println("Error when trying to convert id. Make sure you type the right input.")
		panic(err)
	}

	err = models.DeleteTask(id)
	if err != nil {
		fmt.Println("Error when trying to delete task.")
		panic(err)
	}
}

func ListTasks() {
	tasks, err := models.GetAllTasks()
	if err != nil {
		fmt.Println("Error when trying fetch all tasks.")
		panic(err)
	}

	table := InstansiateTable(tasks)
	table.PrintTable()
}

func ListTasksWhichIsUncomplete() {
	tasks, err := models.GetAllTasksWhichIsUncomplete()
	if err != nil {
		fmt.Println("Error when trying fetch all tasks.")
		panic(err)
	}

	table := InstansiateTable(tasks)
	table.PrintTable()
}

func CompleteTask(taskId string) {
	id, err := utils.ConvertStringToInt64(taskId)
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
