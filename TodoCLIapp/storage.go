package main

import (
	"encoding/json"
	"log"
	"os"
)

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatalf("Unable to marshal struct data due to %v\n", err)
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func LoadTasks() ([]Task, error) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		log.Fatalf("Unable to read file tasks.json due to %v\n", err)
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}
