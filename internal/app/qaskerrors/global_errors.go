package qaskerrors

import "errors"

const (
	//EDUPENTRY ...
	EDUPENTRY = Errno(0x1)
	//EEMPTYBODY ...
	EEMPTYBODY = Errno(0x2)
	//EUNKNOWNFROM ...
	EUNKNOWNFROM = Errno(0x3)
	//EINVALIDMYSQLSYNTAX ...
	EINVALIDMYSQLSYNTAX = Errno(0x4)
	//EINVALIDPOSTGRESQLSYNTAX ...
	EINVALIDPOSTGRESQLSYNTAX = Errno(0x5)
	//EMYSQLDATATOOLONG ...
	EMYSQLDATATOOLONG = Errno(0x6)
)

var _errors = [...]string{
	1: "SQL duplicate entry",
	2: "Empty body",
	3: "Unknown from",
	4: "MySQL invalid query syntax",
	5: "PostgreSQL invalid query syntax",
	6: "MySQL Data too long",
}

var (
	//ErrUserExists is error that reports that the user alredy exists
	ErrUserExists = errors.New("User already exists")

	//ErrEmptyBody is error that returns, when request body is empty
	ErrEmptyBody = errors.New("Empty body")

	//ErrUnknownFrom is error that returns, when "from" field is unknown
	ErrUnknownFrom = errors.New("Unknown from")

	//ErrInvalidSQLSyntax is error that returns, when sql request has invalid syntax
	ErrInvalidSQLSyntax = errors.New("SQL invalid query syntax")

	//ErrSQLDataTooLong is error that returns, when sql data length too long
	ErrSQLDataTooLong = errors.New("SQL Data too long")

	//ErrUnknown is error which returns by default
	ErrUnknown = errors.New("Unknown error")
)
