package main

import (
	"fmt"
	"os"
	"slices"
)

var commands = []string{"add", "delete", "list", "complete"}

func main() {
	args := os.Args

	if len(args) <= 1 {
		return
	}

	for _, command := range commands {
		if slices.Contains(args, command) {
			switch command {
			case "add":
				fmt.Println("Added")
			case "delete":
				fmt.Println("Deleted")
			case "list":
				fmt.Println("Listing tasks")
			case "complete":
				fmt.Println("task is completed")
			}
		}
	}
}
