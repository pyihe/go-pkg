package queue

import "github.com/pyihe/go-pkg/errors"

type Type uint8

const (
	LinearQueue Type = iota + 1 // 线性队列
	ChainQueue                  // 链式队列
)

type Queue interface {
	Reset()                          // 重置队列
	Len() int                        // 返回队列长度
	Head() (interface{}, bool)       // 获取队首(不删除)
	EnQueue(interface{}) error       // 入队列
	DeQueue() (interface{}, error)   // 出队列
	Each(func(ele interface{}) bool) // 遍历每个元素
}

type Config struct {
	Type Type
	Size int // 用于限制切片式队列的长度
}

func New(config *Config) (Queue, error) {
	if config == nil {
		panic("nil config")
	}
	switch config.Type {
	case LinearQueue:
		return newLinearQueue(config.Size), nil
	case ChainQueue:
		return newChainQueue(), nil
	default:
		return nil, errors.New("unrecognized queue type")
	}
}
