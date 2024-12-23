package main

import (
	"os"

	"github.com/oscarvo29/todo-list/db"
	"github.com/oscarvo29/todo-list/services"
)

/*
	TODO:
	- Don't display tasks which is done.
	- Impletement a flag for "list" commend, to show tasks which is completed.
	- implement an opptional column, for which there is a deadline.
	- Consider to make the table more pretty with column borders, row borders, and so on.
*/

func main() {
	db.InitDB()
	args := os.Args

	if len(args) <= 1 {
		return
	}

	services.ProcessCmd(args...)
}
