package postgres_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = "host=172.20.0.2 user=postgres dbname=qask_test sslmode=disable"

	os.Exit(m.Run())
}
