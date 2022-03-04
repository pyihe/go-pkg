package serialize

import (
	"errors"

	"github.com/golang/protobuf/proto"
)

type protoSerializer struct{}

func Proto() Serializer {
	return protoSerializer{}
}

func (ps protoSerializer) Encode(v interface{}) (data []byte, err error) {
	if m, ok := v.(proto.Message); ok {
		return proto.Marshal(m)
	}
	err = errors.New("not proto.Message")
	return
}

func (ps protoSerializer) Decode(data []byte, v interface{}) (err error) {
	if m, ok := v.(proto.Message); ok {
		return proto.Unmarshal(data, m)
	}
	err = errors.New("not proto.Message")
	return
}
