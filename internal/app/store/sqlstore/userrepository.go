package sqlstore

import "qask/internal/app/model"

//UserRepository is a users sql store
type UserRepository struct {
	store *Store
}

//CreateUser is a function for creating user
func (u *UserRepository) CreateUser(user *model.User) error {
	return nil
}

//FindUserByID is a function for searching user by id
func (u *UserRepository) FindUserByID(ID int) *model.User {
	return nil
}

//FindUserByTgID is a function for searching user by telegram id
func (u *UserRepository) FindUserByTgID(TgID int64) *model.User {
	return nil
}

//FindUserByUserName is a function for searching user by username
func (u *UserRepository) FindUserByUserName(userName string) *model.User {
	return nil
}

//GetUsers returns all system users
func (u *UserRepository) GetUsers() []*model.User {
	return nil
}
