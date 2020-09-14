package cmd

import (
	"github.com/benjlevesque/task/pkg/db"
	"github.com/benjlevesque/task/pkg/tasks"
	"github.com/benjlevesque/task/pkg/util"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:               "remove",
	Short:             "Removes a task",
	Aliases:           []string{"rm"},
	ValidArgsFunction: util.GetTaskListValidArgs(db.All, true),
	Run: func(cmd *cobra.Command, args []string) {
		tasks.RemoveTask(db.GetStore(), args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
