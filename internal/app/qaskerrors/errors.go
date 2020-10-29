package qaskerrors

//QaskErr is a main type of error
type QaskErr struct {
	Message string
	Code    uint16
	Err     error
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
func New(s string, c uint16) QaskErr {
	return QaskErr{s, c, nil}
}
