package tasks

import (
	"fmt"
	"strconv"
)

type taskToggler interface {
	ToggleTask(int, bool) error
}

// ToggleTasks sets one or several task to the value passed as done
func ToggleTasks(store taskToggler, args []string, done bool) {
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%d is not a valid ID", id)
		}
		err = store.ToggleTask(id, done)
		if err != nil {
			fmt.Println(err)
		}
	}
}
