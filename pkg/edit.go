package pkg

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/benjlevesque/task/types"
)

type taskEditer interface {
	EditTask(int, string) error
	GetTask(int) (types.Task, error)
}

type textEditer interface {
	EditText(string) (string, error)
}

// EditTask edits  a task using a textEditer
func EditTask(store taskEditer, editer textEditer, args []string) {

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%s is not a valid id\n", args[0])
		os.Exit(1)
	}
	task, err := store.GetTask(id)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	newText, err := editer.EditText(task.Title)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = store.EditTask(id, newText)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Task edited: %d. %s\n", id, newText)
}
