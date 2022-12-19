package tree

import (
	"container/list"
	"fmt"
)

/*
	红黑树也是一种二叉查找树，不同的是红黑树是一种自平衡的二叉查找树，除了具有二叉查找树所具有的特点以外，红黑树还具有以下特点:
	1. 节点是红色或黑色。
	2. 根节点是黑色。
	3. 每个叶子节点都是黑色的空节点（NIL节点）。
	4. 每个红色节点的两个子节点都是黑色。(从每个叶子到根的所有路径上不能有两个连续的红色节点)
	5. 从任一节点到其每个叶子的所有路径都包含相同数目的黑色节点。

	与平衡二叉树不同的是，红黑树的平衡度没有平衡二叉树那么高，平衡二叉树对所有左右子树的高度差要求不能大于1，而红黑树对左右子树
	的高度差要求为：从根节点到叶子节点的最大路径不会超过最短路径的2倍。
*/

type color uint

const (
	Black color = iota + 1
	Red
)

func (c color) String() string {
	if c == Black {
		return "BLACK"
	} else {
		return "RED"
	}
}

//红黑树节点
type RedBlackNode struct {
	Color  color         //节点颜色
	Data   Element       //数据
	Parent *RedBlackNode //父节点
	Left   *RedBlackNode //左孩子
	Right  *RedBlackNode //右孩子
}

//红黑树
type RedBlackTree struct {
	Root *RedBlackNode //根节点
}

func NewRedBlackNode(data Element) *RedBlackNode {
	node := &RedBlackNode{
		Color: Red,
		Data:  data,
	}
	var left = &RedBlackNode{
		Color:  Black,
		Data:   nil,
		Parent: node,
	}
	var right = &RedBlackNode{
		Color:  Black,
		Data:   nil,
		Parent: node,
	}
	node.Left = left
	node.Right = right
	return node
}

func (t *RedBlackNode) PreTraverse() (result string) {
	stack := &Stack{
		List: list.New(),
	}

	node := t
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.Push(node)
			result += fmt.Sprintf("%v(%v)---", node.Color, node.Data)
			node = node.Left
		} else {
			node = stack.Pop().(*RedBlackNode)
			node = node.Right
		}
	}
	return
}

//获取某个节点的祖父节点（爷爷节点）
func (t *RedBlackNode) grandParent() *RedBlackNode {
	//根节点没有父节点和祖父节点
	if t.Parent == nil {
		return nil
	}
	return t.Parent.Parent
}

//获取某个节点的叔父节点（父节点的兄弟节点）
func (t *RedBlackNode) uncleNode() *RedBlackNode {
	if t.grandParent() == nil {
		return nil
	}
	if t.Parent == t.grandParent().Left {
		return t.grandParent().Right
	}
	return t.grandParent().Left
}

//找到某个节点的兄弟节点
func (t *RedBlackNode) siblingNode() *RedBlackNode {
	if t.Parent == nil {
		return nil
	}
	if t == t.Parent.Left {
		return t.Parent.Right
	}
	return t.Parent.Left
}

//获取某个节点的根节点
func (t *RedBlackNode) getRoot() *RedBlackNode {
	node := t
	for node.Parent != nil {
		node = node.Parent
	}
	return node
}

//插入node
func (t *RedBlackNode) insert(data Element) (result *RedBlackNode) {
	if t.Data == nil && t.Left.Data == nil && t.Right.Data == nil {
		t.Data = data
		return t
	}
	cmpResult := t.Data.Compare(data)
	switch {
	case cmpResult > 0:
		if t.Left.Data == nil { //如果左孩子的数据域为空，则插入到左孩子的位置
			t.Left.Data = data
			t.Left.Color = Red
			t.Left.Left = &RedBlackNode{Parent: t.Left, Color: Black}
			t.Left.Right = &RedBlackNode{Parent: t.Left, Color: Black}
			result = t.Left
			return
		}
		result = t.Left.insert(data)
		return
	case cmpResult < 0:
		if t.Right.Data == nil { //如果右孩子的数据域为空，则插入到右孩子的位置
			t.Right.Data = data
			t.Right.Color = Red
			t.Right.Left = &RedBlackNode{Parent: t.Right, Color: Black}
			t.Right.Right = &RedBlackNode{Parent: t.Right, Color: Black}
			result = t.Right
			return
		}
		result = t.Right.insert(data)
		return
	default:
		//t.Color = Red
		result = t
		return
	}
}

//从某一个节点开始找到需要删除的data对应的节点，并返回
func (t *RedBlackNode) remove(data Element) (result *RedBlackNode) {
	cmpResult := t.Data.Compare(data)
	switch {
	case cmpResult > 0:
		if t.Left.Data == nil {
			return
		}
		result = t.Left.remove(data)
	case cmpResult < 0:
		if t.Right.Data == nil {
			return
		}
		result = t.Right.remove(data)
	default:
		//如果左子树不为空
		if t.Left.Data != nil {
			maxNode := maxNode(t.Left).(*RedBlackNode)
			t.Data = maxNode.Data
			result = maxNode
			return
		}
		//如果右子树不为空
		if t.Right.Data != nil {
			minNode := minNode(t.Right).(*RedBlackNode)
			t.Data = minNode.Data
			result = minNode
			return
		}
		//如果左右子树都为空, 则直接返回当前节点为需要删除的节点
		result = t
	}
	return
}

/*
	往红黑树中插入一个节点
	通常将需要插入的节点设置为红色（如果设为黑色，就会导致根到叶子的路径上有一条路上，多一个额外的黑节点，这个是很难调整的。
	但是设为红色节点后，可能会导致出现两个连续红色节点的冲突，那么可以通过颜色调换和树旋转来调整。）
*/

//如果插入的是根节点，则只需要将节点的颜色置为黑色即可
func insertNode(node *RedBlackNode) (root *RedBlackNode) {
	if node.Parent == nil {
		node.Color = Black
		return
	}
	//如果插入的节点的父节点为黑色，则不需要做调整。否则进入case3
	if node.Parent.Color != Black {
		root = insertCase(node)
	}
	return
}

/*
在下列情形中假定新节点的父节点为红色，所以它有祖父节点；因为如果父节点是根节点，那父节点就应当是黑色。
所以新节点总有一个叔父节点，尽管在情形4和5下它可能是叶子节点。

如果父节点和叔父节点二者都是红色，则可以将它们两个重绘为黑色并重绘祖父节点为红色。现在新节点有了一个
黑色的父节点。因为通过父节点或叔父节点的任何路径都必定通过祖父节点，在这些路径上的黑节点数目没有改变。
但是，红色的祖父节点可能是根节点，这就违反了性质2，也有可能祖父节点的父节点是红色的，这就违反了性质4。
为了解决这个问题，我们在祖父节点上递归地进行情形1的整个过程。（把祖父节点当成是新加入的节点进行各种情形的检查）
*/
func insertCase(node *RedBlackNode) (root *RedBlackNode) {
	if uncle := node.uncleNode(); uncle != nil && uncle.Color == Red {
		grandParent := node.grandParent()
		node.Parent.Color = Black
		uncle.Color = Black
		grandParent.Color = Red
		return insertNode(grandParent)
	}
	return insertCase1(node)
}

/*
在余下的情形下，假定父节点是其祖父的左子节点。如果它是右子节点，情形4和情形5中的左和右应当对调。
父节点是红色而叔父节点U是黑色或缺少，并且新节点是其父节点的右子节点而父节点又是其父节点的左子节点。
在这种情形下，我们进行一次左旋转调换新节点和其父节点的角色;接着，我们按情形5处理以前的父节点以解决
仍然失效的性质4。
*/
func insertCase1(node *RedBlackNode) (root *RedBlackNode) {
	grandParent := node.grandParent()
	if node == node.Parent.Right && node.Parent == grandParent.Left {
		leftRotate(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandParent.Right {
		rightRotate(node.Parent)
		node = node.Right
	}
	return insertCase2(node)
}

/*
父节点是红色而叔父节点是黑色或缺少，新节点是其父节点的左子节点，而父节点又是其父节点的左子节点。
在这种情形下，针对父节点进行一次右旋转；在旋转产生的树中，以前的父节点现在是新节点和以前的祖父
节点的父节点。我们知道以前的祖父节点是黑色，否则父节点就不可能是红色（如果父节点和祖父节点都是
红色就违反了性质4，所以祖父节点必须是黑色）。我们切换以前的父节点和祖父节点的颜色，结果的树满足性质
4。性质5也仍然保持满足，因为通过这三个节点中任何一个的所有路径以前都通过祖父节点，现在它们都通
过以前的父节点。在各自的情形下，这都是三个节点中唯一的黑色节点。
*/
func insertCase2(node *RedBlackNode) (root *RedBlackNode) {
	//fmt.Printf("%v\n", node.PreTraverse())
	grandParent := node.grandParent()
	node.Parent.Color = Black
	grandParent.Color = Red
	if node == node.Parent.Left && node.Parent == grandParent.Left {
		root = rightRotate(grandParent)
	} else {
		root = leftRotate(grandParent)
	}
	return
}

//删除node节点
func deleteNode(node *RedBlackNode) (root *RedBlackNode) {
	var child *RedBlackNode
	if node.Left.Data == nil {
		child = node.Right
	} else {
		child = node.Left
	}
	//如果删除的是没有子树的根节点
	if node.Parent == nil && node.Left.Data == nil && node.Right.Data == nil {
		node.Data = nil
		node.Color = Black
		root = node
		return
	}
	//将孩子节点提升, 这里已经将node删除
	if node.Parent.Left == node {
		node.Parent.Left = child
	} else {
		node.Parent.Right = child
	}
	child.Parent = node.Parent

	//如果node为红色，删除红色节点不会影响红黑树的性质，不需要调整
	//如果node是黑色，但孩子节点是红色，此时只需要将孩子节点改为黑色即可
	//如果node是黑色，孩子节点也是黑色，则少了一个黑色节点，需要对树进行调整，以达到平衡
	if node.Color == Black {
		if child.Color == Red { //如果删除的是黑色节点并且提升的孩子节点为红色，则只需要将孩子节点调色为黑色即可
			child.Color = Black
		} else { //否则需要其他操作
			root = deleteCase1(child)
		}
	}
	return
}

//被删除的是根，孩子节点作为新的根，否则进入case2
func deleteCase1(node *RedBlackNode) (root *RedBlackNode) {
	if node.Parent == nil {
		node.Color = Black
		return node
	}
	return deleteCase2(node)
}

//假定孩子节点的兄弟节点为红色，这时我们对被删除节点对孩子节点的新父节点做左旋转(或者右旋转)， 将红色兄弟节点转换为孩子节点的祖父节点，接着
//对调孩子的父亲和祖父的颜色，完成这两个操作后，尽管所有路径上黑色节点的数目没有改变，但现在孩子节点有了一个黑色的兄弟和一个红色的父亲（它的
// 新兄弟是黑色因为它是之前红色兄弟节点的一个儿子），所以我们可以接下去按情形4、情形5或情形6来处理
func deleteCase2(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()
	if siblingNode == nil {
		return
	}

	if siblingNode.Color == Red {
		node.Parent.Color = Red
		siblingNode.Color = Black
		if node == node.Parent.Left {
			leftRotate(node.Parent)
		} else {
			rightRotate(node.Parent)
		}
	}
	return deleteCase3(node)
}

// 孩子节点的父亲、兄弟节点和兄弟节点的儿子都是黑色的。在这种情形下，我们简单的重绘兄弟节点为红色。结果是通过兄弟节点的所有路径，它们就是
// 以前不通过孩子节点的那些路径，都少了一个黑色节点。因为删除N的初始的父亲使通过孩子节点的所有路径少了一个黑色节点，这使事情都平衡了起来。
// 但是，通过P的所有路径现在比不通过P的路径少了一个黑色节点，所以仍然违反性质5。要修正这个问题，我们要从情形1开始，在被删除节点上做重新平衡处理。
func deleteCase3(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()

	if (node.Parent.Color == Black) &&
		(siblingNode.Color == Black) &&
		(siblingNode.Left.Color == Black) &&
		(siblingNode.Right.Color) == Black {
		siblingNode.Color = Red
		return deleteCase1(node.Parent)
	} else {
		return deleteCase4(node)
	}
}

//兄弟节点和兄弟节点的儿子都是黑色，但是孩子节点的父亲是红色。在这种情形下，我们简单的交换孩子节点的兄弟和父亲的颜色。这不影响不通过孩子
// 节点的路径的黑色节点的数目，但是它在通过N的路径上对黑色节点数目增加了一，添补了在这些路径上删除的黑色节点。
func deleteCase4(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()

	if (node.Parent.Color == Red) &&
		(siblingNode.Color == Black) &&
		(siblingNode.Left.Color == Black) &&
		(siblingNode.Right.Color == Black) {
		siblingNode.Color = Red
		node.Parent.Color = Black
	} else {
		root = deleteCase5(node)
	}
	return
}

// 兄弟节点是黑色，兄弟节点的左儿子是红色，兄弟节点的右儿子是黑色，而孩子节点是它父亲的左儿子。在这种情形下我们在兄弟节点上做右旋转，这样
// 兄弟节点的左儿子成为兄弟节点的父亲和孩子节点的新兄弟。我们接着交换兄弟节点和它的新父亲的颜色。所有路径仍有同样数目的黑色节点，
// 但是现在孩子节点有了一个黑色兄弟，他的右儿子是红色的，所以我们进入了情形6。孩子节点和它的父亲都不受这个变换的影响。
func deleteCase5(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()

	if siblingNode.Color == Black {
		if (node == node.Parent.Left) &&
			(siblingNode.Right.Color == Black) &&
			(siblingNode.Left.Color == Red) {
			siblingNode.Color = Red
			siblingNode.Left.Color = Black
			rightRotate(siblingNode)
		} else if (node == node.Parent.Right) &&
			(siblingNode.Left.Color == Black) &&
			(siblingNode.Right.Color == Red) {
			siblingNode.Color = Red
			siblingNode.Right.Color = Black
			leftRotate(siblingNode)
		}
	}
	return deleteCase6(node)
}

//兄弟节点是黑色，兄弟节点的右儿子是红色，而孩子节点是它父亲的左儿子。在这种情形下我们在N的父亲上做左旋转，这样兄弟节点成为孩子节点的父亲
// 和兄弟节点的右儿子的父亲。我们接着交换N的父亲和兄弟节点的颜色，并使兄弟节点的右儿子为黑色。子树在它的根上的仍是同样的颜色，所以性质3没有被违反。
// 但是，孩子节点现在增加了一个黑色祖先：要么孩子节点的父亲变成黑色，要么它是黑色而兄弟节点被增加为一个黑色祖父。所以，通过孩子节点的路径都增加了一个黑色节点。
//此时，如果一个路径不通过孩子节点，则有两种可能性：
//
//它通过孩子节点的新兄弟。那么它以前和现在都必定通过兄弟节点和孩子节点的父亲，而它们只是交换了颜色。所以路径保持了同样数目的黑色节点。
//它通过孩子节点的新叔父，兄弟节点的右儿子。那么它以前通过兄弟节点、兄弟节点的父亲和兄弟节点的右儿子，但是现在只通过兄弟节点，它被假定
// 为它以前的父亲的颜色，和兄弟节点的右儿子，它被从红色改变为黑色。合成效果是这个路径通过了同样数目的黑色节点。
//在任何情况下，在这些路径上的黑色节点数目都没有改变。所以我们恢复了性质4。在示意图中的白色节点可以是红色或黑色，但是在变换前后都必须指定相同的颜色
func deleteCase6(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()

	siblingNode.Color = node.Parent.Color
	node.Parent.Color = Black

	if node == node.Parent.Left {
		siblingNode.Right.Color = Black
		root = leftRotate(node.Parent)
	} else {
		siblingNode.Left.Color = Black
		root = rightRotate(node.Parent)
	}
	return
}

//将某个节点进行旋转操作，结果返回涉及旋转的最上层节点
//左旋: 将某个节点向左旋转为其右孩子的左孩子
func leftRotate(node *RedBlackNode) (root *RedBlackNode) {
	parent := node.Parent
	right := node.Right

	node.Right = right.Left
	right.Left.Parent = node
	right.Left = node
	node.Parent = right
	right.Parent = parent
	if parent != nil {
		if parent.Left == node {
			parent.Left = right
		} else {
			parent.Right = right
		}
		root = parent
	} else {
		root = right
	}
	return
}

//右旋: 将某个节点向右旋转为其左孩子的右节点
func rightRotate(node *RedBlackNode) (root *RedBlackNode) {
	parent := node.Parent
	left := node.Left

	node.Left = left.Right
	left.Right.Parent = node
	left.Right = node
	node.Parent = left
	left.Parent = parent
	if parent != nil {
		if parent.Left == node {
			parent.Left = left
		} else {
			parent.Right = left
		}
		root = parent
	} else {
		root = left
	}
	return
}
func (r *RedBlackTree) AddNode(data Element) {
	var node = NewRedBlackNode(data)

	//作为root节点插入, 根节点为黑色
	if r.Root == nil {
		node.Color = Black
		r.Root = node
		return
	}
	result := r.Root.insert(data)

	//如果不是新插入的节点, 则不需要调整位置
	if result.Color == Black {
		return
	}
	//插入节点后根据结果判断是否需要调整节点颜色和位置
	root := insertNode(result)
	if root != nil && root.Parent == nil {
		r.Root = root
	}
}

func (r *RedBlackTree) RemoveNode(data Element) bool {
	if r.Root.Data == nil {
		return false
	}
	target := r.Root.remove(data)
	if target == nil {
		return false
	}
	root := deleteNode(target)
	if root != nil && root.Parent == nil {
		r.Root = root
	}
	return true
}

func (r *RedBlackTree) String() string {
	if r.Root == nil {
		return ""
	}
	return r.Root.PreTraverse()
}
