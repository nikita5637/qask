package mysql

import (
	"qask/internal/app/model"
)

//UserRepository is a users sql store
type UserRepository struct {
	store *Store
}

//CreateUser is a function for creating user
func (u *UserRepository) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	res, err := u.store.db.Exec("INSERT INTO users (username, firstname, tgid) VALUES (?, ?, ?);", user.UserName, user.FirstName, user.TgID)
	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	user.ID = lastInsertID

	return nil
}

//FindUserByID is a function for searching user by id
func (u *UserRepository) FindUserByID(ID int64) *model.User {
	user := &model.User{}

	if err := u.store.db.QueryRow("SELECT id, tgid, username, firstname FROM users WHERE id = ?", ID).Scan(&user.ID, &user.TgID, &user.UserName, &user.FirstName); err != nil {
		return nil
	}

	return user
}

//FindUserByTgID is a function for searching user by telegram id
func (u *UserRepository) FindUserByTgID(TgID int64) *model.User {
	user := &model.User{}

	if err := u.store.db.QueryRow("SELECT id, username, firstname FROM users WHERE tgid = ?", TgID).Scan(&user.ID, &user.UserName, &user.FirstName); err != nil {
		return nil
	}

	return user
}

//FindUserByUserName is a function for searching user by username
func (u *UserRepository) FindUserByUserName(userName string) *model.User {
	return nil
}

//GetUsers returns all system users
func (u *UserRepository) GetUsers() []*model.User {
	return nil
}
