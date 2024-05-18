package services_test

import (
	"context"
	"testing"

	"github.com/moonorange/go_api/infra/mysql"
)

func Test(t *testing.T) {
	var test = mysql.SetupTestDatabase(t, context.Background(), "todos")
	t.Cleanup(test)
}
