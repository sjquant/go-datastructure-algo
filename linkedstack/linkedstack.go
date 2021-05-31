package linkedstack

import "github.com/sjquant/go-datastructure-algo/doublylinkedlist"

type LinkedStack struct {
	list doublylinkedlist.DoublyLinkedList
}

func New() LinkedStack {
	return LinkedStack{list: doublylinkedlist.DoublyLinkedList{}}
}

func (stack *LinkedStack) Push(value interface{}) {
	node := &doublylinkedlist.Node{Val: value}
	stack.list.AddNode(node)
}

func (stack *LinkedStack) Pop() (interface{}, bool) {
	node := stack.list.PopTail()

	if node == nil {
		return 0, false
	} else {
		return node.Val, true
	}
}

func (stack *LinkedStack) IsEmpty() bool {
	return stack.list.Head == nil
}
