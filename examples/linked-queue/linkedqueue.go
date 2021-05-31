package main

import (
	"fmt"

	"github.com/sjquant/go-datastructure-algo/linkedqueue"
)

func main() {
	queue := linkedqueue.New()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	for {
		val, exists := queue.Pop()
		if exists {
			fmt.Println(val)

		} else {
			break
		}

	}
}
