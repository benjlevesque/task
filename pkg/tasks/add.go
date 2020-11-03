package tasks

import (
	"fmt"
	"strings"
)

type TaskCreater interface {
	CreateTask(title string) (int, error)
}

// AddTask adds a task
func AddTask(store TaskCreater, args []string) {
	title := strings.Join(args, " ")
	title = strings.TrimSpace(title)
	if title == "" {
		fmt.Println("Task is empty")
		return
	}
	id, err := store.CreateTask(title)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("New task created: %d. %s\n", id, title)
}
