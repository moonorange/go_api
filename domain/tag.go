package domain

import "github.com/google/uuid"

type Tag struct {
	// Task ID
	ID uuid.UUID `json:"id,omitempty"`
	// Task description
	Name string `json:"name,omitempty"`
}
