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

	_, err := store.User().CreateUser(user)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	//Duplicate username
	_, err = store.User().CreateUser(user)
	assert.Error(t, err)
}

func TestUserRepository_FindUserByID(t *testing.T) {
	db, tearDown := mysql.TestDB(t, databaseURL)
	defer tearDown("users")

	store := mysql.New(db)
	newUser := model.TestUser()

	_, err := store.User().CreateUser(newUser)
	assert.NoError(t, err)
	assert.NotNil(t, newUser)

	user := store.User().FindUserByID(newUser.ID)
	assert.NotNil(t, user)

	user = store.User().FindUserByID(newUser.ID + 1)
	assert.Nil(t, user)
}

func TestUserRepository_FindUserByTgID(t *testing.T) {
	db, tearDown := mysql.TestDB(t, databaseURL)
	defer tearDown("users")

	store := mysql.New(db)
	newUser := model.TestUser()

	_, err := store.User().CreateUser(newUser)
	assert.NoError(t, err)
	assert.NotNil(t, newUser)

	user := store.User().FindUserByTgID(newUser.TgID)
	assert.NotNil(t, user)

	user = store.User().FindUserByTgID(newUser.TgID + 1)
	assert.Nil(t, user)
}
