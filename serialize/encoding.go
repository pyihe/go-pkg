package serialize

type Serializer interface {
	Encode(v interface{}) ([]byte, error)
	Decode(data []byte, v interface{}) error
}
