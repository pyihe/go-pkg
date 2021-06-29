package encoding

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/vmihailenco/msgpack/v5"
)

type Encoding interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// JSONEncoding json格式
func JSONEncoding() Encoding {
	return &jsonEncoding{}
}

// GobEncoding
func GobEncoding() Encoding {
	return &gobEncoding{}
}

// MsgpackEncoding
func MsgpackEncoding() Encoding {
	return &msgpackEncoding{}
}

// ProtoEncoding
func ProtoEncoding() Encoding {
	return &protoEncoding{}
}

/***JSON Marshaler***/
type jsonEncoding struct{}

// Marshal
func (j *jsonEncoding) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal
func (j *jsonEncoding) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

/***gob Marshaler***/
type gobEncoding struct{}

// Marshal
func (g *gobEncoding) Marshal(v interface{}) ([]byte, error) {
	var b = bytes.NewBuffer(nil)
	err := gob.NewEncoder(b).Encode(v)
	return b.Bytes(), err
}

// Unmarshal
func (g *gobEncoding) Unmarshal(data []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}

/***msgpack Marshaler***/
type msgpackEncoding struct{}

// Marshal
func (m *msgpackEncoding) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// Unmarshal
func (m *msgpackEncoding) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}

/***proto Encoding***/
type protoEncoding struct{}

// Marshal
func (p *protoEncoding) Marshal(v interface{}) ([]byte, error) {
	m, ok := v.(proto.Message)
	if !ok {
		return nil, errors.New("not proto.Message")
	}

	return proto.Marshal(m)
}

// Unmarshal
func (p *protoEncoding) Unmarshal(data []byte, v interface{}) error {
	m, ok := v.(proto.Message)
	if !ok {
		return errors.New("not proto.Message")
	}
	return proto.Unmarshal(data, m)
}
