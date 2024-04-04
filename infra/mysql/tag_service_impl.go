package mysql

import (
	"context"

	"github.com/moonorange/go_api/domain"
)

// TaskService represents a service for managing dials.
type TagService struct {
	db *DB
}

// NewTaskService returns a new instance of TaskService.
func NewTagService(db *DB) *TagService {
	return &TagService{db: db}
}

func (t *TagService) ListTags(ctx context.Context) ([]*domain.Tag, error) {
	return nil, nil
}
