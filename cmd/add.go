package cmd

import (
	"fmt"

	"github.com/benjlevesque/task/pkg/cli"
	"github.com/benjlevesque/task/pkg/db"
	"github.com/benjlevesque/task/pkg/tasks"
	"github.com/benjlevesque/task/pkg/util"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:               "add",
	Short:             "Adds a task",
	ValidArgsFunction: util.NoFileCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			editor := cli.NewTextEditor()
			text, err := editor.CaptureInput()
			if err != nil {
				fmt.Println(err)
				return
			}
			args = []string{text}
		}
		tasks.AddTask(db.GetStore(), args)
	},
}

var (
	remind bool
)

func init() {
	addCmd.Flags().BoolVarP(&remind, "remind", "r", false, "Sets a reminder")
	rootCmd.AddCommand(addCmd)
}
