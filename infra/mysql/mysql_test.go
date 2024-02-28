package mysql_test

import (
	"testing"

	"github.com/moonorange/go_api/configs"
	"github.com/moonorange/go_api/infra/mysql"
)

// Ensure the test database can open & close.
func TestDB(t *testing.T) {
	db := MustOpenDB(t)
	MustCloseDB(t, db)
}

// MustOpenDB returns a new, open DB. Fatal on error.
func MustOpenDB(tb testing.TB) *mysql.DB {
	tb.Helper()

	testDSN := configs.GeTestDSN()
	db := mysql.NewDB(testDSN)
	if err := db.Open(); err != nil {
		tb.Fatal(err)
	}
	return db
}

// MustCloseDB closes the DB. Fatal on error.
func MustCloseDB(tb testing.TB, db *mysql.DB) {
	tb.Helper()

	if err := db.Close(); err != nil {
		tb.Fatal(err)
	}
}
