package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/util"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:               "add",
	Short:             "Adds a task",
	ValidArgsFunction: util.NoFileCompletion,
	Run:               addTask,
}

var (
	remind bool
)

func init() {
	addCmd.Flags().BoolVarP(&remind, "remind", "r", false, "Sets a reminder")
	rootCmd.AddCommand(addCmd)
}

func addTask(cmd *cobra.Command, args []string) {
	store := db.GetStore()

	title := strings.Join(args, " ")
	id, err := store.CreateTask(title)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("New task created: %d. %s\n", id, title)
}
