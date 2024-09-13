package linkedlist

type node struct {
	Value int
	Next  *node
}

type LinkedList struct {
	head *node
	len  uint
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
