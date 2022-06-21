package stack

import "github.com/pyihe/go-pkg/errors"

type Type uint8

const (
	LinearStack Type = iota + 1 // 线性存储
	ChainStack                  // 链式存储
)

type Stack interface {
	Reset()                              // 重置
	Len() int                            // 返回栈元素个数
	Top() (interface{}, bool)            // 如果存在, 返回栈顶元素(不删除)
	Push(interface{})                    // 压栈
	Pop() interface{}                    // 出栈
	Each(handler func(interface{}) bool) // 遍历每个元素
}

type Config struct {
	Type Type // 栈类型
	Size int  // 顺序存储时的切片大小, 并非限制栈的大小, 而是根据实际使用进行容量初始化, 避免使用过程中切片频繁扩容
}

func New(config *Config) (Stack, error) {
	if config == nil {
		panic("nil config")
	}
	switch config.Type {
	case LinearStack:
		return newLinearStack(config.Size), nil
	case ChainStack:
		return newChainStack(), nil
	default:
		return nil, errors.New("unrecognized stack type")
	}
}
