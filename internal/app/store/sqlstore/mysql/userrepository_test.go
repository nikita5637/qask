package mysql_test

import (
	"qask/internal/app/model"
	"qask/internal/app/store/sqlstore/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, tearDown := mysql.TestDB(t, databaseURL)
	defer tearDown("users")

	store := mysql.New(db)
	user := model.TestUser()

	err := store.User().CreateUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindUserByID(t *testing.T) {
	db, tearDown := mysql.TestDB(t, databaseURL)
	defer tearDown("users")

	store := mysql.New(db)
	newUser := model.TestUser()

	err := store.User().CreateUser(newUser)
	assert.NoError(t, err)
	assert.NotNil(t, newUser)

	user := store.User().FindUserByID(newUser.ID)
	assert.NotNil(t, user)
}
