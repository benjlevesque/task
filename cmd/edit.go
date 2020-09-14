package cmd

import (
	"github.com/benjlevesque/task/cli"
	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/pkg"
	"github.com/benjlevesque/task/util"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:               "edit",
	Short:             "Edits a task",
	Long:              "Opens your default $EDITOR to edit a task title",
	ValidArgsFunction: util.GetTaskListValidArgs(db.Undone, false),
	Run: func(cmd *cobra.Command, args []string) {
		editor := cli.NewTextEditor()
		pkg.EditTask(db.GetStore(), editor, args)
	},
}

func init() {
	editCmd.Flags().BoolVarP(&remind, "remind", "r", false, "Sets a reminder")
	rootCmd.AddCommand(editCmd)
}
