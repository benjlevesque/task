package cmd

import (
	"fmt"
	"strconv"

	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/util"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:               "remove",
	Short:             "Removes a task",
	Aliases:           []string{"rm"},
	ValidArgsFunction: util.GetTaskListValidArgs(db.All, true),
	Run:               removeTask,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func removeTask(cmd *cobra.Command, args []string) {
	store := db.GetStore()
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
