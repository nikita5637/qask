package mysql

import (
	"database/sql"
	"fmt"
	"testing"
)

//TestDB returns database and teardown function
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	return db, func(tables ...string) {
		if len(tables) > 0 {
			for _, table := range tables {
				db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", table))
			}
		}
		db.Close()
	}
}
