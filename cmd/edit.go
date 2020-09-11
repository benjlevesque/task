package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/benjlevesque/task/cli"
	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/util"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:               "edit",
	Short:             "Edits a task",
	Long:              "Opens your default $EDITOR to edit a task title",
	ValidArgsFunction: util.GetTaskListValidArgs(db.Undone, false),
	Run:               editTask,
}

func init() {
	editCmd.Flags().BoolVarP(&remind, "remind", "r", false, "Sets a reminder")
	rootCmd.AddCommand(editCmd)
}

func editTask(cmd *cobra.Command, args []string) {
	store := db.GetStore()

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%s is not a valid id\n", args[0])
		os.Exit(1)
	}
	task, err := store.GetTask(id)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	newText, err := cli.EditTextWithEditor(task.Title)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = store.EditTask(id, newText)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Task edited: %d. %s\n", id, newText)
}
