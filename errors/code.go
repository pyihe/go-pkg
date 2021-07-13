package errors

import (
	"fmt"
)

const (
	DefaultErrCode ErrorCode = -1
)

type ErrorCode int64

func (ec ErrorCode) Int() int {
	return int(ec)
}

func (ec ErrorCode) Int64() int64 {
	return int64(ec)
}

func (ec ErrorCode) Int32() int32 {
	return int32(ec)
}

func (ec ErrorCode) ToString() string {
	return fmt.Sprintf("%d", ec)
}

func (ec ErrorCode) Equal(target ErrorCode) bool {
	return ec == target
}
