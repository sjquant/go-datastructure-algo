package main

import (
	"fmt"

	"github.com/sjquant/go-datastructure-algo/arraystack"
)

func main() {
	stack := arraystack.New()
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	for {
		val, exists := stack.Pop()
		if exists {
			fmt.Println(val)

		} else {
			break
		}

	}
}
