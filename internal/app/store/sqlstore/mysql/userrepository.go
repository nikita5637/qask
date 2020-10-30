package mysql

import (
	"errors"
	"fmt"
	"qask/internal/app/model"
	"qask/internal/app/qaskerrors"

	"github.com/go-sql-driver/mysql"
)

//UserRepository is a users sql store
type UserRepository struct {
	store *Store
}

func (u *UserRepository) createUser(user *model.User) (int64, error) {
	res, err := u.store.db.Exec("INSERT INTO users (username, firstname, tgid) VALUES (?, ?, ?);", user.UserName, user.FirstName, user.TgID)
	if err != nil {
		sqlErr, ok := err.(*mysql.MySQLError)
		if !ok {
			return 0, err
		}

		switch sqlErr.Number {
		case 1062:
			return 0, qaskerrors.QaskErr{
				Message: sqlErr.Error(),
				Code:    sqlErr.Number,
				Err:     qaskerrors.EDUPENTRY,
			}
		case 1064:
			return 0, qaskerrors.QaskErr{
				Message: sqlErr.Error(),
				Code:    sqlErr.Number,
				Err:     qaskerrors.EINVALIDMYSQLSYNTAX,
			}
		}
		return 0, sqlErr
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

//CreateUser is a function for creating user
func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	if err := user.Validate(); err != nil {
		return 0, err
	}

	id, err := u.createUser(user)
	if err != nil {
		if errors.Is(err, qaskerrors.ErrUserExists) {
			return 0, fmt.Errorf("%s: %w", qaskerrors.ErrUserExists, err)
		} else if errors.Is(err, qaskerrors.ErrInvalidSQLSyntax) {
			return 0, fmt.Errorf("%s: %w", qaskerrors.ErrInvalidSQLSyntax, err)
		}

		return 0, qaskerrors.ErrUnknown
	}

	user.ID = id
	return id, nil
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
