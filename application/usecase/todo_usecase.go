package usecase

import (
	"context"

	"github.com/moonorange/go_api/api/gen"
	"github.com/moonorange/go_api/domain/entity"
)

type TodoUseCase interface {
	TasksGetAll(ctx context.Context) ([]gen.Task, error)
	TasksCreate(ctx context.Context) (*gen.Task, error)
	TasksDelete(ctx context.Context, id string) error
	TasksRead(ctx context.Context, id string) (*gen.Task, error)
	TasksUpdate(ctx context.Context, id string) (*gen.Task, error)
}

type todoUseCase struct {
	// tqr TODOQueryRepository
	// tcr TODOCommandRepository
}

func (t *todoUseCase) TasksGetAll(ctx context.Context) ([]entity.Todo, error) {
	return nil, nil
}

func (t *todoUseCase) TasksCreate(ctx context.Context) (*entity.Todo, error) {
	return nil, nil
}

func (t *todoUseCase) TasksDelete(ctx context.Context, id string) error {
	return nil
}

func (t *todoUseCase) TasksRead(ctx context.Context, id string) (*entity.Todo, error) {
	return nil, nil
}

func (t *todoUseCase) TasksUpdate(ctx context.Context, id string) (*entity.Todo, error) {
	return nil, nil
}
