package entity

import "github.com/google/uuid"

type (
	TaskId      uuid.UUID
	Description string
	IsCompleted bool

	Todo struct {
		ID          TaskId
		Description Description
		Completed   IsCompleted
	}
)
