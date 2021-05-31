package main

import "github.com/sjquant/go-datastructure-algo/singlylinkedlist"

func main() {
	var list singlylinkedlist.SinglyLinkedList
	list = singlylinkedlist.SinglyLinkedList{}

	n1 := &singlylinkedlist.Node{Val: 3}
	n2 := &singlylinkedlist.Node{Val: 5}
	n3 := &singlylinkedlist.Node{Val: 10}
	n4 := &singlylinkedlist.Node{Val: 15}
	list.AddNode(n1)
	list.AddNode(n2)
	list.AddNode(n3)
	list.AddNode(n4)
	list.PrintNodes()
	list.RemoveNode(n4)
	list.PrintNodes()
}
