package queue

import "github.com/qsoulior/alg/list"

// Вспомогательная функция для обработки узла списка.
func nodeValue[T any](node *list.Node[T]) (T, bool) {
	if node != nil {
		return node.Value, true
	}

	// Если узел равен nil, возвращаем значение по умолчанию для типа T.
	var value T
	return value, false
}

// Обычная очередь, основанная на связном списке.
type Queue[T any] struct {
	data list.List[T]
}

// Конструктор очереди.
func NewQueue[T any]() *Queue[T] { return &Queue[T]{list.NewCircularLinkedList[T]()} }

func (q Queue[T]) Len() int { return q.data.Len() }

func (q Queue[T]) Front() (T, bool) { return nodeValue(q.data.Front()) }

func (q Queue[T]) Back() (T, bool) { return nodeValue(q.data.Back()) }

func (q *Queue[T]) PopFront() (T, bool) { return nodeValue(q.data.PopFront()) }

func (q *Queue[T]) PushBack(value T) T { return q.data.PushBack(value).Value }

// Двусторонняя очередь, основанная на связном списке.
type Deque[T any] struct{ *Queue[T] }

// Конструктор двусторонней очереди.
func NewDeque[T any]() *Deque[T] { return &Deque[T]{NewQueue[T]()} }

func (d *Deque[T]) PopBack() (T, bool) { return nodeValue(d.data.PopBack()) }

func (d *Deque[T]) PushFront(value T) T { return d.data.PushFront(value).Value }
