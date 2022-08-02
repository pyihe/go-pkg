package msgpackserialize

import (
	"github.com/pyihe/go-pkg/serialize"
	"github.com/vmihailenco/msgpack/v5"
)

const Name = "msgpack"

func init() {
	serialize.Register(&msgpackCodec{})
}

type msgpackCodec struct {
}

func (m *msgpackCodec) Name() string {
	return Name
}

func (m *msgpackCodec) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

func (m *msgpackCodec) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
