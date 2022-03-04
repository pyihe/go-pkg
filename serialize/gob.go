package serialize

import (
	"bytes"
	"encoding/gob"
)

type gobSerializer struct{}

func Gob() Serializer {
	return gobSerializer{}
}

func (gb gobSerializer) Encode(v interface{}) (data []byte, err error) {
	buff := bytes.NewBuffer([]byte{})
	err = gob.NewEncoder(buff).Encode(v)
	data = buff.Bytes()
	return
}

func (gb gobSerializer) Decode(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}
