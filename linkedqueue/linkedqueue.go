package linkedqueue

import "github.com/sjquant/go-datastructure-algo/doublylinkedlist"

type LinkedQueue struct {
	list doublylinkedlist.DoublyLinkedList
}

func New() LinkedQueue {
	return LinkedQueue{list: doublylinkedlist.DoublyLinkedList{}}
}

func (queue *LinkedQueue) Push(value interface{}) {
	node := &doublylinkedlist.Node{Val: value}
	queue.list.AddNode(node)
}

func (queue *LinkedQueue) Pop() (interface{}, bool) {
	node := queue.list.PopHead()

	if node == nil {
		return 0, false
	} else {
		return node.Val, true
	}
}

func (queue *LinkedQueue) IsEmpty() bool {
	return queue.list.Head == nil
}
