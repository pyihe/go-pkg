package json

import (
	"encoding/json"

	"github.com/pyihe/go-pkg/serialize"
)

const Name = "json"

func init() {
	serialize.Register(&jsCodec{})
}

type jsCodec struct{}

func (js *jsCodec) Name() string {
	return Name
}

func (js *jsCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (js *jsCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
