package gobserialize

import (
	"bytes"
	"encoding/gob"

	"github.com/pyihe/go-pkg/serialize"
)

const Name = "gob"

func init() {
	serialize.Register(&gobCodec{})
}

type gobCodec struct {
}

func (g *gobCodec) Name() string {
	return Name
}

func (g *gobCodec) Marshal(v interface{}) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	err := gob.NewEncoder(buff).Encode(v)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func (g *gobCodec) Unmarshal(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}
