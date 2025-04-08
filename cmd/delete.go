// cmd/delete.go
package cmd

import (
	"fmt"
	"strconv"
	"task-tracker-cli/tasktracker"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", args[0])
			return
		}

		taskList, err := tasktracker.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		var taskIndex int
		var taskFound bool
		for i, task := range taskList.Tasks {
			if task.ID == id {
				taskIndex = i
				taskFound = true
				break
			}
		}

		if !taskFound {
			fmt.Println("Task not found")
			return
		}

		taskList.Tasks = append(taskList.Tasks[:taskIndex], taskList.Tasks[taskIndex+1:]...)
		err = tasktracker.SaveTasks(taskList)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}

		fmt.Printf("Task ID %d deleted successfully\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
