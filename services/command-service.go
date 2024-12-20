package services

import (
	"errors"
)

func ProcessCmd(commands ...string) {
	for idx, cmd := range commands {
		switch cmd {
		case "add":
			description, err := GetNextArg(idx, commands)
			if err != nil {
				panic(err)
			}
			AddTask(description) // TODO : dynamic inp
		case "delete":
			taskId, err := GetNextArg(idx, commands)
			if err != nil {
				panic(err)
			}
			DeleteTask(taskId)
		case "list":
			ListTasks()
		case "complete":
			taskId, err := GetNextArg(idx, commands)
			if err != nil {
				panic(err)
			}
			CompleteTask(taskId)
		}
	}
}

func GetNextArg(idx int, args []string) (string, error) {
	if idx+1 >= len(args) {
		return "", errors.New("no argument at the end of the input")
	}

	return args[idx+1], nil
}