package postgres

import (
	"qask/internal/app/model"
)

//UserRepository is a users sql store
type UserRepository struct {
	store *Store
}

//CreateUser is a function for creating user
func (u *UserRepository) CreateUser(user *model.User) error {
	if err := u.store.db.QueryRow("INSERT INTO users (username, firstname) VALUES ($1, $2) RETURNING id",
		user.UserName, user.FirstName).Scan(&user.ID); err != nil {
		return err
	}

	return nil
}

//FindUserByID is a function for searching user by id
func (u *UserRepository) FindUserByID(ID int64) *model.User {
	user := &model.User{}

	if err := u.store.db.QueryRow("SELECT id, tgid, username, firstname FROM users WHERE id = $1", ID).Scan(&user.ID, &user.TgID, &user.UserName, &user.FirstName); err != nil {
		return nil
	}

	return user
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
