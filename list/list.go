// Package list implements list data structures.
// It provides doubly linked list and circular doubly linked list implementations.
package list

// List represents abstract list.
type List[T any] interface {
	// Len returns number of nodes contained in list.
	Len() int
	// Front returns first node of list or nil if list is empty.
	Front() *Node[T]
	// Back returns last node of list or nil if list is empty.
	Back() *Node[T]
	// Pop removes specified node from list and returns it.
	Pop(node *Node[T]) *Node[T]
	// PopFront removes first node from list and returns it.
	PopFront() *Node[T]
	// PopBack removes last node from list and returns it.
	PopBack() *Node[T]
	// InsertBefore inserts new node with value
	// before specified node and returns it.
	InsertBefore(value T, at *Node[T]) *Node[T]
	// InsertAfter inserts new node with value
	// after specified node and returns it.
	InsertAfter(value T, at *Node[T]) *Node[T]
	// PushFront inserts new node with value
	// at front of list and returns it.
	PushFront(value T) *Node[T]
	// PushBack inserts new node with value
	// at back of list and returns it.
	PushBack(value T) *Node[T]
}

// LinkedList implements doubly linked list.
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

// Len returns number of nodes contained in list, O(1).
func (l LinkedList[T]) Len() int { return l.len }

// Front returns first node of list or nil if list is empty, O(1).
func (l LinkedList[T]) Front() *Node[T] { return l.head }

// Back returns last node of list or nil if list is empty, O(1).
func (l LinkedList[T]) Back() *Node[T] { return l.tail }

// Pop removes specified node from list and returns it, O(1).
func (l *LinkedList[T]) Pop(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	// Check that node belongs to list.
	if node.list != l {
		return node
	}

	// If node is the first one, update the head,
	// else unlink it from previous node.
	if node.prev == nil {
		l.head = l.head.next
	} else {
		node.prev.next = node.next
	}

	// If node is the last one, update the tail,
	// else unlink it from next node.
	if node.next == nil {
		l.tail = l.tail.prev
	} else {
		node.next.prev = node.prev
	}

	// Avoid memory leaks.
	node.prev = nil
	node.next = nil
	node.list = nil

	l.len--
	return node
}

// PopFront removes first node from list and returns it, O(1).
func (l *LinkedList[T]) PopFront() *Node[T] { return l.Pop(l.head) }

// PopBack removes last node from list and returns it, O(1).
func (l *LinkedList[T]) PopBack() *Node[T] { return l.Pop(l.tail) }

// InsertBefore inserts new node with value before specified node, O(1).
// It returns the inserted node.
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

	// If at is the first node, update the head,
	// else link it to previous node.
	if at.prev == nil {
		l.head = node
	} else {
		at.prev.next = node
	}

	// Link at to new node.
	at.prev = node

	l.len++
	return node
}

// InsertAfter inserts new node with value after specified node, O(1).
// It returns the inserted node.
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

	// If at is the last node, update the tail,
	// else link it to next node.
	if at.next == nil {
		l.tail = node
	} else {
		at.next.prev = node
	}

	// Link at to new node.
	at.next = node

	l.len++
	return node
}

// PushFront inserts new node with value at front of list, O(1).
// It returns the inserted node.
func (l *LinkedList[T]) PushFront(value T) *Node[T] {
	// If list is empty, set head and tail.
	if l.head == nil {
		node := &Node[T]{Value: value, list: l}
		l.head = node
		l.tail = node
		l.len++
		return node
	}

	return l.InsertBefore(value, l.head)
}

// PushBack inserts new node with value at back of list, O(1).
// It returns the inserted node.
func (l *LinkedList[T]) PushBack(value T) *Node[T] {
	// If list is empty, insert node at front of list.
	if l.tail == nil {
		return l.PushFront(value)
	}

	return l.InsertAfter(value, l.tail)
}

// CircularLinkedList implements circular doubly linked list.
type CircularLinkedList[T any] struct {
	head *Node[T]
	len  int
}

// Len returns number of nodes contained in list, O(1).
func (l CircularLinkedList[T]) Len() int { return l.len }

// Front returns first node of list, O(1).
func (l CircularLinkedList[T]) Front() *Node[T] { return l.head }

// Back returns last node of list, O(1).
func (l CircularLinkedList[T]) Back() *Node[T] {
	if l.head == nil {
		return nil
	}
	return l.head.prev
}

// Pop removes specified node from list and returns it, O(1).
func (l *CircularLinkedList[T]) Pop(node *Node[T]) *Node[T] {
	if node == nil {
		return nil
	}

	// Check that node belongs to list.
	if node.list != l {
		return node
	}

	// If node is the first one, update the head,
	if node == l.head {
		l.head = node.Next() // if head links to itself, Next() returns nil
	}

	// Unlink node from previous and next nodes.
	node.prev.next = node.next
	node.next.prev = node.prev

	// Avoid memory leaks.
	node.prev = nil
	node.next = nil
	node.list = nil

	l.len--
	return node
}

// PopFront removes first node from list and returns it, O(1).
func (l *CircularLinkedList[T]) PopFront() *Node[T] { return l.Pop(l.head) }

// PopBack removes last node from list and returns it, O(1).
func (l *CircularLinkedList[T]) PopBack() *Node[T] {
	if l.head == nil {
		return nil
	}
	return l.Pop(l.head.prev)
}

// InsertBefore inserts new node with value before specified node, O(1).
// It returns the inserted node.
func (l *CircularLinkedList[T]) InsertBefore(value T, at *Node[T]) *Node[T] {
	if at == nil {
		return nil
	}

	return l.InsertAfter(value, at.prev)
}

// InsertAfter inserts new node with value after specified node, O(1).
// It returns the inserted node.
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

	// Link at to new node.
	at.next.prev = node
	at.next = node

	l.len++
	return node
}

// PushFront inserts new node with value at front of list, O(1).
// It returns the inserted node.
func (l *CircularLinkedList[T]) PushFront(value T) *Node[T] {
	// If list is empty, set head.
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

// PushBack inserts new node with value at back of list, O(1).
// It returns the inserted node.
func (l *CircularLinkedList[T]) PushBack(value T) *Node[T] {
	// If list is empty, insert node at front of list.
	if l.head == nil {
		return l.PushFront(value)
	}

	return l.InsertAfter(value, l.head.prev)
}
