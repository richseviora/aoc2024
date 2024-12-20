package main

import "container/heap"

// HeapQueue is near copy of https://github.com/lizthegrey/adventofcode/blob/main/2022/heapq/heapqueue.go
type HeapQueue[T comparable] struct {
	elements  *[]T
	score     map[T]int
	positions map[T]int
}

func NewHeapQueue[T comparable]() *HeapQueue[T] {
	return &HeapQueue[T]{
		elements:  &[]T{},
		score:     make(map[T]int),
		positions: make(map[T]int),
	}
}

func (h *HeapQueue[T]) Swap(i, j int) {
	firstElement := (*h.elements)[i]
	secondElement := (*h.elements)[j]
	h.positions[firstElement], h.positions[secondElement] = h.positions[secondElement], h.positions[firstElement]
	(*h.elements)[i], (*h.elements)[j] = secondElement, firstElement
}

func (h *HeapQueue[T]) Len() int {
	return len(*h.elements)
}

func (h *HeapQueue[T]) Push(x interface{}) {
	cast, ok := x.(T)
	if !ok {
		panic("cannot cast")
	}
	h.positions[cast] = len(*h.elements)
	*h.elements = append(*h.elements, cast)
}

func (h *HeapQueue[T]) Pop() any {
	old := *h.elements
	n := len(old)
	if n == 0 {
		return nil
	}
	x := old[n-1]
	*h.elements = old[0 : n-1]
	delete(h.positions, x)
	return x
}

func (h *HeapQueue[T]) Less(i, j int) bool {
	return h.score[(*h.elements)[i]] < h.score[(*h.elements)[j]]
}

func (h *HeapQueue[T]) Upsert(n T, score int) {
	h.score[n] = score
	if pos, ok := h.positions[n]; ok {
		heap.Fix(h, pos)
	} else {
		heap.Push(h, n)
	}
}
