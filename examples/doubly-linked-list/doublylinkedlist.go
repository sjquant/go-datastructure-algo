package main

import "github.com/sjquant/go-datastructure-algo/doublylinkedlist"

func main() {
	var list doublylinkedlist.DoublyLinkedList
	list = doublylinkedlist.DoublyLinkedList{}

	n1 := &doublylinkedlist.Node{Val: 3}
	n2 := &doublylinkedlist.Node{Val: 5}
	n3 := &doublylinkedlist.Node{Val: 10}
	n4 := &doublylinkedlist.Node{Val: 15}
	list.AddNode(n1)
	list.AddNode(n2)
	list.AddNode(n3)
	list.AddNode(n4)
	list.PrintNodes()
	list.RemoveNode(n3)
	list.PrintNodes()
}
