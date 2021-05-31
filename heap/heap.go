package heap

import "fmt"

type Heap struct {
	list []int
}

func New() Heap {
	return Heap{
		list: []int{},
	}
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Pop() (int, bool) {
	if len(h.list) == 0 {
		return 0, false
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	if len(h.list) == 0 {
		return top, true
	}

	h.list[0] = last
	idx := 0

	for idx < len(h.list) {
		leftIdx := 2*idx + 1
		rightIdx := 2*idx + 2

		if leftIdx >= len(h.list) {
			break
		}

		if rightIdx >= len(h.list) || h.list[leftIdx] >= h.list[rightIdx] {
			if h.list[leftIdx] > h.list[idx] {
				h.list[idx], h.list[leftIdx] = h.list[leftIdx], h.list[idx]
				idx = leftIdx
			} else {
				break
			}
		} else {
			h.list[idx], h.list[rightIdx] = h.list[rightIdx], h.list[idx]
			idx = rightIdx
		}
	}
	return top, true
}

func (h *Heap) Print() {
	fmt.Println(h.list)
}
