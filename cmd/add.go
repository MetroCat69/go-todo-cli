// cmd/add.go
package cmd

import (
	"fmt"
	"task-tracker-cli/tasktracker" // Import the tasktracker package
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:   "add [task description]",
    Short: "Add a new task",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        taskList, err := tasktracker.LoadTasks() // Call LoadTasks from tasktracker
        if err != nil {
            fmt.Println("Error loading tasks:", err)
            return
        }

        newTask := tasktracker.Task{
            ID:          tasktracker.GenerateNewTaskID(taskList), // Use GenerateNewTaskID
            Description: args[0],
            Status:      "todo",
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        }
        taskList.Tasks = append(taskList.Tasks, newTask)

        err = tasktracker.SaveTasks(taskList) // Call SaveTasks from tasktracker
        if err != nil {
            fmt.Println("Error saving task:", err)
            return
        }
        fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
