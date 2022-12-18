package tree

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
)

/*
	赫夫曼树
	给定N个权值作为N个叶子结点，构造一棵二叉树，若该树的带权路径长度达到最小，称这样的二叉树为最优二叉树，也称为哈夫曼树(Huffman Tree)。
	哈夫曼树是带权路径长度最短的树，权值较大的结点离根较近。
	哈夫曼树又称最优二叉树，是一种带权路径长度最短的二叉树。所谓树的带权路径长度，就是树中所有的叶结点的权值乘上其到根结点的路径长度（若
	根结点为0层，叶结点到根结点的路径长度为叶结点的层数）。树的路径长度是从树根到每一结点的路径长度之和，记为WPL=（W1*L1+W2*L2+W3*L3+...+Wn*Ln），
	N个权值Wi（i=1,2,...n）构成一棵有N个叶结点的二叉树，相应的叶结点的路径长度为Li（i=1,2,...n）。可以证明哈夫曼树的WPL是最小的。
	在计算机数据处理中，哈夫曼编码使用变长编码表对源符号（如文件中的一个字母）进行编码，其中变长编码表是通过一种评估来源符号出现机率的方法
	得到的，出现机率高的字母使用较短的编码，反之出现机率低的则使用较长的编码，这便使编码之后的字符串的平均长度、期望值降低，从而达到无损压
	缩数据的目的。
*/

type HuffmanNodeList []*HuffmanNode

func (list HuffmanNodeList) Len() int {
	return len(list)
}

func (list HuffmanNodeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list HuffmanNodeList) Less(i, j int) bool {
	return list[i].Weight < list[j].Weight
}

func (list HuffmanNodeList) String() string {
	var result string
	for _, l := range list {
		result += fmt.Sprintf("%v", l.Data) + ":" + strconv.Itoa(int(l.Weight)) + "->"
	}
	return result + "END"
}

// 赫夫曼节点，用于构成赫夫曼树
type HuffmanNode struct {
	Weight uint         //权重
	Data   interface{}  //数据
	Parent *HuffmanNode //父节点
	Left   *HuffmanNode //左孩子
	Right  *HuffmanNode //右孩子
}

func (h *HuffmanNode) PreTraverse() (result string) {
	stack := &Stack{
		List: list.New(),
	}

	node := h
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.Push(node)
			result += fmt.Sprintf("%v(%v)-->", node.Weight, node.Data)
			node = node.Left
		} else {
			node = stack.Pop().(*HuffmanNode)
			node = node.Right
		}
	}
	return
}

// 赫夫曼树结构，这里使用的interface作为源数据类型
type HuffmanTree struct {
	root    *HuffmanNode           //根节点
	leaf    HuffmanNodeList        //所有叶子节点（即数据对应的节点）
	src     map[interface{}]uint   //源数据，key为数据，value为权重
	codeSet map[interface{}]string //编码集，key为数据，value为通过构造赫夫曼树得到的数据的编码
}

// 给定一组字符及其权重的集合，初始化出一棵赫夫曼树
func NewHuffmanTree(src map[interface{}]uint) *HuffmanTree {
	var tree = &HuffmanTree{
		src: src,
	}
	tree.init()
	tree.build()
	tree.parse()
	return tree
}

func (h *HuffmanTree) String() string {
	if h.root == nil {
		return ""
	}
	return h.root.PreTraverse()
}

// 根据数据进行赫夫曼编码
func (h *HuffmanTree) Coding(target interface{}) (result string) {
	if target == nil {
		return
	}
	var s string
	switch t := target.(type) {
	case string:
		s = t
	case []byte:
		s = string(t)
	}
	for _, t := range s {
		v := string(t)
		if c, ok := h.codeSet[v]; !ok {
			panic("invalid code: " + v)
		} else {
			result += c
		}
	}
	return result
}

// 根据赫夫曼编码获取数据
func (h *HuffmanTree) UnCoding(target string) (result string) {
	node := h.root
	for i := 0; i < len(target); i++ {
		switch target[i] {
		case '0':
			node = node.Left
		case '1':
			node = node.Right
		}
		if node.Left == nil && node.Right == nil {
			result = result + node.Data.(string)
			node = h.root
		}
	}
	return
}

// 初始化所有叶子节点
func (h *HuffmanTree) init() {
	if len(h.src) <= 1 {
		panic("invalid src length.")
	}
	h.codeSet = make(map[interface{}]string)
	h.leaf = make(HuffmanNodeList, len(h.src))
	var i int
	for data, weight := range h.src {
		var node = &HuffmanNode{
			Weight: weight,
			Data:   data,
		}
		h.leaf[i] = node
		i++
	}
	//对leaf根据权值排序
	sort.Sort(h.leaf)
}

// 构造赫夫曼树
// src: key为data，value为权值
func (h *HuffmanTree) build() {
	nodeList := h.leaf
	//根据huffman树的规则构造赫夫曼树
	for nodeList.Len() > 1 {
		//1. 选取权值最小的两个node构造出第一个节点
		var temp = &HuffmanNode{
			Weight: nodeList[0].Weight + nodeList[1].Weight,
			Left:   nodeList[0],
			Right:  nodeList[1],
		}
		nodeList[0].Parent = temp
		nodeList[1].Parent = temp

		//2.将生成的新节点插入节点序列中
		nodeList = regroup(nodeList[2:], temp)
	}
	h.root = nodeList[0]
}

// 获取每个byte的编码，目的是为了下次需要编码的时候不用再次遍历树以获取每个byte的编码了
// 在赫夫曼树中的所有节点要么没有孩子节点，要么有两个孩子节点，不存在只有一个孩子节点的节点
// 此处的编码为由底至顶获取，也可以由顶至底的获取
func (h *HuffmanTree) parse() {
	if h.root == nil {
		return
	}
	var temp *HuffmanNode
	var code string
	for _, n := range h.leaf {
		temp = n
		for temp.Parent != nil {
			if temp == temp.Parent.Left {
				code = "0" + code
			} else {
				code = "1" + code
			}
			temp = temp.Parent
		}
		h.codeSet[n.Data] = code
		code = ""
	}
}

// 重组，将生成的节点放入既有的list，排序后返回，权值最小的始终在最前面
func regroup(src HuffmanNodeList, temp *HuffmanNode) HuffmanNodeList {
	//将temp添加进src，然后取出weight最小的一个
	length := len(src)
	result := make(HuffmanNodeList, len(src)+1)
	if length == 0 {
		result[0] = temp
		return result
	}
	if src[length-1].Weight <= temp.Weight {
		copy(result, src)
		result[length] = temp
		return result
	}
	for i := range src {
		if src[i].Weight <= temp.Weight {
			result[i] = src[i]
		} else {
			result[i] = temp
			copy(result[i+1:], src[i:])
			break
		}
	}
	return result
}
