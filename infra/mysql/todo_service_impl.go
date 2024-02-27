package mysql

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/moonorange/go_api/domain"
)

// TaskService represents a service for managing dials.
type TaskService struct {
	db *DB
}

// NewTaskService returns a new instance of TaskService.
func NewTODOService(db *DB) *TaskService {
	return &TaskService{db: db}
}

// FindDialByID retrieves a single dial by ID along with associated memberships.
// Only the dial owner & members can see a dial. Returns ENOTFOUND if dial does
// not exist or user does not have permission to view it.
func (s *TaskService) TasksGetAll(ctx context.Context) ([]*domain.Task, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := tx.Rollback()
		if err != nil {
			log.Printf("transaction roll back error")
		}
	}()

	// Fetch todo objects
	todo, err := s.listTasks(ctx, tx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TaskService) listTasks(ctx context.Context, tx *Tx) ([]*domain.Task, error) {

	args := []interface{}{}
	// Execute query
	rows, err := tx.QueryContext(ctx, `
		SELECT
		    id,
		    description,
		    is_completed
		FROM todos
		`,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows and deserialize into Task objects.
	tasks := make([]*domain.Task, 0)
	for rows.Next() {
		var todo domain.Task
		if err := rows.Scan(
			&todo.ID,
			&todo.Description,
			&todo.IsCompleted,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, &todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) TasksCreate(ctx context.Context, task *domain.Task) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		err := tx.Rollback()
		if err != nil {
			log.Printf("transaction roll back error")
		}
	}()
	// Create todo objects
	err = s.createTask(ctx, tx, task)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *TaskService) createTask(ctx context.Context, tx *Tx, task *domain.Task) error {
	task.ID = uuid.New()
	// Execute query
	_, err := tx.ExecContext(ctx, `
		INSERT INTO todos (id, description, is_completed)
		VALUES (?, ?, ?)
	`,
		task.ID, task.Description, task.IsCompleted,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) TasksDelete(ctx context.Context, id string) error {
	return nil
}

func (s *TaskService) TasksRead(ctx context.Context, id string) (domain.Task, error) {
	return domain.Task{}, nil
}

func (s *TaskService) TasksUpdate(ctx context.Context, task *domain.Task) error {
	return nil
}
