package mysql

import (
	"fmt"
	"qask/internal/app/model"
)

//UserRepository is a users sql store
type UserRepository struct {
	store *Store
}

//CreateUser is a function for creating user
func (u *UserRepository) CreateUser(user *model.User) error {
	res, err := u.store.db.Exec("INSERT INTO users (username, firstname) VALUES (?, ?);",
		user.UserName, user.FirstName)
	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	user.ID = int(lastInsertID)

	return nil
}

//FindUserByID is a function for searching user by id
func (u *UserRepository) FindUserByID(ID int) *model.User {
	user := &model.User{}

	if err := u.store.db.QueryRow("SELECT id, tgid, username, firstname FROM users WHERE id = ?", ID).Scan(&user.ID, &user.TgID, &user.UserName, &user.FirstName); err != nil {
		fmt.Println(err)
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
