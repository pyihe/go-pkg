package tree

import (
	"fmt"
)

type Element interface {
	Value() interface{}
	Compare(Element) int // 相等返回0，小于返回负数，大于返回正数
}

/*
	平衡二叉树
	特点:
	1. 平衡二叉树是一种二叉查找树
	2. 每个结点的左子树的高度减去右子树的高度的绝对值不超过1
	3. 空树和左右子树都是平衡二叉树
	4. 相比红黑树，平衡二叉树比较适用于没有删除的情况

	平衡二叉树与二叉搜索树的区别在于是否平衡，所以在对树做更改的时候需要判断树是否平衡，其他操作与二叉搜索树相同。
	所以这里只写添加和删除节点。

	平衡二叉树的操作是通过对树进行旋转，旋转分为向左和向右旋转，以及这两个的组合。
	1. 向左旋转
	2. 向右旋转
*/

type AVLNode struct {
	Data   Element  // 存放的元素
	Height int      // 存放节点的高度
	Left   *AVLNode // 左子树
	Right  *AVLNode // 右子树
}

func NewAVLTree(data Element) *AVLNode {
	return &AVLNode{
		Data:   data,
		Height: 1,
		Left:   nil,
		Right:  nil,
	}
}

// 根左右
func (avl *AVLNode) String() string {
	if avl == nil {
		return ""
	}
	var result string
	result += fmt.Sprintf("%v ", avl.Data.Value())
	result += avl.Left.String()
	result += avl.Right.String()
	return result
}

// AddNode 添加节点
func (avl *AVLNode) AddNode(data Element) (root *AVLNode, ok bool) {
	defer func() {
		if ok {
			root = balance(avl)
			root.Height = maxHeight(getHeight(root.Left), getHeight(root.Right)) + 1
		}
	}()

	var target = &AVLNode{
		Data:   data,
		Height: 1,
		Left:   nil,
		Right:  nil,
	}

	cmpResult := avl.Data.Compare(data)
	switch {
	case cmpResult > 0:
		if avl.Left == nil {
			avl.Left = target
			ok = true
			return
		}
		avl.Left, ok = avl.Left.AddNode(data)
		return
	case cmpResult < 0:
		if avl.Right == nil {
			avl.Right = target
			ok = true
			return
		}
		avl.Right, ok = avl.Right.AddNode(data)
		return
	default:
		ok = true
		return
	}
}

// RemoveNode 移除节点
func (avl *AVLNode) RemoveNode(data Element) (root *AVLNode, ok bool) {
	defer func() {
		if ok && root != nil {
			root = balance(root)
			root.Height = maxHeight(getHeight(root.Left), getHeight(root.Right)) + 1
		}
	}()

	var temp *AVLNode
	cmpResult := avl.Data.Compare(data)
	switch {
	case cmpResult > 0:
		if avl.Left == nil {
			return nil, false
		}
		temp, ok = avl.Left.RemoveNode(data)
		if ok {
			avl.Left = temp
		}
		root = avl
		return
	case cmpResult < 0:
		if avl.Right == nil {
			return nil, false
		}
		temp, ok = avl.Right.RemoveNode(data)
		if ok {
			avl.Right = temp
		}
		root = avl
		return
	default:
		if avl.Left == nil && avl.Right == nil {
			return nil, true
		}
		if avl.Left == nil {
			root, ok = avl.Right, true
			return
		}
		if avl.Right == nil {
			root, ok = avl.Left, true
			return
		}
		// 将左子树最大节点提升到头节点，并将该最大节点从左子树中删除
		// avl.Data = avl.Left.Max().Data
		avl.Data = maxNode(avl.Left).(*AVLNode).Data
		avl.Left, ok = avl.Left.RemoveNode(avl.Data)
		root = avl
		return
	}
}

// 调整树为二叉平衡树
func balance(avl *AVLNode) *AVLNode {
	if avl == nil {
		return nil
	}
	if getHeight(avl.Right)-getHeight(avl.Left) == 2 {
		if getHeight(avl.Right.Right) > getHeight(avl.Right.Left) {
			avl = avl.leftRotate()
		} else {
			avl = avl.rightLeftRotate()
		}
	} else if getHeight(avl.Left)-getHeight(avl.Right) == 2 {
		if getHeight(avl.Left.Left) > getHeight(avl.Left.Right) {
			avl = avl.rightRotate()
		} else {
			avl = avl.leftRightRotate()
		}
	}
	return avl
}

// 向左旋转，单旋转，返回旋转后的头节点
func (avl *AVLNode) leftRotate() *AVLNode {
	node := avl.Right
	avl.Right = node.Left
	node.Left = avl

	avl.Height = maxHeight(getHeight(avl.Left), getHeight(avl.Right)) + 1
	node.Height = maxHeight(getHeight(node.Left), getHeight(node.Right)) + 1
	return node
}

// 先将右孩子右旋转，然后自己右旋转
func (avl *AVLNode) rightLeftRotate() *AVLNode {
	avl.Right = avl.Right.rightRotate()
	return avl.leftRotate()
}

// 向右旋转，单旋转，旋转结果为
func (avl *AVLNode) rightRotate() *AVLNode {
	node := avl.Left      // 左子树
	avl.Left = node.Right // 左子树的右子树变为左子树
	node.Right = avl      // avl降为左子树的右子树

	// 更新节点高度
	avl.Height = maxHeight(getHeight(avl.Left), getHeight(avl.Right)) + 1
	node.Height = maxHeight(getHeight(node.Left), getHeight(node.Right)) + 1
	return node
}

// 先将左孩子左旋转，自己再右旋转
func (avl *AVLNode) leftRightRotate() *AVLNode {
	avl.Left = avl.Left.leftRotate()
	return avl.rightRotate()
}

// Max 获取最大节点
func (avl *AVLNode) Max() *AVLNode {
	node := avl
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func maxNode(node interface{}) interface{} {
	switch t := node.(type) {
	case *AVLNode:
		temp := t
		for temp.Right != nil {
			temp = temp.Right
		}
		return temp
	case *RedBlackNode:
		temp := t
		for temp.Right.Data != nil {
			temp = temp.Right
		}
		return temp
	default:
		return nil
	}
}

func minNode(node interface{}) interface{} {
	switch t := node.(type) {
	case *AVLNode:
		temp := t
		for temp.Left != nil {
			temp = temp.Left
		}
		return temp
	case *RedBlackNode:
		temp := t
		for temp.Left.Data != nil {
			temp = temp.Left
		}
		return temp
	default:
		return nil
	}
}

// 返回更大的那个值
func maxHeight(h1, h2 int) int {
	if h1 > h2 {
		return h1
	} else {
		return h2
	}
}

func getHeight(avl *AVLNode) int {
	if avl != nil {
		return avl.Height
	}
	return 0
}
