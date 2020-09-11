package cmd

import (
	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/util"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:               "do",
	Short:             "Marks a task as done",
	ValidArgsFunction: util.GetTaskListValidArgs(db.Undone, true),
	Run: func(cmd *cobra.Command, args []string) {
		toggleTasks(args, true)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
