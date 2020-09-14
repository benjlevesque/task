package util

import (
	"fmt"
	"strconv"

	"github.com/benjlevesque/task/pkg/db"
	"github.com/benjlevesque/task/types"
	"github.com/spf13/cobra"
)

// NoFileCompletion is a stub for ShellCompDirectiveNoFileComp
func NoFileCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveNoFileComp
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func toCompletion(task types.Task) string {
	return fmt.Sprintf("%d\t%s", task.ID, task.Title)
}

// ToCompletionArray returns a bash completion
func toCompletionArray(list []types.Task, args []string) []string {
	ids := make([]string, len(list))
	for i, task := range list {
		if contains(args, strconv.Itoa(task.ID)) {
			continue
		}
		ids[i] = toCompletion(task)
	}
	return ids
}

// GetTaskListValidArgs returns bash completion for a task list
func GetTaskListValidArgs(completionType db.ListType, array bool) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		store := db.GetStore()
		list, err := store.List(completionType)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		var ids []string
		if array || len(args) == 0 {
			ids = toCompletionArray(list, args)
		} else if len(list) > 0 && len(args) == 0 {
			ids = make([]string, 1)
			ids[0] = toCompletion(list[0])
		} else {
			ids = make([]string, 0)
		}
		return ids, cobra.ShellCompDirectiveNoFileComp
	}
}

// GetReminderListValidArgs ...
func GetReminderListValidArgs() func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		store := db.GetStore()
		reminders, err := store.ListReminders()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		validIds := make([]string, len(reminders))
		for i, reminder := range reminders {
			if contains(args, strconv.Itoa(reminder.Task.ID)) {
				continue
			}
			validIds[i] = toCompletion(reminder.Task)
		}
		return validIds, cobra.ShellCompDirectiveNoFileComp
	}
}
