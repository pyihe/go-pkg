package protobuffer

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/pyihe/go-pkg/serialize"
)

const Name = "proto"

func init() {
	serialize.Register(&protoCodec{})
}

type protoCodec struct {
}

func (p *protoCodec) Name() string {
	return Name
}

func (p *protoCodec) Marshal(v interface{}) ([]byte, error) {
	vv, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("faild to marshal: message is %T, want proto.Message", v)
	}
	return proto.Marshal(vv)
}

func (p *protoCodec) Unmarshal(data []byte, v interface{}) error {
	vv, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("failed to unmarshal: message is %T, want proto.Message", v)
	}
	return proto.Unmarshal(data, vv)
}
