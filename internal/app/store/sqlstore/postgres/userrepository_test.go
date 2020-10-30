package postgres_test

import (
	"qask/internal/app/model"
	"qask/internal/app/store/sqlstore/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, tearDown := postgres.TestDB(t, databaseURL)
	defer tearDown("users")

	store := postgres.New(db)
	user := model.TestUser()

	_, err := store.User().CreateUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindUserByID(t *testing.T) {
	db, tearDown := postgres.TestDB(t, databaseURL)
	defer tearDown("users")

	store := postgres.New(db)
	newUser := model.TestUser()

	_, err := store.User().CreateUser(newUser)
	assert.NoError(t, err)
	assert.NotNil(t, newUser)

	user := store.User().FindUserByID(newUser.ID)
	assert.NotNil(t, user)
}
