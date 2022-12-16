package tree

import (
	"fmt"
	"math"
	"sort"
)

/*
	B树（B-tree）是一种自平衡的树，能够保持数据有序。这种数据结构能够让查找数据、顺序访问、插入数据及删除的动作，都在对数时间内完成。
	B树，概括来说是一个一般化的二叉查找树（binary search tree）一个节点可以拥有2个以上的子节点。与自平衡二叉查找树不同，B树适用于读写相
	对大的数据块的存储系统，例如磁盘。B树减少定位记录时所经历的中间过程，从而加快存取速度。B树这种数据结构可以用来描述外部存储。这种数据结构
	常被应用在数据库和文件系统的实现上。

	描述一颗B树时需要指定它的阶数，阶数表示了一个结点最多有多少个孩子结点，一般用字母m表示阶数。当m取2时，就是我们常见的二叉搜索树。
	一颗m阶的B树有如下性质：

	1. 每个结点最多有m个孩子结点（子树），最多有m-1个关键字。
	2. 如果根结点不是叶子结点，则根结点至少有2个子树，至少有一个关键字
	3. 除根结点外的所有非叶子结点至少有math.Ceil(m/2)个孩子结点，至少有math.Ceil(m/2)-1个关键字
	4. 所有叶子结点位于同一层
	5. 每个结点中的关键字都按照从小到大排序
*/

type KVList []*KeyData

func (k KVList) Len() int {
	return len(k)
}

func (k KVList) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k KVList) Less(i, j int) bool {
	return k[i].Key < k[j].Key
}

func (k KVList) String() (result string) {
	for _, v := range k {
		result += fmt.Sprintf("(%v, %v)", v.Key, v.Value)
	}
	return
}

// 节点中包含的关键字及其对应的数据：key+data
type KeyData struct {
	Key   int         // 关键字
	Value interface{} // 数据
}

// B树
type BTree struct {
	M    int    // 阶数
	Root *BNode // 根结点
}

func NewBTree(m int) *BTree {
	if m < 2 {
		panic("invalid m")
	}
	return &BTree{
		M:    m,
		Root: nil,
	}
}

// 便于打印调试
func (b *BTree) String() (result string) {
	return b.Root.String()
}

// 便于打印调试
func (b *BNode) String() (result string) {
	if b == nil {
		return
	}
	for _, v := range b.Data {
		result += fmt.Sprintf("%v->", v.Key)
	}
	for _, h := range b.Children {
		result += h.String()
	}
	return
}

// 插入关键字
func (b *BTree) Insert(key int, value interface{}) {
	kv := &KeyData{
		Key:   key,
		Value: value,
	}
	// 如果插入的是根结点
	if b.Root == nil {
		b.Root = &BNode{
			Data:   []*KeyData{kv},
			Parent: nil,
		}
		return
	}
	root := b.Root.insert(b.M, kv)
	for root.Parent != nil {
		root = root.Parent
	}
	b.Root = root
}

// 查询关键字所在的结点
func (b *BTree) Search(key int) *BNode {
	if b.Root == nil {
		return nil
	}
	return b.Root.search(key)
}

// 删除key对应的关键字
func (b *BTree) Remove(targetKey int) bool {
	if b.Root == nil {
		return false
	}

	return b.Root.remove(b.M, targetKey)
}

// ****************************************
// B树的结点
type BNode struct {
	Parent   *BNode     // 父节点
	Data     []*KeyData // 关键字信息
	Children []*BNode   // 孩子节点
}

// 移除targetKey对应的KeyData
func (b *BNode) remove(m int, targetKey int) bool {
	/*
		目标结点为关键字targetKey所在的结点
		相邻关键字：对于不在终端结点上的关键字，其相邻关键字为该关键字左子树中最大的关键字或者右子树中最小的关键字

		删除操作总体分两种：
		1. targetKey在叶子结点上
			a) 目标结点内的关键字数量大于math.Ceil(m/2)-1，这时删除不会破坏B树的性质，可以直接删除
			b) 目标结点内的关键字数量等于math.Ceil(m/2)-1，并且其左右兄弟结点中存在关键字数量大于math.Ceil(m/2)-1的结点，
			   则删除后需要向兄弟结点借关键字(将兄弟结点中的某个关键字提升到父结点，将父结点中的某个关键字下沉到目标结点)
			c) 目标结点内的关键字数量等于math.Ceil(m/2)-1，而兄弟结点中不存在关键字数量大于math.Ceil(m/2)-1的结点，则需要进行结点
			   合并从父结点中取一个关键字与兄弟结点合并，并将取出的结点从父结点中删除，同时更新父子关系(如果需要)

		2. targetKey不在叶子结点上
			a) targetKey存在关键字数量大于math.Ceil(m/2)-1结点的左子树或者右子树，在对应子树上找到该关键字的相邻关键字，并交换相邻关键字与目标
			   关键字，然后在替换后的位置上删除targetKey（此时已经转换成了在叶子结点上删除关键字）
			b) targetKey左右子树的关键字数量均等于math.Ceil(m/2)-1，则将这两个左右子树结点进行合并，然后删除targetKey，并调整关系
	*/

	// 找到需要删除的结点
	targetNode := b.search(targetKey)
	if targetNode == nil {
		return false
	}
	ceil := int(math.Ceil(float64(m)/float64(2))) - 1
	// 如果targetKey在叶子结点中，则直接移除
	if len(targetNode.Children) == 0 {
		return targetNode.removeAtLeaf(ceil, targetKey)
	}

	// 如果不在叶子结点中
	return targetNode.removeAtMid(ceil, targetKey)
}

// 如果targetKey在非叶子结点中
func (b *BNode) removeAtMid(ceil int, targetKey int) bool {
	targetKeyNode := b.getTargetKeyNode(targetKey)
	if targetKeyNode == nil {
		return false
	}
	// 找到targetKey对应的左右子树
	children := b.getChildrenByKey(targetKey)
	if len(children) == 0 {
		return false
	}
	// 从子树中找到相邻关键字所在的结点，并交换
	for _, c := range children {
		if len(c.Data) <= ceil {
			continue
		}
		// 如果是右子树
		if c.Data[0].Key > targetKey {
			minNode := c.minKeyNode()
			minKey := minNode.Data[0]
			b.removeKeyData(targetKey)
			b.insertKeyData(minKey)
			minNode.removeKeyData(minKey.Key)
			minNode.insertKeyData(targetKeyNode)
			return minNode.removeAtLeaf(ceil, targetKey)
		}
		// 如果是左子树
		if c.Data[len(c.Data)-1].Key < targetKey {
			maxNode := c.maxKeyNode()
			maxKey := maxNode.Data[len(maxNode.Data)-1]
			b.removeKeyData(targetKey)
			b.insertKeyData(maxKey)
			maxNode.removeKeyData(maxKey.Key)
			maxNode.insertKeyData(targetKeyNode)
			return maxNode.removeAtLeaf(ceil, targetKey)
		}
	}

	// 走到这里说明targetKey的两棵子树的关键字数量都为ceil，则将这两个左右子树结点进行合并，然后删除targetKey，并调整关系
	b.removeKeyData(targetKey)
	var newNode = &BNode{
		Parent: b,
	}
	var datas []*KeyData
	for _, c := range children {
		datas = append(datas, c.Data...)
		for _, h := range c.Children {
			newNode.insertChild(h)
		}
		b.removeChild(c)
	}
	sort.Sort(KVList(datas))

	newNode.Data = datas

	b.insertChild(newNode)

	// 因为子树的关键字都等于ceil，所以即使合并了也达不到分裂的条件
	// newNode.split(m)

	return true
}

// 如果targetKey在叶子结点中
func (b *BNode) removeAtLeaf(ceil int, targetKey int) bool {
	dataLen := len(b.Data)
	// 目标结点内的关键字数量大于math.Ceil(m/2)-1，这时删除不会破坏B树的性质，可以直接删除
	if dataLen > ceil {
		b.removeKeyData(targetKey)
		return true
	}

	if dataLen != ceil {
		return false
	}
	// 目标结点内的关键字数量等于math.Ceil(m/2)-1，并且其左右兄弟结点中存在关键字数量大于math.Ceil(m/2)-1的结点，
	// 则删除后需要向兄弟结点借关键字(将兄弟结点中的某个关键字提升到父结点，将父结点中的某个关键字下沉到目标结点)
	siblingNodes := b.siblingNode()
	if len(siblingNodes) == 0 { // 没有兄弟结点
		return false
	}

	parentKey := b.getParentKey()
	if parentKey == nil {
		return false
	}

	for _, s := range siblingNodes {
		if len(s.Data) <= ceil {
			continue
		}
		// 如果存在关键字数量大于ceil的兄弟结点
		// 如果是parentKey的右子树
		if s.Data[0].Key > parentKey.Key {
			b.removeKeyData(targetKey)            // 删除目标关键值
			b.insertKeyData(parentKey)            // 将父结点中的关键值放进目标结点
			b.Parent.removeKeyData(parentKey.Key) // 从父结点中移除parentKey
			b.Parent.insertKeyData(s.Data[0])     // 从兄弟结点借关键值放进父结点
			s.removeKeyData(s.Data[0].Key)        // 从兄弟结点中移除关键值
			return true
		}
		// 如果是parentKey的左子树
		if s.Data[len(s.Data)-1].Key < parentKey.Key {
			b.removeKeyData(targetKey)
			b.insertKeyData(parentKey)
			b.Parent.removeKeyData(parentKey.Key)
			b.Parent.insertKeyData(s.Data[len(s.Data)-1])
			s.removeKeyData(s.Data[len(s.Data)-1].Key)
			return true
		}
	}

	// 目标结点内的关键字数量等于math.Ceil(m/2)-1，而兄弟结点中不存在关键字数量大于math.Ceil(m/2)-1的结点，则需要进行结点
	// 合并: 从父结点中取一个关键字与兄弟结点合并，并将取出的结点从父结点中删除，同时更新父子关系(如果需要)
	b.removeKeyData(targetKey)            // 从当前结点中移除targetKey
	b.Parent.removeKeyData(parentKey.Key) // 从父结点中移除parentKey
	b.Parent.removeChild(b)               // 从父结点中移除当前结点
	b.Parent.removeChild(siblingNodes[0]) // 从父结点中移除被合并的兄弟结点

	// 利用parentKey、兄弟结点中的key以及本结点剩余的key重新组成一个结点
	newNode := &BNode{
		Parent: b.Parent,
		Data:   []*KeyData{parentKey},
	}
	newNode.Data = append(newNode.Data, b.Data...)
	newNode.Data = append(newNode.Data, siblingNodes[0].Data...)
	sort.Sort(KVList(newNode.Data))

	// 将新结点插入父结点中去
	b.Parent.insertChild(newNode)
	return true
}

// 搜索key对应的结点
func (b *BNode) search(targetKey int) *BNode {
	for i, v := range b.Data {
		if v.Key == targetKey {
			return b
		}
		if v.Key > targetKey && len(b.Children) > 0 {
			return b.Children[i].search(targetKey)
		}
	}
	// 如果b中所有key都比targetKey小，则在b的最后一个孩子结点中搜索，否则表示没有targetKey对应的结点
	childCnt := len(b.Children)
	if childCnt == 0 {
		return nil
	}

	return b.Children[childCnt-1].search(targetKey)
}

// 插入的位置一定是叶子结点
func (b *BNode) insert(m int, keyData *KeyData) (root *BNode) {
	// 1. 如果b是叶子结点，则直接插入到data中，然后根据data的长度来判断是否需要分裂
	if len(b.Children) == 0 {
		// 插入
		b.insertKeyData(keyData)
		// 判断是否需要分裂
		root = b.split(m)
		return
	}
	// 2. 如果b不是叶子节点，则需要先找到插入的位置
	var leaf *BNode // 插入的位置
	for i, d := range b.Data {
		if keyData.Key < d.Key {
			leaf = b.Children[i]
			break
		}
	}

	if leaf == nil {
		leaf = b.Children[len(b.Children)-1]
	}
	root = leaf.insert(m, keyData)
	return
}

// 执行分裂
func (b *BNode) split(m int) (root *BNode) {
	var dataLen = len(b.Data)
	if dataLen < m {
		return b
	}
	var mid int
	if len(b.Data)%2 != 0 {
		mid = dataLen / 2
	} else {
		mid = dataLen/2 - 1
	}
	// 晋升的节点
	promoteKey := b.Data[mid]
	parent := b.Parent
	// 分裂操作步骤为:
	// 1. 将promoteKey从当前节点中剔除，KeyData分成大小两部分，生成两个新的节点，并重新分配孩子结点
	smallerKey := b.Data[:mid]
	biggerKey := b.Data[mid+1:]
	smallerNode := &BNode{
		Parent: parent,
		Data:   smallerKey,
	}
	biggerNode := &BNode{
		Parent: parent,
		Data:   biggerKey,
	}
	// 这里需要重新分配孩子节点，以promoteKey为中心，小的放在smallerNode里面，其他的放在biggerNode里面，并绑定父子结点关系
	for i := range b.Children {
		if i <= mid {
			b.Children[i].Parent = smallerNode
			smallerNode.Children = append(smallerNode.Children, b.Children[i])
		} else {
			b.Children[i].Parent = biggerNode
			biggerNode.Children = append(biggerNode.Children, b.Children[i])
		}
	}
	// 2. 将promoteKey晋升到父节点中去，如果父节点不存在则需要新生成一个父节点
	if parent == nil {
		parent = &BNode{
			Parent:   nil,
			Data:     []*KeyData{},
			Children: []*BNode{smallerNode, biggerNode},
		}
		smallerNode.Parent = parent
		biggerNode.Parent = parent
	} else {
		// 从父结点中删除当前节点，并且在父节点中找到步骤一中新生成的两个节点应该插入的位置
		parent.removeChild(b)
		parent.insertChild(smallerNode)
		parent.insertChild(biggerNode)
	}

	// 3. 将晋升的关键字插入父结点
	parent.insertKeyData(promoteKey)

	// 4. 然后递归的对父节点执行分裂
	root = parent.split(m)
	return
}

// 插入KeyData
func (b *BNode) insertKeyData(target *KeyData) {
	for i, v := range b.Data {
		// 如果key相等，则直接替换value
		if v.Key == target.Key {
			b.Data[i].Value = target.Value
			return
		}
		// 如果不存在相等的key，则插入对应的位置
		if v.Key > target.Key {
			var temp = make([]*KeyData, len(b.Data)+1)
			copy(temp[:i], b.Data[:i])
			temp[i] = target
			copy(temp[i+1:], b.Data[i:])
			b.Data = temp
			return
		}
	}
	// 执行到这里证明b.Data为空或者target是最大的，直接append
	b.Data = append(b.Data, target)
}

// 移除key对应的KeyData
func (b *BNode) removeKeyData(key int) {
	for i, v := range b.Data {
		if v.Key == key {
			b.Data = b.Data[:i+copy(b.Data[i:], b.Data[i+1:])]
			break
		}
	}
}

// 插入孩子结点
func (b *BNode) insertChild(child *BNode) {
	for i, v := range b.Children {
		if v.Data[0].Key > child.Data[len(child.Data)-1].Key {
			var temp = make([]*BNode, len(b.Children)+1)
			copy(temp[:i], b.Children[:i])
			temp[i] = child
			copy(temp[i+1:], b.Children[i:])
			b.Children = temp
			return
		}
	}
	b.Children = append(b.Children, child)
}

// 移除孩子结点
func (b *BNode) removeChild(child *BNode) {
	for i, v := range b.Children {
		if v == child {
			b.Children = append(b.Children[:i], b.Children[i+1:]...)
			break
		}
	}
}

// 获取一个结点的左右兄弟结点
func (b *BNode) siblingNode() []*BNode {
	childCnt := len(b.Parent.Children)
	if childCnt <= 1 {
		return nil
	}
	var result []*BNode
	for i, h := range b.Parent.Children {
		if h != b {
			continue
		}
		if i > 0 { // 左兄弟
			result = append(result, b.Parent.Children[i-1])
		}
		if i < childCnt-1 {
			result = append(result, b.Parent.Children[i+1])
		}
	}
	return result
}

// 获取一个结点中某个关键字的左右子树
func (b *BNode) getChildrenByKey(key int) []*BNode {
	childLen := len(b.Children)
	if childLen == 0 {
		return nil
	}
	var result []*BNode
	for i, v := range b.Data {
		if v.Key != key {
			continue
		}
		if i == 0 {
			result = append(result, b.Children[0], b.Children[1])
		} else if i == len(b.Data)-1 {
			result = append(result, b.Children[childLen-1], b.Children[childLen-2])
		} else {
			result = append(result, b.Children[i], b.Children[i+1])
		}
		break
	}
	return result
}

// 获取结点在父结点中对应的KeyData,即该KeyData的孩子中包含b
func (b *BNode) getParentKey() *KeyData {
	if b.Parent == nil {
		return nil
	}
	dataLen := len(b.Parent.Data)
	if dataLen == 0 {
		return nil
	}
	for _, d := range b.Parent.Data {
		if d.Key > b.Data[len(b.Data)-1].Key {
			return d
		}
	}
	// 如果父结点中的关键字都比b中的关键字小，则直接返回父结点中的最后一个关键字
	return b.Parent.Data[dataLen-1]
}

// 找到targetKey对应的KeyData
func (b *BNode) getTargetKeyNode(targetKey int) *KeyData {
	for _, v := range b.Data {
		if v.Key == targetKey {
			return v
		}
	}
	return nil
}

// 找到某个子树中最小关键字所在的结点
func (b *BNode) minKeyNode() *BNode {
	node := b
	for len(node.Children) > 0 {
		node = node.Children[0]
	}
	return node
}

// 找到某个子树中最大关键字所在的结点
func (b *BNode) maxKeyNode() *BNode {
	node := b
	for len(node.Children) > 0 {
		node = node.Children[len(node.Children)-1]
	}
	return node
}
