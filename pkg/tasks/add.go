package tasks

import (
	"fmt"
	"log"
	"strings"
)

type TaskCreater interface {
	CreateTask(title string) (int, error)
}

// AddTask adds a task
func AddTask(store TaskCreater, args []string) {
	title := strings.Join(args, " ")
	id, err := store.CreateTask(title)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("New task created: %d. %s\n", id, title)
}
