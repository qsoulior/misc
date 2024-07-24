package queue

import "container/heap"

// Priority queue interface.
type PriorityQueue[T any] interface {
	Len() int
	Front() (T, int, bool)
	PopFront() (T, int, bool)
	Push(value T, priority int) T
}

// Max priority queue based on container/heap.
type priorityQueue[T any] struct {
	data prioritySlice[T]
}

func NewPriorityQueue[T any]() PriorityQueue[T] { return &priorityQueue[T]{make(prioritySlice[T], 0)} }

func (p priorityQueue[T]) Len() int { return p.data.Len() }

func (p priorityQueue[T]) Front() (T, int, bool) {
	if p.data.Len() > 0 {
		item := p.data[0]
		return item.value, item.priority, true
	}

	var value T
	return value, 0, false
}

func (p *priorityQueue[T]) PopFront() (T, int, bool) {
	if p.data.Len() > 0 {
		item := heap.Pop(&p.data).(*priorityItem[T])
		return item.value, item.priority, true
	}

	var value T
	return value, 0, false
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
}

type prioritySlice[T any] []*priorityItem[T]

func (h prioritySlice[T]) Len() int { return len(h) }

func (h prioritySlice[T]) Less(i, j int) bool { return h[i].priority > h[j].priority }

func (h prioritySlice[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *prioritySlice[T]) Push(x any) { *h = append(*h, x.(*priorityItem[T])) }

func (h *prioritySlice[T]) Pop() any {
	data := *h
	i := len(data) - 1
	item := data[i]
	data[i] = nil
	*h = data[:i]
	return item
}
