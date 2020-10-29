package qaskerrors

var (
	//ErrUserExists is error that reports that the user alredy exists
	ErrUserExists = New("User already exists", 50)

	//ErrEmptyBody is error that returns, when request body is empty
	ErrEmptyBody = New("Empty body", 4)

	//ErrUnknownFrom is error that returns, when "from" field is unknown
	ErrUnknownFrom = New("Unknown from", 5)
)
