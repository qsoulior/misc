package queue

import "github.com/qsoulior/misc/list"

// Deque represents abstract double-ended queue.
type Deque[T any] interface {
	Queue[T]
	// PopBack removes last element from queue and returns it.
	// If queue is empty, it returns default value of type T and false as second value.
	PopBack() (T, bool)
	// PushFront inserts new value at front of queue.
	// It returns the inserted value.
	PushFront(value T) T
}

// listDeque implements double-ended queue based on linked list.
type listDeque[T any] struct{ *listQueue[T] }

// NewListDeque returns new deque based on linked list.
func NewListDeque[T any]() Deque[T] {
	return &listDeque[T]{&listQueue[T]{new(list.CircularLinkedList[T])}}
}

// PopBack removes last element from queue and returns it, O(1).
// If queue is empty, it returns default value of type T and false as second value.
func (d *listDeque[T]) PopBack() (T, bool) { return nodeValue(d.data.PopBack()) }

// PushFront inserts new value at front of queue, O(1).
// It returns the inserted value.
func (d *listDeque[T]) PushFront(value T) T { return d.data.PushFront(value).Value }
