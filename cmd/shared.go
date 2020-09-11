package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/benjlevesque/task/db"
)

func toggleTasks(args []string, done bool) {
	store := db.GetStore()
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

func exitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
