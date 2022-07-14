package packets

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/pyihe/go-pkg/errors"
)

var (
	ErrMsgTooLong  = errors.New("message too long")
	ErrMsgTooShort = errors.New("message too short")
	ErrHeaderSize  = errors.New("header size can be: 1/2/4/8")
)

type Option func(*Packet)

type IPacket interface {
	Packet(message []byte) (data []byte, err error)
	UnPacket(reader io.Reader) ([]byte, error)
}

type Packet struct {
	bigEndian   bool                         // 是否小端序
	initialized bool                         // 是否已经初始化
	headerSize  int                          // 头部长度
	maxMsgSize  int                          // 消息最大长度
	minMsgSize  int                          // 消息最小长度
	encrypter   func([]byte) ([]byte, error) // 消息加密方法
	decrypter   func([]byte) ([]byte, error) // 消息解密方法
}

func NewPacket(opts ...Option) IPacket {
	p := &Packet{
		initialized: true,
		headerSize:  4, // 默认头部长度4字节
	}
	for _, op := range opts {
		op(p)
	}
	return p
}

func (p *Packet) assert() {
	if !p.initialized {
		panic("Packet must be initialize by NewPacket")
	}
}

func (p *Packet) endian() binary.ByteOrder {
	if p.bigEndian {
		return binary.BigEndian
	}
	return binary.LittleEndian
}

func (p *Packet) checkMessageSize(size int) (err error) {
	if p.maxMsgSize > 0 && size > p.maxMsgSize {
		err = ErrMsgTooLong
		return
	}
	if p.minMsgSize > 0 && size < p.minMsgSize {
		err = ErrMsgTooShort
		return
	}
	return
}

// Packet 封包
func (p *Packet) Packet(message []byte) (data []byte, err error) {
	p.assert()

	if p.encrypter != nil {
		message, err = p.encrypter(message)
		if err != nil {
			return
		}
	}

	// 判断消息大小是否正确
	size := len(message)
	if err = p.checkMessageSize(size); err != nil {
		return
	}

	byteOrder := p.endian()
	// 将消息长度写入header
	data = make([]byte, p.headerSize+size)
	switch p.headerSize {
	case 1:
		data[0] = byte(size)
	case 2:
		byteOrder.PutUint16(data, uint16(size))
	case 4:
		byteOrder.PutUint32(data, uint32(size))
	case 8:
		byteOrder.PutUint64(data, uint64(size))
	default:
		err = ErrHeaderSize
		return
	}

	// 写入数据
	copy(data[p.headerSize:], message)
	return
}

// UnPacket 拆包
func (p *Packet) UnPacket(reader io.Reader) (b []byte, err error) {
	p.assert()

	if reader == nil {
		return
	}

	// 先读取header中的数据长度
	header := make([]byte, p.headerSize)
	n, err := io.ReadFull(reader, header)
	if err != nil {
		return
	}

	var size uint32
	if err = binary.Read(bytes.NewReader(header[:n]), p.endian(), &size); err != nil {
		return
	}

	// 判断数据长度是否合法
	if err = p.checkMessageSize(int(size)); err != nil {
		return
	}

	// 读数据，根据数据长度从reader中读取对应的数据
	b = make([]byte, size)
	n, err = io.ReadFull(reader, b)
	if err != nil {
		return
	}

	if p.decrypter != nil {
		return p.decrypter(b[:n])
	}
	return b[:n], nil
}

func WithHeaderSize(header int) Option {
	return func(packet *Packet) {
		packet.headerSize = header
	}
}

func WithMaxMsgSize(size int) Option {
	return func(packet *Packet) {
		packet.maxMsgSize = size
	}
}

func WithMinMsgSize(size int) Option {
	return func(packet *Packet) {
		packet.minMsgSize = size
	}
}

func WithBigEndian(b bool) Option {
	return func(packet *Packet) {
		packet.bigEndian = b
	}
}

func WithEncrypter(encrypter func([]byte) ([]byte, error)) Option {
	return func(packet *Packet) {
		packet.encrypter = encrypter
	}
}

func WithDecrypter(decrypter func([]byte) ([]byte, error)) Option {
	return func(packet *Packet) {
		packet.decrypter = decrypter
	}
}
