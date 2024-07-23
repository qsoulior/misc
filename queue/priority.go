package queue

import "container/heap"

type PriorityQueue[T any] interface {
	Len() int
	Front() (T, bool)
	PopFront() (T, bool)
	Push(value T, priority int) T
}

type priorityQueue[T any] struct {
	data prioritySlice[T]
}

func NewPriorityQueue[T any]() *priorityQueue[T] { return &priorityQueue[T]{make(prioritySlice[T], 0)} }

func (p priorityQueue[T]) Len() int { return p.data.Len() }

func (p priorityQueue[T]) Front() (T, bool) {
	if p.data.Len() > 0 {
		return p.data[0].value, true
	}

	var value T
	return value, false
}

func (p *priorityQueue[T]) PopFront() (T, bool) {
	if p.data.Len() > 0 {
		item := heap.Pop(&p.data).(*priorityItem[T])
		return item.value, true
	}

	var value T
	return value, false
}

func (p *priorityQueue[T]) Push(value T, priority int) T {
	item := &priorityItem[T]{
		value:    value,
		priority: priority,
	}
	heap.Push(&p.data, item)
	return item.value
}

type priorityItem[T any] struct {
	value    T
	priority int
	index    int
}

type prioritySlice[T any] []*priorityItem[T]

func (h prioritySlice[T]) Len() int { return len(h) }

func (h prioritySlice[T]) Less(i, j int) bool { return h[i].priority > h[j].priority }

func (h prioritySlice[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *prioritySlice[T]) Push(x any) {
	item := x.(*priorityItem[T])
	item.index = len(*h)
	*h = append(*h, item)
}

func (h *prioritySlice[T]) Pop() any {
	data := *h
	i := len(data) - 1
	item := data[i]
	item.index = -1
	data[i] = nil
	*h = data[:i]
	return item
}
