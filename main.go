package main

import (
	"os"

	"github.com/oscarvo29/todo-list/db"
	"github.com/oscarvo29/todo-list/services"
)

/*
	TODO:
	- Format everything into a table for printing
	- Format time stamps.
	- Consider the need to make it a loop, so the program wont end after on input, or make multiple tasks in one go
*/

func main() {
	db.InitDB()
	args := os.Args

	if len(args) <= 1 {
		return
	}

	services.ProcessCmd(args...)
}
