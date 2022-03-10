package packets

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/maths"
)

type Packet interface {
	Packet(message Message) (data []byte, err error)
	UnPacket(reader io.Reader, message Message) error
}

type Message interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type packet struct {
	headerLen  int // 头部长度
	maxDataLen int // 数据最大长度
}

func NewPacket(headerLen, maxDataLen int) Packet {
	if headerLen <= 0 {
		headerLen = 4
	}
	return &packet{
		headerLen:  headerLen,
		maxDataLen: maths.MaxInt(0, maxDataLen),
	}
}

// Packet 封包
func (p *packet) Packet(message Message) (data []byte, err error) {
	if message == nil {
		err = errors.New("nil Message")
		return
	}
	mBytes, err := message.Marshal()
	if err != nil {
		return
	}
	if p.maxDataLen > 0 && len(mBytes) > p.maxDataLen {
		err = errors.New("packet: message is too large")
		return
	}
	data = make([]byte, p.headerLen+len(mBytes))
	// 头headerLen个字节存放数据长度
	binary.LittleEndian.PutUint32(data[:4], uint32(len(mBytes)))
	// 将数据写进剩余的字节
	copy(data[4:], mBytes)
	return
}

// UnPacket 拆包
func (p *packet) UnPacket(reader io.Reader, message Message) (err error) {
	if message == nil {
		err = errors.New("nil Message")
		return
	}

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
	if p.maxDataLen > 0 && dataLen > int32(p.maxDataLen) {
		err = errors.New("unpacket: message is too large")
		return
	}

	// 读数据，根据数据长度从reader中读取对应的数据
	data := make([]byte, dataLen)
	n, err = io.ReadFull(reader, data)
	if err == nil {
		// 反序列化数据到对应的结构体中
		err = message.Unmarshal(data[:n])
	}
	return
}
