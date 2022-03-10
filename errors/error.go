package errors

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
	return e.err
}

func (e *Error) Code() int32 {
	return e.code
}
