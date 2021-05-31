package tree

import (
	"fmt"

	"github.com/sjquant/go-datastructure-algo/linkedqueue"
	"github.com/sjquant/go-datastructure-algo/linkedstack"
)

type TreeNode struct {
	Val      int
	Children []*TreeNode
}

type Tree struct {
	root *TreeNode
}

func New(root *TreeNode) Tree {
	return Tree{root: root}
}

func (parent *TreeNode) AddChild(child *TreeNode) {
	parent.Children = append(parent.Children, child)
}

func dfsRecursion(node *TreeNode) {
	fmt.Printf("%d -> ", node.Val)

	for _, each := range node.Children {
		dfsRecursion(each)
	}
}

func (tree *Tree) DFSwithRecursion() {
	dfsRecursion(tree.root)
	fmt.Println()
}

func (tree *Tree) DFSwithStack() {
	stack := linkedstack.New()
	root := tree.root
	stack.Push(root)

	for !stack.IsEmpty() {
		tempNode, exists := stack.Pop()
		node := tempNode.(*TreeNode)
		if exists {
			fmt.Printf("%d -> ", node.Val)
		} else {
			break
		}

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.Push(node.Children[i])
		}
	}
	fmt.Println()
}

func (tree *Tree) BFS() {
	queue := linkedqueue.New()
	root := tree.root
	queue.Push(root)

	for !queue.IsEmpty() {
		tempNode, exists := queue.Pop()
		node := tempNode.(*TreeNode)
		if exists {
			fmt.Printf("%d -> ", node.Val)
		} else {
			break
		}

		for i := 0; i < len(node.Children); i++ {
			queue.Push(node.Children[i])
		}
	}
}
