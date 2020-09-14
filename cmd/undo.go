package cmd

import (
	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/pkg"
	"github.com/benjlevesque/task/util"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:               "undo",
	Short:             "Mark a task as not done",
	ValidArgsFunction: util.GetTaskListValidArgs(db.Done, true),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.ToggleTasks(db.GetStore(), args, false)
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}
