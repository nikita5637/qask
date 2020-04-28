package teststore

import (
	"qask/internal/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	store := New()

	//Creating a first user
	user1 := &model.User{}
	user1.FirstName = "FirstName_1"
	user1.UserName = "UserName_1"
	user1.TgID = 100

	err := store.User().CreateUser(user1)
	assert.NoError(t, err)

	//Creating a second user with same telegram id
	user2 := &model.User{}
	user2.FirstName = "FirstName_2"
	user2.UserName = "UserName_2"
	user2.TgID = 100

	err = store.User().CreateUser(user2)
	assert.Error(t, err)

	//Creating a third user with same username
	user3 := &model.User{}
	user3.FirstName = "FirstName_3"
	user3.UserName = "UserName_1"
	user3.TgID = 200

	err = store.User().CreateUser(user3)
	assert.Error(t, err)
}

func TestUserRepository_FindUserByID(t *testing.T) {
	store := New()

	testUser := model.TestUser()
	err := store.User().CreateUser(testUser)
	assert.NoError(t, err)

	user := store.User().FindUserByID(1)
	assert.NotNil(t, user)

	user = store.User().FindUserByID(2)
	assert.Nil(t, user)
}

func TestUserRepository_FindUserByTgID(t *testing.T) {
	store := New()

	testUser := model.TestUser()
	err := store.User().CreateUser(testUser)
	assert.NoError(t, err)

	user := store.User().FindUserByTgID(100)
	assert.NotNil(t, user)

	user = store.User().FindUserByTgID(200)
	assert.Nil(t, user)
}

func TestUserRepository_FindUserByUserName(t *testing.T) {
	store := New()

	testUser := model.TestUser()
	err := store.User().CreateUser(testUser)
	assert.NoError(t, err)

	user := store.User().FindUserByUserName("TestUser_UserName")
	assert.NotNil(t, user)

	user = store.User().FindUserByUserName("TestUser_UserName2")
	assert.Nil(t, user)
}
