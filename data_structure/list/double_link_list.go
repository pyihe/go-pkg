package list

type dulNode struct {
	data  interface{}
	prior *dulNode // 指向前驱节点
	next  *dulNode // 指向后继节点
}

type doubleLinkedList struct {
	length int
	head   *dulNode
}
