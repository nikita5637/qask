package store

import "qask/internal/app/model"

//UserRepository is a interface
//for working with users
type UserRepository interface {
	CreateUser(*model.User) (int64, error)
	FindUserByID(int64) *model.User
	FindUserByTgID(int64) *model.User
	FindUserByUserName(string) *model.User
	GetUsers() []*model.User
}
