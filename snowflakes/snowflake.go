package snowflakes

import (
	"strconv"
	"sync"
	"time"
)

type Worker interface {
	GetInt64() int64
	GetString() string
}

//生成的64位ID的组成(从左到右编号由1至64):
//1: 1位符号位
//2-42: 41位时间戳
//43-64: 剩下22位由两部分组成: 分布式节点ID所占的位数和每个节点每毫秒内可以生成的ID数构成, 这两部分所占位数不固定, 但是总和必须等于22

const (
	nodeBits    uint8 = 10                      // 每台机器(节点)的ID位数 10位最大可以有2^10=1024个节点
	numberBits  uint8 = 12                      // 表示每个集群下的每个节点，1毫秒内可生成的id序号的二进制位数 即每毫秒可生成 2^12-1=4096个唯一ID
	nodeMax     int64 = -1 ^ (-1 << nodeBits)   // 节点ID的最大值，用于防止溢出
	numberMax   int64 = -1 ^ (-1 << numberBits) // 同上，用来表示生成id序号的最大值
	timeShift         = nodeBits + numberBits   // 时间戳向左的偏移量
	workerShift       = numberBits              // 节点ID向左的偏移量
)

type Option func(*builder) error

type builder struct {
	mu        sync.Mutex // 添加互斥锁 确保并发安全
	epoch     int64
	timestamp int64 // 记录时间戳
	workerId  int64 // 该节点的ID
	number    int64 // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
}

func NewWorker(workerId int64) Worker {
	assertWorkId(workerId)
	b := &builder{
		epoch:    time.Now().Unix() * 1000,
		workerId: workerId,
	}
	return b
}

func assertWorkId(workerId int64) {
	if workerId < 0 || workerId > nodeMax {
		panic("work id cannot more than 1024")
	}
}

func (w *builder) GetInt64() (id int64) {
	w.mu.Lock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	id = (now-w.epoch)<<timeShift | (w.workerId << workerShift) | (w.number)
	w.mu.Unlock()
	return
}

func (w *builder) GetString() string {
	return strconv.FormatInt(w.GetInt64(), 10)
}
