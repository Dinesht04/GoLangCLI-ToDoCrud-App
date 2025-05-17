package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go add <task> | list | done <id> | delete <id>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task.")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		id := 1
		if len(tasks) > 0 {
			id = tasks[len(tasks)-1].ID + 1
		}
		task := Task{ID: id, Title: title, Done: false}
		tasks = append(tasks, task)
		saveTasks()
		println("Added Task: ", task.Title)
		return

	case "list":
		if len(tasks) == 0 {
			fmt.Println("No Tasks To Display")
			return
		}

		for _, t := range tasks {
			status := "[O]"
			if t.Done {
				status = "[X]"
			}
			fmt.Printf("%d. %s %s\n", t.ID, status, t.Title)
		}
		return

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid Task ID: %s. Please provide a number.", os.Args[2])
			return
		}
		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Done = true
				saveTasks()
				fmt.Println("Marked as Done: ", t.Title)
				return
			}
		}
		fmt.Println("No Task Found with ID: ", os.Args[2])
		return

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid Task ID: %s. Please provide a number.", os.Args[2])
			return
		}
		for i, t := range tasks {
			if t.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				saveTasks()
				fmt.Println("Deleted Task: ", t.Title)
				return
			}
		}
		fmt.Println("No Task Found with ID: ", os.Args[2])
		return

	default:
		fmt.Println("Unknown Command: ", command)
	}
}
