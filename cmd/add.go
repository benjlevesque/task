package cmd

import (
	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/pkg"
	"github.com/benjlevesque/task/util"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:               "add",
	Short:             "Adds a task",
	ValidArgsFunction: util.NoFileCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.AddTask(db.GetStore(), args)
	},
}

var (
	remind bool
)

func init() {
	addCmd.Flags().BoolVarP(&remind, "remind", "r", false, "Sets a reminder")
	rootCmd.AddCommand(addCmd)
}
