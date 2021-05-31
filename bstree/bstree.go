// Binary Search Tree
package bstree

import (
	"fmt"

	"github.com/sjquant/go-datastructure-algo/linkedqueue"
)

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
}

type BSTree struct {
	root *TreeNode
}

func New() BSTree {
	return BSTree{}
}

func (tree *BSTree) AddNode(node *TreeNode) {

	if tree.root == nil {
		tree.root = node
		return
	}

	n := tree.root
	for {
		if node.Val < n.Val {
			if n.left == nil {
				n.left = node
				break
			} else {
				n = n.left
			}
		} else {
			if n.right == nil {
				n.right = node
				break
			} else {
				n = n.right
			}

		}
	}
}

type depthNode struct {
	depth int
	node  *TreeNode
}

func (tree *BSTree) Print() {
	queue := linkedqueue.New()
	root := tree.root
	depth := 0
	queue.Push(depthNode{depth: depth, node: root})

	for !queue.IsEmpty() {
		item, exists := queue.Pop()
		dNode := item.(depthNode)
		if exists {
			if depth == 0 || dNode.depth != depth {
				fmt.Println()
				fmt.Printf("Depth %d: ", dNode.depth)
				depth = dNode.depth
			}
			fmt.Print(dNode.node.Val, " ")
		} else {
			break
		}

		if dNode.node.left != nil {
			queue.Push(depthNode{depth: depth + 1, node: dNode.node.left})
		}
		if dNode.node.right != nil {
			queue.Push(depthNode{depth: depth + 1, node: dNode.node.right})
		}
	}
}

func (tree *BSTree) Search(v int) (bool, int) {
	node := tree.root
	cnt := 1

	for {
		if node == nil {
			return false, cnt
		}

		if node.Val == v {
			return true, cnt
		}

		if v < node.Val {
			node = node.left
		} else {
			node = node.right
		}

		cnt += 1
	}
}
