package tree

/*
	前缀树，又称前缀树或字典树，其中的键通常是字符串。与二叉查找树不同，键不是直接保存在节点中，而是由节点在树中的位置决定。
	一个节点的所有子孙都有相同的前缀，也就是这个节点对应的字符串，而根节点对应空字符串。一般情况下，不是所有的节点都有对应的值，
	只有叶子节点和部分内部节点所对应的键才有相关的值。

	前缀树的性质：
	1. 根节点不包含字符，除根节点外每一个节点都只包含一个字符
	2. 从根节点到某一节点，路径上经过的字符连接起来，为该节点对应的字符串
	3. 每个节点的所有子节点包含的字符都不相同
*/

type Trie struct {
	KeyTag   bool           //节点所代表的字符串是否是完整的字符串
	Val      rune           //节点值，每个节点所包含的字符：考虑汉字的情况，这里用rune
	Children map[rune]*Trie //叶子节点，也可以用slice
}

func NewTrie() *Trie {
	t := &Trie{
		KeyTag:   false,
		Val:      -1,
		Children: make(map[rune]*Trie),
	}
	return t
}

func (t *Trie) init() {
	t.Children = make(map[rune]*Trie)
}

//向前缀树中插入字符串
func (t *Trie) Insert(word string) {
	var node = t
	for i := 0; i < len(word); i++ {
		if node.Children == nil {
			node.init()
		}
		if next := node.Children[rune(word[i])]; next == nil { //如果不存在，则插入
			next = &Trie{
				KeyTag:   i == len(word)-1,
				Val:      rune(word[i]),
				Children: nil,
			}
			node.Children[rune(word[i])] = next
			node = next
		} else {
			if i == len(word)-1 {
				next.KeyTag = true
				break
			}
			node = next
		}
	}
}

//搜索字符串
func (t *Trie) Search(word string) bool {
	var node = t
	for i := 0; i < len(word); i++ {
		if node.Children == nil {
			return false
		}
		if next := node.Children[rune(word[i])]; next == nil {
			return false
		} else {
			if i == len(word)-1 {
				return next.KeyTag
			}
			node = next
		}
	}
	return false
}

//判断是否有前缀为prefix的字符串
func (t *Trie) StartWith(prefix string) bool {
	var node = t
	for i := 0; i < len(prefix); i++ {
		if node.Children == nil {
			return false
		}
		if next := node.Children[rune(prefix[i])]; next == nil {
			return false
		} else {
			node = next
		}
	}
	return true
}

func (t *Trie) String() string {
	var result string
	if t.Val != -1 {
		result = string(t.Val)
	}
	if t.Children != nil {
		for _, v := range t.Children {
			result += v.String()
		}
	}
	return result
}
