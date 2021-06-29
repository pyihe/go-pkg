package errors

import (
	"fmt"
)

const (
	ErrorCodeFail = 100000 //默认错误代码
)

type Error interface {
	Error() string
	Code() int
	Message() string
	WithCode(code int) Error
	WithMessage(msg string) Error
}

type myError struct {
	code    int
	message string
}

func NewError(code int, message string) Error {
	return &myError{
		code:    code,
		message: message,
	}
}

func Errorf(m string, args ...interface{}) Error {
	e := &myError{}
	e.message = fmt.Sprintf(m, args...)

	return e
}

func (m *myError) WithCode(code int) Error {
	if m == nil {
		return nil
	}
	m.code = code
	return m
}

func (m *myError) WithMessage(msg string) Error {
	if m == nil {
		return nil
	}
	m.message = msg
	return m
}

func (m *myError) Error() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("Code: %d, Message: %s", m.code, m.message)
}

func (m *myError) Message() string {
	if m == nil {
		return ""
	}
	return m.message
}

func (m *myError) Code() int {
	if m != nil {
		return m.code
	}
	return 0
}
