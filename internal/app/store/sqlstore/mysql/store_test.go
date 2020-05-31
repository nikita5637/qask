package mysql_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = "root:12345678@tcp(172.20.0.5)/qask_test"

	os.Exit(m.Run())
}
