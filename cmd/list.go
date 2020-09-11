package cmd

import (
	"fmt"

	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/util"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Aliases: []string{
		"ls",
	},
	SuggestFor:        []string{},
	Short:             "Lists all tasks",
	ValidArgsFunction: util.NoFileCompletion,
	Run:               listTasks,
}

func init() {
	listCmd.Flags().BoolP("all", "a", false, "List tasks, even done")
	rootCmd.AddCommand(listCmd)
}

func listTasks(cmd *cobra.Command, args []string) {
	store := db.GetStore()
	all := cmd.Flag("all").Value.String()
	listType := db.Undone
	if all == "true" {
		listType = db.All
	}
	list, err := store.List(listType)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	if len(list) == 0 {
		fmt.Println("No tasks")
	}

	for _, task := range list {
		if task.Done {
			fmt.Printf("%d. %s ✔\n", task.ID, task.Title)
		} else {
			fmt.Printf("%d. %s\n", task.ID, task.Title)
		}
	}
}
