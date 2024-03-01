package mysql_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/moonorange/go_api/configs"
	"github.com/moonorange/go_api/infra/mysql"
)

// Ensure the test database can open & close.
func TestDB(t *testing.T) {
	db := MustOpenDB(t)
	MustCloseDB(t, db)
}

// SetupTestDatabase returns a new, open DB and cleanup func to truncated tables passed by args. Fatal on error.
func SetupTestDatabase(t *testing.T, ctx context.Context, tables ...string) (*mysql.DB, func()) {
	t.Helper()

	db := MustOpenDB(t)
	cleanupFunc := func() {
		TruncateTables(t, ctx, db, tables...)
		MustCloseDB(t, db)
	}
	return db, cleanupFunc
}

var registeredDBs = make(map[string]struct{})

func MustOpenDB(t *testing.T) *mysql.DB {
	testDSN := configs.GeTestDSN()
	db := mysql.NewDB(testDSN)
	_, ok := registeredDBs[testDSN]
	if !ok {
		txdb.Register("mydb_test", "mysql", testDSN)
		registeredDBs[testDSN] = struct{}{}
	}
	if err := db.Open(); err != nil {
		t.Fatal(err)
	}
	return db
}

// MustCloseDB closes the DB. Fatal on error.
func MustCloseDB(t *testing.T, db *mysql.DB) {
	t.Helper()

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}

// TruncateTables truncates specified tables in the database.
func TruncateTables(t *testing.T, ctx context.Context, db *mysql.DB, tables ...string) {
	t.Helper()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, table := range tables {
		_, err := tx.Exec("TRUNCATE TABLE " + table)
		if err != nil {
			t.Fatal(err)
		}
	}
}
