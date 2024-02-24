package mysql

import (
	"context"

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
func (s *TodoService) TasksGetAll(ctx context.Context, id int) ([]*domain.Todo, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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
