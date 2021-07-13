package errors

import (
	"errors"
	"fmt"
)

type Error interface {
	Error() string        // 包含code和msg
	Code() ErrorCode      // 错误码
	Desc() string         // 错误描述
	Data() interface{}    // 附加信息
	WithData(interface{}) // 给错误添加附加信息
	Is(target error) bool
	As(target interface{}) bool
}

type _err struct {
	data interface{} // 如果需要包含特定的数据
	code ErrorCode   // 错误码
	msg  string      // 错误描述
}

func New(err string) Error {
	return &_err{
		code: DefaultErrCode,
		msg:  err,
	}
}

func NewWithCode(err string, code ErrorCode) Error {
	return &_err{
		code: code,
		msg:  err,
	}
}

func (e *_err) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("err: %d-%s", e.code, e.msg)
}

func (e *_err) Code() ErrorCode {
	if e == nil {
		return 0
	}
	return e.code
}

func (e *_err) Desc() string {
	if e == nil {
		return ""
	}
	return e.msg
}

func (e *_err) Data() interface{} {
	if e == nil {
		return nil
	}
	return e.data
}

func (e *_err) WithData(data interface{}) {
	if data == nil || e == nil {
		return
	}
	e.data = data
}

func (e *_err) Is(target error) bool {
	return errors.Is(e, target)
}

func (e *_err) As(target interface{}) bool {
	return errors.As(e, target)
}

func (e *_err) Unwrap() error {
	return errors.Unwrap(e)
}
