package mysql_test

import (
	"qask/internal/app/model"
	"qask/internal/app/store/sqlstore/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportRepository_CreateReport(t *testing.T) {
	db, tearDown := mysql.TestDB(t, databaseURL)
	defer tearDown("users", "reports")

	store := mysql.New(db)
	user := model.TestUser()
	assert.NotNil(t, user)

	err := store.User().CreateUser(user)
	assert.NoError(t, err)

	tests := []struct {
		name        string
		user        int64
		message     string
		wantedError bool
	}{
		{
			name:        "valid",
			user:        user.ID,
			message:     "Report message",
			wantedError: false,
		},
		{
			name:        "not valid (empty message)",
			user:        user.ID,
			message:     "",
			wantedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := store.Report().CreateReport(tt.user, tt.message)
			if tt.wantedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
