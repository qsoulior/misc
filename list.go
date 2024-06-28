package alg

// Интерфейс связного списка.
type List[T any] interface {
	Len() int
	Front() *Node[T]
	Back() *Node[T]
	Pop(*Node[T]) *Node[T]
	PopFront() *Node[T]
	PopBack() *Node[T]
	InsertBefore(T, *Node[T]) *Node[T]
	InsertAfter(T, *Node[T]) *Node[T]
	PushFront(T) *Node[T]
	PushBack(T) *Node[T]
}

// Узел связного списка.
type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	list  List[T]
	Value T
}

// Получение предыдущего узла списка, работающее за O(1).
func (n *Node[T]) Prev() *Node[T] {
	// Если узел не принадлежит списку или циклично ссылается на последний узел,
	// возвращаем nil.
	if n.list == nil || n.prev == n.list.Back() {
		return nil
	}
	return n.prev
}

// Получение следующего узла списка, работающее за O(1).
func (n *Node[T]) Next() *Node[T] {
	// Если узел не принадлежит списку или циклично ссылается на первый узел,
	// возвращаем nil.
	if n.list == nil || n.next == n.list.Front() {
		return nil
	}
	return n.next
}

// Обычный двусвязный список.
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

// Конструктор двусвязного списка.
func NewLinkedList[T any]() *LinkedList[T] { return &LinkedList[T]{len: 0} }

// Получение длины списка, работающее за O(1).
func (l LinkedList[T]) Len() int { return l.len }

// Получение первого узла, работающее за O(1).
func (l LinkedList[T]) Front() *Node[T] { return l.head }

// Получение последнего узла, работающее за O(1).
func (l LinkedList[T]) Back() *Node[T] { return l.tail }

// Извлечение/удаление произвольного узла, работающее за O(1).
func (l *LinkedList[T]) Pop(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	// Проверяем принадлежность к списку.
	if node.list != l {
		return node
	}

	// Если узел первый, обновляем head,
	// иначе отвязываем от предыдущего узла.
	if node.prev == nil {
		l.head = l.head.next
	} else {
		node.prev.next = node.next
	}

	// Если узел последний, обновляем tail,
	// иначе отвязываем от следующего узла.
	if node.next == nil {
		l.tail = l.tail.prev
	} else {
		node.next.prev = node.prev
	}

	// Избегаем утечек памяти.
	node.prev = nil
	node.next = nil
	node.list = nil

	l.len--
	return node
}

// Извлечение/удаление первого узла, работающее за O(1).
func (l *LinkedList[T]) PopFront() *Node[T] { return l.Pop(l.head) }

// Извлечение/удаление последнего узла, работающее за O(1).
func (l *LinkedList[T]) PopBack() *Node[T] { return l.Pop(l.tail) }

// Вставка перед узлом, работающая за O(1).
func (l *LinkedList[T]) InsertBefore(value T, at *Node[T]) *Node[T] {
	if at == nil || at.list != l {
		return nil
	}

	node := &Node[T]{
		Value: value,
		list:  l,
		prev:  at.prev,
		next:  at,
	}

	// Если узел вставляется перед первым, обновляем head,
	// иначе связываем с предыдущим.
	if at.prev == nil {
		l.head = node
	} else {
		at.prev.next = node
	}

	// Связываем узел со следующим.
	at.prev = node

	l.len++
	return node
}

// Вставка после узла, работающая за O(1).
func (l *LinkedList[T]) InsertAfter(value T, at *Node[T]) *Node[T] {
	if at == nil || at.list != l {
		return nil
	}

	node := &Node[T]{
		Value: value,
		list:  l,
		prev:  at,
		next:  at.next,
	}

	// Если узел вставляется после последнего, обновляем tail,
	// иначе связываем со следующим.
	if at.next == nil {
		l.tail = node
	} else {
		at.next.prev = node
	}

	// Связываем узел с предыдущим.
	at.next = node

	l.len++
	return node
}

// Вставка в начало списка, работающая за O(1).
func (l *LinkedList[T]) PushFront(value T) *Node[T] {
	// Если список пустой, инициируем head и tail.
	if l.head == nil {
		node := &Node[T]{Value: value, list: l}
		l.head = node
		l.tail = node
		l.len++
		return node
	}

	return l.InsertBefore(value, l.head)
}

// Вставка в конец списка, работающая за O(1).
func (l *LinkedList[T]) PushBack(value T) *Node[T] {
	// Если список пустой, вставляем узел в начало списка.
	if l.tail == nil {
		return l.PushFront(value)
	}

	return l.InsertAfter(value, l.tail)
}

// Кольцевой двусвязный список.
type CircularLinkedList[T any] struct {
	head *Node[T]
	len  int
}

// Конструктор кольцевого двусвязного списка.
func NewCircularLinkedList[T any]() *CircularLinkedList[T] { return &CircularLinkedList[T]{len: 0} }

// Получение длины списка, работающее за O(1).
func (l CircularLinkedList[T]) Len() int { return l.len }

// Получение первого узла, работающее за O(1).
func (l CircularLinkedList[T]) Front() *Node[T] { return l.head }

// Получение последнего узла, работающее за O(1).
func (l CircularLinkedList[T]) Back() *Node[T] {
	if l.head == nil {
		return nil
	}
	return l.head.prev
}

// Извлечение/удаление произвольного узла, работающее за O(1).
func (l *CircularLinkedList[T]) Pop(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	// Проверяем принадлежность к списку.
	if node.list != l {
		return node
	}

	// Отвязываем от соседних узлов.
	node.prev.next = node.next
	node.next.prev = node.prev

	// Избегаем утечек памяти.
	node.prev = nil
	node.next = nil
	node.list = nil

	l.len--
	return node
}

// Извлечение/удаление первого узла, работающее за O(1).
func (l *CircularLinkedList[T]) PopFront() *Node[T] { return l.Pop(l.head) }

// Извлечение/удаление последнего узла, работающее за O(1).
func (l *CircularLinkedList[T]) PopBack() *Node[T] {
	if l.head == nil {
		return nil
	}
	return l.Pop(l.head.prev)
}

// Вставка перед узлом, работающая за O(1).
func (l *CircularLinkedList[T]) InsertBefore(value T, at *Node[T]) *Node[T] {
	if at == nil {
		return nil
	}

	return l.InsertAfter(value, at.prev)
}

// Вставка после узла, работающая за O(1).
func (l *CircularLinkedList[T]) InsertAfter(value T, at *Node[T]) *Node[T] {
	if at == nil || at.list != l {
		return nil
	}

	node := &Node[T]{
		Value: value,
		list:  l,
		prev:  at,
		next:  at.next,
	}

	// Связываем узел с предыдущим.
	at.next.prev = node //
	at.next = node

	l.len++
	return node
}

// Вставка в начало списка, работающая за O(1).
func (l *CircularLinkedList[T]) PushFront(value T) *Node[T] {
	// Если список пустой, инициируем head.
	if l.head == nil {
		node := &Node[T]{Value: value, list: l}
		node.prev = node
		node.next = node
		l.head = node
		l.len++
		return node
	}

	return l.InsertBefore(value, l.head)
}

// Вставка в конец списка, работающая за O(1).
func (l *CircularLinkedList[T]) PushBack(value T) *Node[T] {
	// Если список пустой, вставляем узел в начало списка.
	if l.head == nil {
		return l.PushFront(value)
	}

	return l.InsertAfter(value, l.head.prev)
}
