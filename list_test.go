package alg

import (
	"reflect"
	"testing"
)

func TestNode_Prev(t *testing.T) {
	list := NewList[int]()
	head := list.PushFront(0)
	tail := list.PushBack(0)

	tests := []struct {
		name string
		n    *Node[int]
		want *Node[int]
	}{
		{"List", tail, head},
		{"NilList", &Node[int]{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Prev(); got != tt.want {
				t.Errorf("Node.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Next(t *testing.T) {
	list := NewList[int]()
	head := list.PushFront(0)
	tail := list.PushBack(0)

	tests := []struct {
		name string
		n    *Node[int]
		want *Node[int]
	}{
		{"List", head, tail},
		{"NilList", &Node[int]{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Next(); got != tt.want {
				t.Errorf("Node.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewList(t *testing.T) {
	tests := []struct {
		name string
		want *DoublyLinkedList[int]
	}{
		{"EmptyList", &DoublyLinkedList[int]{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Len(t *testing.T) {
	list := NewList[int]()
	list.PushBack(0)

	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		want int
	}{
		{"EmptyList", NewList[int](), 0},
		{"List", list, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("List.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Front(t *testing.T) {
	list := NewList[int]()
	node := list.PushBack(0)

	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		want *Node[int]
	}{
		{"EmptyList", NewList[int](), nil},
		{"List", list, node},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Front(); got != tt.want {
				t.Errorf("List.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Back(t *testing.T) {
	list := NewList[int]()
	node := list.PushBack(0)

	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		want *Node[int]
	}{
		{"EmptyList", NewList[int](), nil},
		{"List", list, node},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Back(); got != tt.want {
				t.Errorf("List.Back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Pop(t *testing.T) {
	list := NewList[int]()
	headNode := list.PushBack(0)
	innerNode := list.PushBack(0)
	tailNode := list.PushBack(0)
	emptyNode := new(Node[int])

	type args struct {
		node *Node[int]
	}
	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		args args
		want *Node[int]
	}{
		{"Node", list, args{innerNode}, innerNode},
		{"HeadNode", list, args{headNode}, headNode},
		{"TailNode", list, args{tailNode}, tailNode},
		{"NilNode", list, args{nil}, nil},
		{"NilList", list, args{emptyNode}, emptyNode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Pop(tt.args.node); got != tt.want {
				t.Errorf("List.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_PopFront(t *testing.T) {
	list := NewList[int]()
	head := list.PushBack(0)

	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		want *Node[int]
	}{
		{"HeadNode", list, head},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PopFront(); got != tt.want {
				t.Errorf("List.PopFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_PopBack(t *testing.T) {
	list := NewList[int]()
	tail := list.PushBack(0)

	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		want *Node[int]
	}{
		{"TailNode", list, tail},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PopBack(); got != tt.want {
				t.Errorf("List.PopBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_InsertBefore(t *testing.T) {
	list := NewList[int]()
	headNode := list.PushBack(0)
	tailNode := list.PushBack(0)

	type args struct {
		value int
		at    *Node[int]
	}
	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		args args
		want *Node[int]
	}{
		{"HeadNode", list, args{0, headNode}, &Node[int]{Value: 0, prev: nil, next: headNode, list: list}},
		{"TailNode", list, args{0, tailNode}, &Node[int]{Value: 0, prev: headNode, next: tailNode, list: list}},
		{"NilNode", list, args{0, nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.InsertBefore(tt.args.value, tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.InsertBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_InsertAfter(t *testing.T) {
	list := NewList[int]()
	headNode := list.PushBack(0)
	tailNode := list.PushBack(0)

	type args struct {
		value int
		at    *Node[int]
	}
	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		args args
		want *Node[int]
	}{
		{"TailNode", list, args{0, tailNode}, &Node[int]{Value: 0, prev: tailNode, next: nil, list: list}},
		{"HeadNode", list, args{0, headNode}, &Node[int]{Value: 0, prev: headNode, next: tailNode, list: list}},
		{"NilNode", list, args{0, nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.InsertAfter(tt.args.value, tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.InsertAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_PushFront(t *testing.T) {
	emptyList := NewList[int]()
	list := NewList[int]()
	headNode := list.PushBack(0)

	type args struct {
		value int
	}
	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		args args
		want *Node[int]
	}{
		{"EmptyList", emptyList, args{0}, &Node[int]{Value: 0, prev: nil, next: nil, list: emptyList}},
		{"List", list, args{0}, &Node[int]{Value: 0, prev: nil, next: headNode, list: list}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PushFront(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.PushFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_PushBack(t *testing.T) {
	emptyList := NewList[int]()
	list := NewList[int]()
	headNode := list.PushFront(0)

	type args struct {
		value int
	}
	tests := []struct {
		name string
		l    *DoublyLinkedList[int]
		args args
		want *Node[int]
	}{
		{"EmptyList", emptyList, args{0}, &Node[int]{Value: 0, prev: nil, next: nil, list: emptyList}},
		{"List", list, args{0}, &Node[int]{Value: 0, prev: headNode, next: nil, list: list}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PushBack(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.PushBack() = %v, want %v", got, tt.want)
			}
		})
	}
}
