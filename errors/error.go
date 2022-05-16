package errors

import "fmt"

type Error struct {
	err  string
	code int32
}

func New(err string, codes ...int32) error {
	e := &Error{
		err: err,
	}
	if len(codes) > 0 {
		e.code = codes[0]
	}
	return e
}

func (e *Error) Error() (err string) {
	if e.code == 0 {
		return e.err
	}
	return fmt.Sprintf("%d-%s", e.code, e.err)
}

func (e *Error) Code() int32 {
	return e.code
}
