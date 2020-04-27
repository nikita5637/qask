package model

import (
	"fmt"
)

var (
	minFirstNameLength       = 3
	minUserNameLength        = 6
	minTgID            int64 = 1
)

type userPublic struct {
	FirstName string `json:"firstName"`
}

type userPrivate struct {
	UserName string `json:"userName"`
	TgID     int64  `json:"tgID"`
}

//User is a user model
type User struct {
	userPublic
	userPrivate
}

//Validate is a function for validating user model
func (u *User) Validate() error {
	if len(u.FirstName) < minFirstNameLength {
		return fmt.Errorf("Firstname less than %d", minFirstNameLength)
	}

	if len(u.UserName) < minUserNameLength {
		return fmt.Errorf("Username less than %d", minUserNameLength)
	}

	if u.TgID < minTgID {
		return fmt.Errorf("Telegram ID less than %d", minTgID)
	}

	return nil
}
