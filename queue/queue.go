package queue

import "github.com/qsoulior/alg/list"

// nodeValue retrieves value from list node.
// If node is nil, returns default value of type T and false as second value.
func nodeValue[T any](node *list.Node[T]) (T, bool) {
	if node != nil {
		return node.Value, true
	}

	var value T
	return value, false
}

// Queue interface.
type Queue[T any] interface {
	Len() int
	Front() (T, bool)
	Back() (T, bool)
	PopFront() (T, bool)
	PushBack(value T) T
}

// Queue based on linked list.
type queue[T any] struct {
	data list.List[T]
}

func NewQueue[T any]() Queue[T] { return &queue[T]{new(list.CircularLinkedList[T])} }

func (q queue[T]) Len() int { return q.data.Len() }

func (q queue[T]) Front() (T, bool) { return nodeValue(q.data.Front()) }

func (q queue[T]) Back() (T, bool) { return nodeValue(q.data.Back()) }

func (q *queue[T]) PopFront() (T, bool) { return nodeValue(q.data.PopFront()) }

func (q *queue[T]) PushBack(value T) T { return q.data.PushBack(value).Value }
