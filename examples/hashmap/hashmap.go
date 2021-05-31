package main

import (
	"fmt"

	"github.com/sjquant/go-datastructure-algo/hashmap"
)

func main() {
	h := hashmap.New()
	h.Add("A", "1")
	h.Add("B", "2")
	h.Add("C", "3")
	h.Add("D", "4")
	h.Add("E", "5")
	h.Add("F", "6")
	h.Add("G", "7")

	var v string
	var exists bool

	v, exists = h.Get("A")
	fmt.Println(v, exists)
	v, exists = h.Get("B")
	fmt.Println(v, exists)
	v, exists = h.Get("What")
	fmt.Println(v, exists)
	v, exists = h.Pop("C")
	fmt.Println(v, exists)
	v, exists = h.Get("C")
	fmt.Println(v, exists)
}
