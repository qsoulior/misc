package list

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
