package pkg

import (
	"fmt"
	"log"
	"strings"
)

type taskCreater interface {
	CreateTask(title string) (int, error)
}

// AddTask adds a task
func AddTask(store taskCreater, args []string) {
	title := strings.Join(args, " ")
	id, err := store.CreateTask(title)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("New task created: %d. %s\n", id, title)
}
