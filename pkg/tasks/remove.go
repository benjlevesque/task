package tasks

import (
	"fmt"
	"strconv"
)

type TaskDeleter interface {
	DeleteTask(int) error
}

// RemoveTask removes a task
func RemoveTask(store TaskDeleter, args []string) {
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%d is not a valid ID", id)
		}
		err = store.DeleteTask(id)
		if err != nil {
			fmt.Println(err)
		}
	}
}
