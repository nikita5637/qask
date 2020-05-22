package sqlstore_test

import (
	"qask/internal/app/model"
	"qask/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, tearDown := sqlstore.TestDB(t, databaseURL)
	defer tearDown("users")

	store := sqlstore.New(db)
	user := model.TestUser()

	err := store.User().CreateUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
