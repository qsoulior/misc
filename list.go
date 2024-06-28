package alg

// Узел связного списка.
type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	list  *DoublyLinkedList[T]
	Value T
}

// Получение предыдущего узла списка, работающее за O(1).
func (n *Node[T]) Prev() *Node[T] {
	if n.list == nil {
		return nil
	}
	return n.prev
}

// Получение следующего узла списка, работающее за O(1).
func (n *Node[T]) Next() *Node[T] {
	if n.list == nil {
		return nil
	}
	return n.next
}

// Обычный двусвязный список.
type DoublyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

// Конструктор двусвязного списка.
func NewList[T any]() *DoublyLinkedList[T] { return &DoublyLinkedList[T]{len: 0} }

// Получение длины списка, работающее за O(1).
func (l DoublyLinkedList[T]) Len() int { return l.len }

// Получение первого узла, работающее за O(1).
func (l DoublyLinkedList[T]) Front() *Node[T] { return l.head }

// Получение последнего узла, работающее за O(1).
func (l DoublyLinkedList[T]) Back() *Node[T] { return l.tail }

// Извлечение/удаление произвольного узла, работающее за O(1).
func (l *DoublyLinkedList[T]) Pop(node *Node[T]) *Node[T] {
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
func (l *DoublyLinkedList[T]) PopFront() *Node[T] {
	return l.Pop(l.head)
}

// Извлечение/удаление последнего узла, работающее за O(1).
func (l *DoublyLinkedList[T]) PopBack() *Node[T] {
	return l.Pop(l.tail)
}

// Вставка перед узлом, работающая за O(1).
func (l *DoublyLinkedList[T]) InsertBefore(value T, at *Node[T]) *Node[T] {
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
func (l *DoublyLinkedList[T]) InsertAfter(value T, at *Node[T]) *Node[T] {
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
func (l *DoublyLinkedList[T]) PushFront(value T) *Node[T] {
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
func (l *DoublyLinkedList[T]) PushBack(value T) *Node[T] {
	// Если список пустой, вставляем узел в начало списка.
	if l.tail == nil {
		return l.PushFront(value)
	}

	return l.InsertAfter(value, l.tail)
}
