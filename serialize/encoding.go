package serialize

import (
	"strings"
)

type Codec interface {
	Name() string
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

var (
	m = make(map[string]Codec)
)

func Register(c Codec) {
	m[strings.ToLower(c.Name())] = c
}

func Get(name string) Codec {
	return m[strings.ToLower(name)]
}
