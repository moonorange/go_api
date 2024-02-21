package services

import (
	"context"

	"github.com/moonorange/go_api/gen"
)

type TodoService interface {
	TasksGetAll(ctx context.Context) ([]gen.Task, error)
	TasksCreate(ctx context.Context) (*gen.Task, error)
	TasksDelete(ctx context.Context, id string) error
	TasksRead(ctx context.Context, id string) (*gen.Task, error)
	TasksUpdate(ctx context.Context, id string) (*gen.Task, error)
}

type todoService struct {
	// tqr TODOQueryRepository
	// tcr TODOCommandRepository
}

func (t *todoService) TasksGetAll(ctx context.Context) ([]gen.Task, error) {
	return nil, nil
}

func (t *todoService) TasksCreate(ctx context.Context) (*gen.Task, error) {
	return nil, nil
}

func (t *todoService) TasksDelete(ctx context.Context, id string) error {
	return nil
}

func (t *todoService) TasksRead(ctx context.Context, id string) (*gen.Task, error) {
	return nil, nil
}

func (t *todoService) TasksUpdate(ctx context.Context, id string) (*gen.Task, error) {
	return nil, nil
}
