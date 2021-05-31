package main

import (
	"fmt"

	"github.com/sjquant/go-datastructure-algo/arrayqueue"
)

func main() {
	queue := arrayqueue.New()
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	queue.Push(6)

	for {
		val, exists := queue.Pop()
		if exists {
			fmt.Println(val)

		} else {
			break
		}

	}
}
