package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/benjlevesque/task/db"
	"github.com/benjlevesque/task/util"
	"github.com/spf13/cobra"
)

var reminderCommand = &cobra.Command{
	Use:               "reminder",
	Short:             "Create, list, delete reminders",
	ValidArgsFunction: util.GetTaskListValidArgs(db.Undone, false),
	Run:               getReminder,
}

var addReminderCommand = &cobra.Command{
	Use:               "add",
	Short:             "Adds a reminder",
	ValidArgsFunction: util.GetTaskListValidArgs(db.Undone, false),
	Run:               addReminder,
}

var removeReminderCommand = &cobra.Command{
	Use:               "remove",
	Short:             "Deletes a reminder",
	Aliases:           []string{"rm"},
	ValidArgsFunction: util.GetReminderListValidArgs(),
	Run:               deleteReminder,
}

var listReminderCommand = &cobra.Command{
	Use:               "list",
	Short:             "Lists all tasks with overdue reminder",
	Aliases:           []string{"ls"},
	ValidArgsFunction: util.GetTaskListValidArgs(db.Undone, false),
	Run:               dumpReminders,
}

func init() {
	reminderCommand.AddCommand(addReminderCommand)
	reminderCommand.AddCommand(listReminderCommand)
	reminderCommand.AddCommand(removeReminderCommand)
	rootCmd.AddCommand(reminderCommand)
}

func addReminder(cmd *cobra.Command, args []string) {
	id := getReminderID(cmd, args)
	store := db.GetStore()
	err := store.SetReminder(id, time.Now())
	exitIfError(err)
	fmt.Printf("Reminder added for task %d\n", id)
}

func getReminder(cmd *cobra.Command, args []string) {
	id := getReminderID(cmd, args)
	store := db.GetStore()

	t, err := store.GetReminder(id)
	exitIfError(err)
	fmt.Println(t)
}

func dumpReminders(cmd *cobra.Command, args []string) {
	store := db.GetStore()
	reminders, err := store.ListReminders()
	exitIfError(err)
	for _, reminder := range reminders {
		fmt.Printf("%d. %s (%s)\n", reminder.Task.ID, reminder.Task.Title, reminder.Time)
	}
}

func deleteReminder(cmd *cobra.Command, args []string) {
	id := getReminderID(cmd, args)
	store := db.GetStore()
	err := store.DeleteReminder(id)
	exitIfError(err)
}

func getReminderID(cmd *cobra.Command, args []string) int {
	if len(args) == 0 {
		fmt.Println("missing required positional argument [ID]")
		cmd.Help()
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%s is not a valid ID\n", args[0])
		os.Exit(1)
	}
	return id
}
