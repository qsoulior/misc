package queue

import "github.com/qsoulior/alg/list"

// nodeValue retrieves value from list node and returns it.
// If node is nil, returns default value of type T and false as second value.
func nodeValue[T any](node *list.Node[T]) (T, bool) {
	if node != nil {
		return node.Value, true
	}

	var value T
	return value, false
}

// Queue represents abstract queue.
type Queue[T any] interface {
	// Len returns number of elements contained in queue.
	Len() int

	// Front returns first element of queue.
	// If queue is empty, it returns default value of type T and false as second value.
	Front() (T, bool)

	// Back returns last element of queue.
	// If queue is empty, it returns default value of type T and false as second value.
	Back() (T, bool)

	// PopFront removes first element from queue and returns it.
	// If queue is empty, it returns default value of type T and false as second value.
	PopFront() (T, bool)

	// PushBack inserts new value at back of queue.
	// It returns the inserted value.
	PushBack(value T) T
}

// listQueue implements queue based on linked list.
type listQueue[T any] struct {
	data list.List[T]
}

// NewListQueue returns new queue based on linked list.
func NewListQueue[T any]() Queue[T] { return &listQueue[T]{new(list.CircularLinkedList[T])} }

// Len returns number of elements contained in queue, O(1).
func (q listQueue[T]) Len() int { return q.data.Len() }

// Front returns first element of queue, O(1).
// If queue is empty, it returns default value of type T and false as second value.
func (q listQueue[T]) Front() (T, bool) { return nodeValue(q.data.Front()) }

// Back returns last element of queue, O(1).
// If queue is empty, it returns default value of type T and false as second value.
func (q listQueue[T]) Back() (T, bool) { return nodeValue(q.data.Back()) }

// PopFront removes first element from queue and returns it, O(1).
// If queue is empty, it returns default value of type T and false as second value.
func (q *listQueue[T]) PopFront() (T, bool) { return nodeValue(q.data.PopFront()) }

// PushBack inserts new value at back of queue, O(1).
// It returns the inserted value.
func (q *listQueue[T]) PushBack(value T) T { return q.data.PushBack(value).Value }
