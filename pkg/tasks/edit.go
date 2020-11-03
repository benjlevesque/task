package tasks

import (
	"fmt"
	"os"
	"strconv"

	"github.com/benjlevesque/task/types"
)

// TaskEditer allows to modify an existing task
type TaskEditer interface {
	EditTask(int, string) error
	GetTask(int) (types.Task, error)
}

// TextEditer allows to modify a string
type TextEditer interface {
	EditText(string) (string, error)
}

// EditTask edits  a task using a TextEditer
func EditTask(store TaskEditer, editer TextEditer, args []string) {

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%s is not a valid id\n", args[0])
		os.Exit(1)
	}
	task, err := store.GetTask(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newText, err := editer.EditText(task.Title)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = store.EditTask(id, newText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Task edited: %d. %s\n", id, newText)
}
