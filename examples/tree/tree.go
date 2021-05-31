package main

import (
	"fmt"

	"github.com/sjquant/go-datastructure-algo/tree"
)

func main() {
	root := &tree.TreeNode{Val: 0}
	t := tree.New(root)

	n1 := &tree.TreeNode{Val: 1}
	n2 := &tree.TreeNode{Val: 2}
	n3 := &tree.TreeNode{Val: 3}
	n4 := &tree.TreeNode{Val: 4}
	n5 := &tree.TreeNode{Val: 5}
	n6 := &tree.TreeNode{Val: 6}
	n7 := &tree.TreeNode{Val: 7}
	n8 := &tree.TreeNode{Val: 8}
	n9 := &tree.TreeNode{Val: 9}
	n10 := &tree.TreeNode{Val: 10}

	root.AddChild(n1)
	root.AddChild(n2)
	n1.AddChild(n3)
	n1.AddChild(n4)
	n3.AddChild(n5)
	n5.AddChild(n6)
	n5.AddChild(n7)
	n2.AddChild(n8)
	n2.AddChild(n9)
	n8.AddChild(n10)

	fmt.Println("----------DFS----------")
	t.DFSwithRecursion()
	t.DFSwithStack()
	fmt.Println("----------BFS----------")
	t.BFS()
	fmt.Println()
}
