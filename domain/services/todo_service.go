package services

import (
	"github.com/moonorange/go_api/domain"
)

type TodoService interface {
	TasksGetAll() ([]domain.Todo, error)
	TasksCreate() (*domain.Todo, error)
	TasksDelete(id domain.TaskId) error
	TasksRead(id domain.TaskId) (*domain.Todo, error)
	TasksUpdate(id domain.TaskId) (*domain.Todo, error)
}
