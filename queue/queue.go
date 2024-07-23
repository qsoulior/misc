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

// Интерфейс обычной очереди.
type Queue[T any] interface {
	Len() int
	Front() (T, bool)
	Back() (T, bool)
	PopFront() (T, bool)
	PushBack(value T) T
}

// Обычная очередь, основанная на связном списке.
type queue[T any] struct {
	data list.List[T]
}

// Конструктор очереди.
func NewQueue[T any]() *queue[T] { return &queue[T]{new(list.CircularLinkedList[T])} }

func (q queue[T]) Len() int { return q.data.Len() }

func (q queue[T]) Front() (T, bool) { return nodeValue(q.data.Front()) }

func (q queue[T]) Back() (T, bool) { return nodeValue(q.data.Back()) }

func (q *queue[T]) PopFront() (T, bool) { return nodeValue(q.data.PopFront()) }

func (q *queue[T]) PushBack(value T) T { return q.data.PushBack(value).Value }
