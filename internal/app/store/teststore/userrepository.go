package teststore

import (
	"qask/internal/app/model"
	"qask/internal/app/store"
)

//UserRepository is a users store for testing
type UserRepository struct {
	users map[int]*model.User
}

//CreateUser is a function for creating user
func (u *UserRepository) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if tmpUser := u.findUser(user); tmpUser != nil {
		return store.ErrUserExists
	}

	user.ID = len(u.users) + 1
	u.users[user.ID] = user

	return nil
}

func (u *UserRepository) findUser(user *model.User) *model.User {
	if tmpUser := u.FindUserByTgID(user.TgID); tmpUser != nil {
		return tmpUser
	}
	if tmpUser := u.FindUserByUserName(user.UserName); tmpUser != nil {
		return tmpUser
	}

	return nil
}

//FindUserByTgID is a function for searching user by telegram id
func (u *UserRepository) FindUserByTgID(tgID int64) *model.User {
	for _, u := range u.users {
		if u.TgID == tgID {
			return u
		}
	}

	return nil
}

//FindUserByUserName is a function for searching user by username
func (u *UserRepository) FindUserByUserName(userName string) *model.User {
	for _, u := range u.users {
		if u.UserName == userName {
			return u
		}
	}

	return nil
}

//GetUsers returns all system users
func (u *UserRepository) GetUsers() []*model.User {
	users := make([]*model.User, len(u.users))
	i := 0
	for _, user := range u.users {
		users[i] = user
		i++
	}

	return users
}
