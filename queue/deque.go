package queue

import "github.com/qsoulior/alg/list"

// Double-ended queue interface.
type Deque[T any] interface {
	Queue[T]
	PopBack() (T, bool)
	PushFront(value T) T
}

// Double-ended queue based on linked list.
type deque[T any] struct{ *queue[T] }

func NewDeque[T any]() Deque[T] { return &deque[T]{&queue[T]{new(list.CircularLinkedList[T])}} }

func (d *deque[T]) PopBack() (T, bool) { return nodeValue(d.data.PopBack()) }

func (d *deque[T]) PushFront(value T) T { return d.data.PushFront(value).Value }
