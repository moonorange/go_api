package domain

import "github.com/google/uuid"

type Task struct {
	// Task ID
	ID uuid.UUID `json:"id,omitempty"`
	// Task description
	Description string `json:"description,omitempty"`
	// boolean value to show if the task is completed or not
	IsCompleted bool `json:"completed,omitempty"`
}

func (t *Task) CompleteTask() {
	t.IsCompleted = true
}
