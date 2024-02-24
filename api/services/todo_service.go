package services

import (
	"context"

	"github.com/moonorange/go_api/domain"
)

type TodoService interface {
	TasksGetAll(ctx context.Context) ([]*domain.Todo, error)
	TasksCreate(ctx context.Context) (domain.Todo, error)
	TasksDelete(ctx context.Context, id string) error
	TasksRead(ctx context.Context, id string) (domain.Todo, error)
	TasksUpdate(ctx context.Context, id string) (domain.Todo, error)
}
