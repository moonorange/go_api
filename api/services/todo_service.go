package services

import (
	"context"

	"github.com/moonorange/go_api/domain"
)

type TaskService interface {
	TasksGetAll(ctx context.Context) ([]*domain.Task, error)
	TasksCreate(ctx context.Context, task *domain.Task) error
	TasksDelete(ctx context.Context, id string) error
	TasksRead(ctx context.Context, id string) (domain.Task, error)
	TasksUpdate(ctx context.Context, task *domain.Task) error
}
