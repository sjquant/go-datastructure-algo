package main

import (
	"fmt"

	"github.com/sjquant/go-datastructure-algo/linkedstack"
)

func main() {
	stack := linkedstack.New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	for {
		val, exists := stack.Pop()
		if exists {
			fmt.Println(val)

		} else {
			break
		}

	}
}
