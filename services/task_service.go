package services

import (
	"context"

	"github.com/moonorange/go_api/models"
)

type TaskService interface {
	TasksGetAll(ctx context.Context) ([]*models.Task, error)
	TasksCreate(ctx context.Context, task *models.Task) error
	TasksDelete(ctx context.Context, id string) error
	TasksRead(ctx context.Context, id string) (*models.Task, error)
	TasksUpdate(ctx context.Context, task *models.Task) error
}
