package serialize

import "github.com/vmihailenco/msgpack/v5"

type msgPack struct{}

func Msgpack() Serializer {
	return msgPack{}
}

func (mp msgPack) Encode(v interface{}) (data []byte, err error) {
	return msgpack.Marshal(v)
}

func (mp msgPack) Decode(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
