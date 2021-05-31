package singlylinkedlist

import "fmt"

type Node struct {
	Val  interface{}
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func (list *SinglyLinkedList) AddNode(node *Node) {
	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
	}
	list.tail = node
}

func (list *SinglyLinkedList) GetPrev(node *Node) *Node {
	n := list.head
	for n.next != node {
		n = n.next
	}
	return n
}

func (list *SinglyLinkedList) RemoveNode(node *Node) {
	if node == list.head {
		list.head = node.next
	} else {
		prev := list.GetPrev(node)
		prev.next = node.next
		if node == list.tail {
			list.tail = prev
		}
	}
}

func (list *SinglyLinkedList) PrintNodes() {
	node := list.head

	if node == nil {
		fmt.Print("[ ]")
		return
	}

	fmt.Print("[ ")
	for node.next != nil {
		fmt.Printf("%d, ", node.Val)
		node = node.next
	}

	fmt.Printf("%d", node.Val)
	fmt.Println(" ]")
}
