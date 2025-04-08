// cmd/mark.go
package cmd

import (
	"fmt"
	"strconv"
	"task-tracker-cli/tasktracker"
	"time"

	"github.com/spf13/cobra"
)

// markInProgressCmd marks a task as in progress.
var markInProgressCmd = &cobra.Command{
    Use:   "mark-in-progress [task ID]",
    Short: "Mark a task as in-progress",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // Parse the task ID
        taskID, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Invalid task ID:", err)
            return
        }

        // Load tasks from the JSON file
        taskList, err := tasktracker.LoadTasks()
        if err != nil {
            fmt.Println("Error loading tasks:", err)
            return
        }

        // Find the task by ID
        task, err := tasktracker.FindTaskByID(taskList, taskID)
        if err != nil {
            fmt.Println("Task not found:", err)
            return
        }

        // Update the task status to "in-progress"
        task.Status = "in-progress"
        task.UpdatedAt = time.Now()

        // Save the updated task list back to the JSON file
        err = tasktracker.SaveTasks(taskList)
        if err != nil {
            fmt.Println("Error saving task:", err)
            return
        }

        fmt.Printf("Task (ID: %d) marked as in-progress.\n", task.ID)
    },
}

// markDoneCmd marks a task as done.
var markDoneCmd = &cobra.Command{
    Use:   "mark-done [task ID]",
    Short: "Mark a task as done",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // Parse the task ID
        taskID, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Invalid task ID:", err)
            return
        }

        // Load tasks from the JSON file
        taskList, err := tasktracker.LoadTasks()
        if err != nil {
            fmt.Println("Error loading tasks:", err)
            return
        }

        // Find the task by ID
        task, err := tasktracker.FindTaskByID(taskList, taskID)
        if err != nil {
            fmt.Println("Task not found:", err)
            return
        }

        // Update the task status to "done"
        task.Status = "done"
        task.UpdatedAt = time.Now()

        // Save the updated task list back to the JSON file
        err = tasktracker.SaveTasks(taskList)
        if err != nil {
            fmt.Println("Error saving task:", err)
            return
        }

        fmt.Printf("Task (ID: %d) marked as done.\n", task.ID)
    },
}

func init() {
    // Add the mark commands to the root command
    rootCmd.AddCommand(markInProgressCmd)
    rootCmd.AddCommand(markDoneCmd)
}
