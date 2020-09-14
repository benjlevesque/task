package cmd

import (
	"github.com/benjlevesque/task/pkg/db"
	"github.com/benjlevesque/task/pkg/tasks"
	"github.com/benjlevesque/task/pkg/util"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:               "do",
	Short:             "Marks a task as done",
	ValidArgsFunction: util.GetTaskListValidArgs(db.Undone, true),
	Run: func(cmd *cobra.Command, args []string) {
		tasks.ToggleTasks(db.GetStore(), args, true)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
