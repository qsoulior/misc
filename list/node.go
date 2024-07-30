package list

// Node implements a linked list node.
type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	list  List[T]
	Value T
}

// Prev returns previous node of list, O(1).
func (n *Node[T]) Prev() *Node[T] {
	// If node does not belong to list or links to the last node,
	// return nil.
	if n.list == nil || n.prev == n.list.Back() {
		return nil
	}
	return n.prev
}

// Next returns next node of list, O(1).
func (n *Node[T]) Next() *Node[T] {
	// If node does not belong to list or links to the first node,
	// return nil.
	if n.list == nil || n.next == n.list.Front() {
		return nil
	}
	return n.next
}
