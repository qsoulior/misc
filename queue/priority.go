package queue

import "container/heap"

// PriorityQueue represents abstract priority queue.
type PriorityQueue[T any] interface {
	// Len returns number of elements contained in queue.
	Len() int
	// Front returns the highest-priority element and its priority.
	// If queue is empty, it returns default value of type T and false as third value.
	Front() (T, int, bool)
	// PopFront removes the highest-priority element, returns it and its priority.
	// If queue is empty, it returns default value of type T and false as third value.
	PopFront() (T, int, bool)
	// Push inserts new value with priority into queue.
	// It returns the inserted value and its priority.
	Push(value T, priority int) (T, int)
}

// priorityQueue implements priority queue based on min/max heap.
// Element with the highest priority has min/max value in heap.
type priorityQueue[T any] struct {
	data prioritySlice[T]
}

// NewMaxPriorityQueue returns new priority queue based on min heap.
func NewMinPriorityQueue[T any]() PriorityQueue[T] {
	return &priorityQueue[T]{new(minPrioritySlice[T])}
}

// NewMaxPriorityQueue returns new priority queue based on max heap.
func NewMaxPriorityQueue[T any]() PriorityQueue[T] {
	return &priorityQueue[T]{new(maxPrioritySlice[T])}
}

// Len returns number of elements contained in queue, O(1).
func (p priorityQueue[T]) Len() int { return p.data.Len() }

// Front returns element of queue that has the highest priority in heap, O(1).
// It also returns element's priority as second value.
// If queue is empty, it returns default value of type T and false as third value.
func (p priorityQueue[T]) Front() (T, int, bool) {
	if p.data.Len() > 0 {
		item := p.data.First()
		return item.value, item.priority, true
	}

	var value T
	return value, 0, false
}

// PopFront removes element of queue that has the highest priority in heap, O(log(n)).
// It returns this element and its priority as second value.
// If queue is empty, it returns default value of type T and false as third value.
func (p *priorityQueue[T]) PopFront() (T, int, bool) {
	if p.data.Len() > 0 {
		item := heap.Pop(p.data).(*PriorityItem[T])
		return item.value, item.priority, true
	}

	var value T
	return value, 0, false
}

// Push inserts new value with priority into queue, O(log(n)).
// It returns the inserted value and its priority.
func (p *priorityQueue[T]) Push(value T, priority int) (T, int) {
	item := &PriorityItem[T]{
		value:    value,
		priority: priority,
	}
	heap.Push(p.data, item)
	return item.value, priority
}

// PriorityItem implements a priority heap item.
type PriorityItem[T any] struct {
	value    T
	priority int
}

// prioritySlice represents heap based on slice.
// It includes heap.Interface to use container/heap operations.
type prioritySlice[T any] interface {
	heap.Interface
	// First returns first item of priority slice.
	First() *PriorityItem[T]
}

// minPrioritySlice implements priority slice
// in which the first item has minimum priority value.
type minPrioritySlice[T any] []*PriorityItem[T]

// First returns first item of priority slice.
func (h minPrioritySlice[T]) First() *PriorityItem[T] { return h[0] }

// Len returns number of items contained in priority slice.
func (h minPrioritySlice[T]) Len() int { return len(h) }

// Less returns true, if priority of item with index i
// is less than priority of item with index j.
func (h minPrioritySlice[T]) Less(i, j int) bool { return h[i].priority < h[j].priority }

// Swap swaps priority items with indexes i and j.
func (h minPrioritySlice[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push inserts new priority item with value x at end of slice.
func (h *minPrioritySlice[T]) Push(x any) { *h = append(*h, x.(*PriorityItem[T])) }

// Pop removes last item from priority slice and returns it.
func (h *minPrioritySlice[T]) Pop() any {
	data := *h
	i := len(data) - 1
	item := data[i]
	data[i] = nil
	*h = data[:i]
	return item
}

// maxPrioritySlice implements priority slice
// in which the first item has maximum priority value.
type maxPrioritySlice[T any] struct{ minPrioritySlice[T] }

// Less returns true, if priority of item with index i
// is greater than priority of item with index j.
func (h maxPrioritySlice[T]) Less(i, j int) bool { return h.minPrioritySlice.Less(j, i) }
