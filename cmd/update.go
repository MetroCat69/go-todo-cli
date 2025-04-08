// cmd/update.go
package cmd

import (
	"fmt"
	"strconv"
	"task-tracker-cli/tasktracker"
	"time"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [task ID] [new description]",
	Short: "Update a task description",
	Args:  cobra.ExactArgs(2),
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

		task, err := tasktracker.FindTaskByID(taskList, id)
		if err != nil {
			fmt.Println(err)
			return
		}

		task.Description = args[1]
		task.UpdatedAt = time.Now()

		err = tasktracker.SaveTasks(taskList)
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}
		fmt.Printf("Task ID %d updated successfully\n", id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
