package mysql

import (
	"database/sql"
	"fmt"
	"strings"
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

	return db, func(tables ...string) {
		if len(tables) > 0 {
			truncateString := strings.Builder{}
			for _, table := range tables {
				truncateString.WriteString(fmt.Sprintf("TRUNCATE TABLE %s;", table))

			}
			db.Exec(truncateString.String())
		}

		db.Close()
	}
}
