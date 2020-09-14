package cmd

import (
	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/pkg"
	"github.com/benjlevesque/task/util"
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
		pkg.ListTasks(db.GetStore(), all, args)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "List tasks, even done")
	rootCmd.AddCommand(listCmd)
}
