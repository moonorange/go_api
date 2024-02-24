package domain

import "github.com/google/uuid"

type Todo struct {
	// Task ID
	ID uuid.UUID
	// Task description
	Description string
	// boolean value to show if the task is completed or not
	IsCompleted bool
}
