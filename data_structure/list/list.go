package list

import "github.com/pyihe/go-pkg/errors"

var (
	ErrInvalidPos = errors.New("invalid pos")
)

// Lister 线性表
type Lister interface {
	Len() int                              // 返回线性表长度
	Reset()                                // 重置线性表
	Get(pos int) (interface{}, bool)       // 获取某个位置的元素, 注意pos从1开始, 而不是0开始
	Pos(ele interface{}) int               // 获取某个元素的位置, 返回的位置从1开始
	Insert(pos int, ele interface{}) error // 往队列中插入元素
	Delete(pos int) bool                   // 删除某个元素
}
