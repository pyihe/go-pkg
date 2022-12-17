package tree

import (
	"container/list"
	"fmt"
)

/*
	二叉搜索树(二叉查找树)
	二叉搜索树特点:
	1. 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
	2. 若任意节点的右子树不空，则右子树上所有节点的值均大于或等于它的根节点的值；
	3. 任意节点的左、右子树也分别为二叉查找树；
	1. 添加节点
	2. 删除节点
	3. 查找节点
	4. 最大/最小值
	5. 递归前中后遍历
	6. 非递归前中后遍历
	7. 深度/广度优先遍历
*/

// 二叉搜索树: 左边的节点都比跟节点小，右边的节点都比跟节点大
type BinaryTreeNode struct {
	ValueCnt int             //表示data的数量
	Data     Element         //数据
	Left     *BinaryTreeNode //左子树: 假定左子树都比跟节点小
	Right    *BinaryTreeNode //右子树: 假定右子树逗比跟节点大
}

func NewTree(element Element) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data:     element,
		ValueCnt: 1,
	}
}

// 输出整棵树
func (tree *BinaryTreeNode) String() string {
	if tree == nil {
		return ""
	}
	var result string
	result += tree.Left.String()
	result += fmt.Sprintf("%v ", tree.Data.Value())
	result += tree.Right.String()
	return result
}

// 输出节点数据
func (tree *BinaryTreeNode) DataString() string {
	return fmt.Sprintf("%v ", tree.Data.Value())
}

// 添加节点
func (tree *BinaryTreeNode) AddNode(data Element) (*BinaryTreeNode, bool) {
	if data == nil {
		return nil, false
	}

	var target = &BinaryTreeNode{
		Data:     data,
		ValueCnt: 1,
		Left:     nil,
		Right:    nil,
	}
	cmpResult := tree.Data.Compare(target.Data)
	switch {
	case cmpResult > 0: //如果target比给定根节点小，则将target插入tree的左子树
		if tree.Left == nil {
			tree.Left = target
			return target, true
		}
		return tree.Left.AddNode(data)

	case cmpResult < 0: //如果目标节点比给定根节点大，则将target插入tree的右子树
		if tree.Right == nil {
			tree.Right = target
			return target, true
		}
		return tree.Right.AddNode(data)

	default: //如果相等，证明给定的树中已经存在对应值的节点了，不需要再插入
		tree.ValueCnt += 1
		return tree, true
	}
}

// 从二叉树中移除某一个节点
func (tree *BinaryTreeNode) RemoveNode(data Element) (*BinaryTreeNode, bool) {
	if data == nil {
		return nil, false
	}
	var ok bool
	var sub *BinaryTreeNode
	cmpResult := tree.Data.Compare(data)
	switch {
	case cmpResult > 0: //小于根节点，则从左子树递归删除
		if tree.Left == nil {
			return nil, ok
		}
		sub, ok = tree.Left.RemoveNode(data)
		if ok {
			tree.Left = sub
		}
		return tree, ok
	case cmpResult < 0: //大于根节点, 则从右子树递归删除
		if tree.Right == nil {
			return nil, ok
		}
		sub, ok = tree.Right.RemoveNode(data)
		if ok {
			tree.Right = sub
		}
		return tree, ok
	default: //根节点即为需要删除的节点，删除的同时需要提升子节点为根节点
		if tree.Left == nil && tree.Right == nil { //如果只有根节点，则返回nil
			return nil, true
		}
		if tree.Left == nil { //如果右子树不为空，左子树为空
			return tree.Right, true
		}
		if tree.Right == nil { //如果左子树不为空，右子树为空
			return tree.Left, true
		}
		//如果左右子树都不为空则找出左子树中最大的节点或者右子树中最小的节点作为根节点
		leftMaxNode := tree.Left.Max() //找出左子树最大的节点
		_, _ = tree.RemoveNode(leftMaxNode.Data)
		leftMaxNode.Left = tree.Left
		leftMaxNode.Right = tree.Right
		return leftMaxNode, true
	}
}

// 搜索对应值的节点
func (tree *BinaryTreeNode) Search(ele Element) *BinaryTreeNode {
	cmpResult := tree.Data.Compare(ele)

	switch {
	case cmpResult < 0: //查找的元素节点在右子树
		if tree.Right == nil {
			return nil
		}
		return tree.Right.Search(ele)
	case cmpResult > 0: //查找的元素节点在左子树
		if tree.Left == nil {
			return nil
		}
		return tree.Left.Search(ele)

	default: //查找的元素节点刚好为tree
		return tree
	}

}

// 找出树中最大的节点
func (tree *BinaryTreeNode) Max() *BinaryTreeNode {
	//递归
	//if tree.Right == nil {
	//	return tree
	//}
	//return tree.Right.Max()

	//非递归
	node := tree
	for node.Right != nil {
		node = node.Right
	}
	return node
}

// 找出树中最小的节点
func (tree *BinaryTreeNode) Min() *BinaryTreeNode {
	//递归
	//if tree.Left == nil {
	//	return tree
	//}
	//return tree.Left.Min()

	//非递归
	node := tree
	for node.Left != nil {
		node = node.Left
	}
	return node
}

/*二叉树遍历（递归）*/
//前序遍历:以当前节点为根节点，根——>左——>右
func (tree *BinaryTreeNode) PreTraverseRecursion() (treeString string) {
	if tree == nil {
		return
	}
	treeString += tree.DataString()
	treeString += tree.Left.PreTraverseRecursion()
	treeString += tree.Right.PreTraverseRecursion()
	return
}

// 中序遍历:以当前节点为根节点，左——>根——>右
func (tree *BinaryTreeNode) MidTraverseRecursion() (treeString string) {
	if tree == nil {
		return
	}
	treeString += tree.Left.MidTraverseRecursion()
	treeString += tree.DataString()
	treeString += tree.Right.MidTraverseRecursion()
	return
}

// 后续遍历：以当前节点为根节点，左——>右——>根
func (tree *BinaryTreeNode) PostTraverseRecursion() (treeString string) {
	if tree == nil {
		return
	}
	treeString += tree.Left.PostTraverseRecursion()
	treeString += tree.Right.PostTraverseRecursion()
	treeString += tree.DataString()
	return
}

/*非递归遍历：利用栈结构*/
//栈结构
type Stack struct {
	*list.List
}

// 出栈
func (s *Stack) Pop() interface{} {
	if s == nil || s.Len() <= 0 {
		return nil
	}
	value := s.Back()
	s.Remove(value)
	return value.Value
}

// 进栈
func (s *Stack) Push(d interface{}) {
	if s == nil {
		return
	}
	s.PushBack(d)
}

// 获取栈顶元素
func (s *Stack) Top() interface{} {
	if s == nil {
		return nil
	}
	return s.Back().Value
}

/*
非递归前序遍历
根——>左——>右
1、从根节点开始访问，每访问一个元素，执行入栈操作并输出当前节点
2、访问到最左边的子节点时，开始出栈
3、每出栈一个元素需要该节点是否存在右节点，如果存在则重复操作1
*/
func (tree *BinaryTreeNode) PreTraverse() (result string) {
	stack := &Stack{
		List: list.New(),
	}

	node := tree
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.Push(node)
			result += node.DataString()
			node = node.Left
		} else {
			node = stack.Pop().(*BinaryTreeNode)
			node = node.Right
		}
	}
	return
}

/*
非递归中序遍历
左——>根——>右
1、从根节点开始遍历到最左边的子节点，每访问一个节点就入栈（此处用node访问每个节点）
2、访问到最左边的子节点时开始出栈，出栈时做输出操作
3、每次出栈一个元素，需要判断该元素是否存在右节点，如果存在，则重复步骤1
*/
func (tree *BinaryTreeNode) MidTraverse() (result string) {
	stack := list.New()

	node := tree
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {
			ele := stack.Back()
			stack.Remove(ele)
			node = ele.Value.(*BinaryTreeNode)
			result += node.DataString()
			node = node.Right
		}
	}
	return
}

/*
非递归后续遍历
左——>右——>根
1、从根节点开始遍历到最左边的子节点，每访问一个节点就入栈（此处用node访问每个节点）
2、最后一个左子节点入栈后开始出栈操作，出栈时做输出操作
3、出栈条件：栈顶元素的右子节点为空或者右子节点已经出栈（此处用top纪录当前栈顶元素，last纪录最后出栈的元素）
4、如果栈顶元素的右子节点不为空且未出栈，则继续步骤1
为什么要纪录最后出站的元素？
如果一个节点同时存在左右子节点，按照后序遍历的规则，最后一个出栈元素为一定为该节点的右子节点，此时该节点的子节点已经遍历完，需要将该节点出栈并输出
*/
func (tree *BinaryTreeNode) PostTraverse() (result string) {
	stack := list.New()

	node := tree
	var topNode, lastNode *BinaryTreeNode //top为栈顶元素、last为最后出栈的元素

	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.PushBack(node)
			node = node.Left
		} else {
			ele := stack.Back().Value
			topNode = ele.(*BinaryTreeNode)
			if topNode.Right == nil || topNode.Right == lastNode {
				stack.Remove(stack.Back())
				result += topNode.DataString()
				lastNode = topNode
			} else {
				node = topNode.Right
			}
		}
	}
	return
}

/*
	广度优先遍历(BFS), 即层次遍历, 从根节点开始从左向右每一层遍历。
	这里利用的队列，将根节点入列，当队列中元素大于0时，挨个出列，每出列一个元素，同时将该元素的左右节点依次入列，直到队列为空
*/

func (tree *BinaryTreeNode) BFSTraverse() (result string) {
	treeList := list.New()
	treeList.PushBack(tree)

	for treeList.Len() > 0 {
		element := treeList.Front()
		node := element.Value.(*BinaryTreeNode)

		result += node.DataString()
		treeList.Remove(element)

		if node.Left != nil {
			treeList.PushBack(node.Left)
		}
		if node.Right != nil {
			treeList.PushBack(node.Right)
		}
	}
	return
}

/*
深度优先遍历(DFS), 从根节点开始向下访问每个子节点，直到最后一个节点或者没有节点可以访问了为止，然后在向上返回至最近一个仍然有子节点未被访问的节点的子节点开始访问。
算法实现利用栈的特性，先根节点入栈，然后出栈(遍历)，然后依次入栈右子树和左子树，继续出栈。
*/
func (tree *BinaryTreeNode) DFSTraverse() (result string) {
	stack := &Stack{
		List: list.New(),
	}

	stack.PushBack(tree)

	for stack.Len() > 0 {
		ele := stack.Back()
		node := ele.Value.(*BinaryTreeNode)

		result += node.DataString()

		stack.Remove(ele)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return
}
