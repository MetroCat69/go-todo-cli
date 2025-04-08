// tasktracker/task.go
package tasktracker

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

// Task represents a task structure.
type Task struct {
    ID          int       `json:"id"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}

// TaskList holds all the tasks.
type TaskList struct {
    Tasks []Task `json:"tasks"`
}

// TaskFile represents the name of the task JSON file.
var TaskFile = "tasks.json"

// LoadTasks loads tasks from the JSON file.
func LoadTasks() (TaskList, error) {
    var taskList TaskList
    _, err := os.Stat(TaskFile)
    if os.IsNotExist(err) {
        return taskList, nil // Return empty list if file doesn't exist
    }

    data, err := ioutil.ReadFile(TaskFile)
    if err != nil {
        return taskList, err
    }

    err = json.Unmarshal(data, &taskList)
    if err != nil {
        return taskList, err
    }
    return taskList, nil
}

// SaveTasks saves the tasks back to the JSON file.
func SaveTasks(taskList TaskList) error {
    data, err := json.MarshalIndent(taskList, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(TaskFile, data, 0644)
}

// FindTaskByID finds a task by ID.
func FindTaskByID(taskList TaskList, id int) (*Task, error) {
    for i, task := range taskList.Tasks {
        if task.ID == id {
            return &taskList.Tasks[i], nil
        }
    }
    return nil, errors.New("task not found")
}

// GenerateNewTaskID generates a new ID for a task.
func GenerateNewTaskID(taskList TaskList) int {
    if len(taskList.Tasks) == 0 {
        return 1
    }
    return taskList.Tasks[len(taskList.Tasks)-1].ID + 1
}
