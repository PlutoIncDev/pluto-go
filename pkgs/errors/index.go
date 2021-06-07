package errors

type PlutoError struct {
	msg string
}

func (e PlutoError) Error() string {
	return e.msg
}

func NewErr(msg string) PlutoError {
	return PlutoError{msg: msg}
}