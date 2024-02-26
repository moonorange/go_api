package mysql

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/moonorange/go_api/domain"
)

// TodoService represents a service for managing dials.
type TodoService struct {
	db *DB
}

// NewTodoService returns a new instance of TodoService.
func NewTODOService(db *DB) *TodoService {
	return &TodoService{db: db}
}

// FindDialByID retrieves a single dial by ID along with associated memberships.
// Only the dial owner & members can see a dial. Returns ENOTFOUND if dial does
// not exist or user does not have permission to view it.
func (s *TodoService) TasksGetAll(ctx context.Context) ([]*domain.Todo, error) {
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
	todo, err := s.listTodos(ctx, tx)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) listTodos(ctx context.Context, tx *Tx) ([]*domain.Todo, error) {

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

	// Iterate over rows and deserialize into Todo objects.
	todos := make([]*domain.Todo, 0)
	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(
			&todo.ID,
			&todo.Description,
			&todo.IsCompleted,
		); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) TasksCreate(ctx context.Context, task *domain.Todo) error {
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
	err = s.createTodo(ctx, tx, task)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *TodoService) createTodo(ctx context.Context, tx *Tx, task *domain.Todo) error {
	task.ID = uuid.New()
	// Execute query
	_, err := tx.ExecContext(ctx, `
		INSERT INTO todos (id, description, is_completed)
		VALUES (?, ?, ?)
	`,
		task.ID, task.Description, task.IsCompleted,
	)
	fmt.Fprintf(os.Stdout, "task\n: %+v", task)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) TasksDelete(ctx context.Context, id string) error {
	return nil
}

func (s *TodoService) TasksRead(ctx context.Context, id string) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (s *TodoService) TasksUpdate(ctx context.Context, task *domain.Todo) error {
	return nil
}
