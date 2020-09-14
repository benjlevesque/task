package cmd

import (
	"github.com/benjlevesque/task/pkg/db"
	"github.com/benjlevesque/task/pkg/tasks"
	"github.com/benjlevesque/task/pkg/util"
	"github.com/spf13/cobra"
)

var all bool

var listCmd = &cobra.Command{
	Use: "list",
	Aliases: []string{
		"ls",
	},
	SuggestFor:        []string{},
	Short:             "Lists all tasks",
	ValidArgsFunction: util.NoFileCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		tasks.ListTasks(db.GetStore(), all, args)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "List tasks, even done")
	rootCmd.AddCommand(listCmd)
}
