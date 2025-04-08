// cmd/list.go
package cmd

import (
	"fmt"
	"task-tracker-cli/tasktracker"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List tasks",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskList, err := tasktracker.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		var filteredTasks []tasktracker.Task
		if len(args) == 0 {
			filteredTasks = taskList.Tasks
		} else {
			status := args[0]
			for _, task := range taskList.Tasks {
				if task.Status == status {
					filteredTasks = append(filteredTasks, task)
				}
			}
		}

		for _, task := range filteredTasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s\n", task.ID, task.Description, task.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
