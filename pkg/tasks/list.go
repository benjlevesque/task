package tasks

import (
	"fmt"

	"github.com/benjlevesque/task/pkg/db"
	"github.com/benjlevesque/task/types"
)

type taskLister interface {
	List(t db.ListType) ([]types.Task, error)
}

// ListTasks list all tasks
func ListTasks(store taskLister, all bool, args []string) {
	listType := db.Undone
	if all {
		listType = db.All
	}
	list, err := store.List(listType)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	if len(list) == 0 {
		fmt.Println("No tasks")
	}

	for _, task := range list {
		if task.Done {
			fmt.Printf("%d. %s ✔\n", task.ID, task.Title)
		} else {
			fmt.Printf("%d. %s\n", task.ID, task.Title)
		}
	}
}
