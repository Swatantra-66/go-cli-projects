package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done|delete]")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add \"task title\"")
			return
		}
		title := os.Args[2]
		tasks = AddTask(tasks, title)
		if err := SaveTasks(tasks); err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
		}
		fmt.Println("Task added.")

	case "list":
		ListTasks(tasks)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		tasks, err = MarkDone(tasks, id)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := SaveTasks(tasks); err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
		}
		fmt.Println("Task marked as done.")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		tasks, err = DeleteTask(tasks, id)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := SaveTasks(tasks); err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
		}
		fmt.Println("Task deleted.")

	default:
		fmt.Println("Unknown command")
	}
}
