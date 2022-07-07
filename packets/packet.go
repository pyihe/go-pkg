package packets

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/maths"
)

type IPacket interface {
	Packet(message []byte) (data []byte, err error)
	UnPacket(reader io.Reader) ([]byte, error)
}

type Message interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type Packet struct {
	headerLen   int  // 头部长度
	dataSize    int  // 数据最大长度
	initialized bool // 是否已经初始化
}

func NewPacket(headerLen, dataMaxSize int) *Packet {
	if maths.MaxInt(0, headerLen) == 0 {
		headerLen = 4
	}
	if maths.MaxInt(0, dataMaxSize) == 0 {
		dataMaxSize = 4096
	}
	return &Packet{
		headerLen:   headerLen,
		dataSize:    maths.MaxInt(0, dataMaxSize),
		initialized: true,
	}
}

func (p *Packet) assert() {
	if !p.initialized {
		panic("Packet must be initialize by NewPacket")
	}
}

// Packet 封包
func (p *Packet) Packet(message []byte) (data []byte, err error) {
	p.assert()
	if p.dataSize > 0 && len(message) > p.dataSize {
		err = errors.New("packet: message is too large")
		return
	}
	data = make([]byte, p.headerLen+len(message))
	// 头headerLen个字节存放数据长度
	binary.LittleEndian.PutUint32(data[:4], uint32(len(message)))
	// 将数据写进剩余的字节
	copy(data[4:], message)
	return
}

// UnPacket 拆包
func (p *Packet) UnPacket(reader io.Reader) (b []byte, err error) {
	if reader == nil {
		return
	}
	p.assert()
	// 先读取header中的数据长度
	header := make([]byte, p.headerLen)
	n, err := io.ReadFull(reader, header)
	if err != nil {
		return
	}
	var dataLen int32
	if err = binary.Read(bytes.NewReader(header[:n]), binary.LittleEndian, &dataLen); err != nil {
		return
	}

	// 判断数据长度是否合法
	if p.dataSize > 0 && dataLen > int32(p.dataSize) {
		err = errors.New("unpacket: message is too large")
		return
	}

	// 读数据，根据数据长度从reader中读取对应的数据
	data := make([]byte, dataLen)
	n, err = io.ReadFull(reader, data)
	if err == nil {
		b = data[:n]
	}
	return
}
