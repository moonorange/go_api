package mysql

import (
	"context"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/moonorange/go_api/models"
	"github.com/moonorange/go_api/terrors"
)

// TaskService represents a service for managing dials.
type TaskService struct {
	db *DB
}

// NewTaskService returns a new instance of TaskService.
func NewTaskService(db *DB) *TaskService {
	return &TaskService{db: db}
}

// FindDialByID retrieves a single dial by ID along with associated memberships.
// Only the dial owner & members can see a dial. Returns ENOTFOUND if dial does
// not exist or user does not have permission to view it.
func (s *TaskService) TasksGetAll(ctx context.Context) ([]*models.Task, error) {
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
	todo, err := s.findTasks(ctx, tx, models.TaskFilter{})
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// findTasks retrieves a list of matching tasks. Also returns a total matching
// count which may different from the number of results if filter.Limit is set.
func (s *TaskService) findTasks(ctx context.Context, tx *Tx, filter models.TaskFilter) ([]*models.Task, error) {

	// Build WHERE clause. Each part of the WHERE clause is AND-ed together.
	// Values are appended to an arg list to avoid SQL injection.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}

	// Execute query
	rows, err := tx.QueryContext(ctx, `
		SELECT
		    id,
		    description,
		    is_completed
		FROM todos
		WHERE `+strings.Join(where, " AND ")+`
		ORDER BY id ASC
		`,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows and deserialize into Task objects.
	tasks := make([]*models.Task, 0)
	for rows.Next() {
		var todo models.Task
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

func (s *TaskService) TasksCreate(ctx context.Context, task *models.Task) error {
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

func (s *TaskService) createTask(ctx context.Context, tx *Tx, task *models.Task) error {
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

func (s *TaskService) TasksRead(ctx context.Context, id string) (*models.Task, error) {
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

	// Fetch a task object
	todo, err := s.findTaskByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// findTaskByID is a helper function to retrieve a task by ID.
// Returns ENOTFOUND if task doesn't exist.
func (s *TaskService) findTaskByID(ctx context.Context, tx *Tx, id string) (*models.Task, error) {
	tasks, err := s.findTasks(ctx, tx, models.TaskFilter{ID: &id})
	if err != nil {
		return nil, err
	} else if len(tasks) == 0 {
		return nil, &terrors.Error{Code: terrors.ENOTFOUND, Message: "Task not found."}
	}
	return tasks[0], nil
}

func (s *TaskService) TasksUpdate(ctx context.Context, task *models.Task) error {
	return nil
}
