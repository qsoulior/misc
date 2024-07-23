package queue

// Интерфейс двусторонней очереди.
type Deque[T any] interface {
	Queue[T]
	PopBack() (T, bool)
	PushFront(value T) T
}

// Двусторонняя очередь, основанная на связном списке.
type deque[T any] struct{ *queue[T] }

// Конструктор двусторонней очереди.
func NewDeque[T any]() *deque[T] { return &deque[T]{NewQueue[T]()} }

func (d *deque[T]) PopBack() (T, bool) { return nodeValue(d.data.PopBack()) }

func (d *deque[T]) PushFront(value T) T { return d.data.PushFront(value).Value }
