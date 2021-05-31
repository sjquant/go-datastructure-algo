package main

import (
	"fmt"

	"github.com/sjquant/go-datastructure-algo/bstree"
)

func main() {
	t := bstree.New()

	n0 := &bstree.TreeNode{Val: 5}
	n1 := &bstree.TreeNode{Val: 3}
	n2 := &bstree.TreeNode{Val: 2}
	n3 := &bstree.TreeNode{Val: 4}
	n4 := &bstree.TreeNode{Val: 8}
	n5 := &bstree.TreeNode{Val: 7}
	n6 := &bstree.TreeNode{Val: 6}
	n7 := &bstree.TreeNode{Val: 10}
	n8 := &bstree.TreeNode{Val: 9}

	t.AddNode(n0)
	t.AddNode(n1)
	t.AddNode(n2)
	t.AddNode(n3)
	t.AddNode(n4)
	t.AddNode(n5)
	t.AddNode(n6)
	t.AddNode(n7)
	t.AddNode(n8)

	t.Print()

	if found, cnt := t.Search(6); found {
		fmt.Println("found 6 cnt:", cnt)
	} else {
		fmt.Println("Not found 6 cnt:", cnt)
	}

	if found, cnt := t.Search(11); found {
		fmt.Println("fount 11 cnt:", cnt)
	} else {
		fmt.Println("Not found 11 cnt:", cnt)
	}
}
