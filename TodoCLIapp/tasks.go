package main

import (
	"fmt"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func AddTask(tasks []Task, title string) []Task {
	newTask := Task{
		ID:    len(tasks) + 1,
		Title: title,
	}
	return append(tasks, newTask)
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet.")
		return
	}
	for _, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
	}
}

func MarkDone(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			return tasks, nil
		}
	}
	return tasks, fmt.Errorf("task with ID %d not found", id)
}

func DeleteTask(tasks []Task, id int) ([]Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}
	return tasks, fmt.Errorf("task with ID %d not found", id)
}
