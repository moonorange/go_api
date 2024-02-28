package mysql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/moonorange/go_api/infra/mysql"
	"github.com/moonorange/go_api/models"
)

func TestTaskService_TasksCreate(t *testing.T) {
	// Ensure a dial can be created by a user & a membership for the user is automatically created.
	t.Run("OK", func(t *testing.T) {
		db := MustOpenDB(t)
		defer MustCloseDB(t, db)

		ctx := context.Background()

		s := mysql.NewTaskService(db)
		task := &models.Task{Description: "test", IsCompleted: false}

		// Create new dial. Ensure the current user is the owner & an invite code is generated.
		if err := s.TasksCreate(ctx, task); err != nil {
			t.Fatal(err)
		} else if task.ID == (uuid.UUID{}) {
			t.Fatal("expected description")
		} else if got, want := task.Description, "test"; got != want {
			t.Fatal("expected description")
		} else if got, want := task.IsCompleted, false; got != want {
			t.Fatal("expected is_completed")
		}

		// Fetch dial from database & compare.
		if actual, err := s.TasksRead(ctx, task.ID.String()); err != nil {
			t.Fatal(err)
		} else if !reflect.DeepEqual(task, actual) {
			t.Fatalf("mismatch: %#v != %#v", task, actual)
		}
	})
}
