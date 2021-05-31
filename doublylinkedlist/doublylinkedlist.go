package doublylinkedlist

import "fmt"

type Node struct {
	Val  interface{}
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (list *DoublyLinkedList) AddNode(node *Node) {
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
		node.Prev = list.Tail
		list.Tail = node
	}
}

func (list *DoublyLinkedList) PopHead() *Node {
	if list.Head == nil {
		return nil
	} else {
		node := list.Head
		list.RemoveNode(list.Head)
		return node
	}
}

func (list *DoublyLinkedList) PopTail() *Node {
	if list.Tail == nil {
		return nil
	} else {
		node := list.Tail
		list.RemoveNode(list.Tail)
		return node
	}
}

func (list *DoublyLinkedList) RemoveNode(node *Node) {
	if node == list.Head && node == list.Tail {
		list.Head = nil
		list.Tail = nil
	} else if node == list.Head {
		list.Head = node.Next
		node.Next.Prev = nil
	} else if node == list.Tail {
		list.Tail = node.Prev
		node.Prev.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
}

func (list *DoublyLinkedList) PrintNodes() {
	node := list.Head

	if node == nil {
		fmt.Print("[ ]")
		return
	}

	fmt.Print("[ ")
	for node.Next != nil {
		fmt.Printf("%d, ", node.Val)
		node = node.Next
	}

	fmt.Printf("%d", node.Val)
	fmt.Println(" ]")
}
