package qaskerrors

import (
	"strconv"
)

//Errno ...
type Errno int

func (e Errno) Error() string {
	if 0 <= int(e) && int(e) < len(_errors) {
		s := _errors[e]
		if s != "" {
			return s
		}
	}
	return "errno: " + strconv.Itoa(int(e))
}

//Is compares Errno error with target error
func (e Errno) Is(target error) bool {
	switch target {
	case ErrUserExists:
		return e == EDUPENTRY
	case ErrInvalidSQLSyntax:
		return e == EINVALIDMYSQLSYNTAX || e == EINVALIDPOSTGRESQLSYNTAX
	}
	return false
}

//QaskErr is a main type of error
type QaskErr struct {
	Message string `json:"errMessage"`
	Code    uint16 `json:"errCode"`
	Err     error  `json:"-"`
}

//Error returns an error description message
func (e QaskErr) Error() string {
	return e.Message
}

//Unwrap returns a wrapped error
func (e QaskErr) Unwrap() error {
	return e.Err
}

//New returns a new value with filled description and error code
func New(s string, c uint16) *QaskErr {
	return &QaskErr{s, c, nil}
}
