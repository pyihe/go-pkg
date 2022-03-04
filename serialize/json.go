package serialize

import "encoding/json"

type jsonSerializer struct{}

func JSON() Serializer {
	return jsonSerializer{}
}

func (js jsonSerializer) Encode(v interface{}) (data []byte, err error) {
	return json.Marshal(v)
}

func (js jsonSerializer) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
