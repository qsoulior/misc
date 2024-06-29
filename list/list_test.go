package list

import (
	"reflect"
	"testing"
)

func TestLinkedList_Len(t *testing.T) {
	list := new(LinkedList[int])
	list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want int
	}{
		{"EmptyList", new(LinkedList[int]), 0},
		{"List", list, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("LinkedList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Front(t *testing.T) {
	list := new(LinkedList[int])
	node := list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want *Node[int]
	}{
		{"EmptyList", new(LinkedList[int]), nil},
		{"List", list, node},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Front(); got != tt.want {
				t.Errorf("LinkedList.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Back(t *testing.T) {
	list := new(LinkedList[int])
	node := list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want *Node[int]
	}{
		{"EmptyList", new(LinkedList[int]), nil},
		{"List", list, node},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Back(); got != tt.want {
				t.Errorf("LinkedList.Back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	list := new(LinkedList[int])
	headNode := list.PushBack(0)
	innerNode := list.PushBack(0)
	tailNode := list.PushBack(0)
	emptyNode := new(Node[int])

	type args struct {
		node *Node[int]
	}
	tests := []struct {
		name string
		l    List[int]
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
				t.Errorf("LinkedList.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PopFront(t *testing.T) {
	list := new(LinkedList[int])
	head := list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want *Node[int]
	}{
		{"EmptyList", new(LinkedList[int]), nil},
		{"List", list, head},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PopFront(); got != tt.want {
				t.Errorf("LinkedList.PopFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PopBack(t *testing.T) {
	list := new(LinkedList[int])
	tail := list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want *Node[int]
	}{
		{"EmptyList", new(LinkedList[int]), nil},
		{"List", list, tail},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PopBack(); got != tt.want {
				t.Errorf("LinkedList.PopBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_InsertBefore(t *testing.T) {
	list := new(LinkedList[int])
	headNode := list.PushBack(0)
	tailNode := list.PushBack(0)

	type args struct {
		value int
		at    *Node[int]
	}
	tests := []struct {
		name string
		l    List[int]
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
				t.Errorf("LinkedList.InsertBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_InsertAfter(t *testing.T) {
	list := new(LinkedList[int])
	headNode := list.PushBack(0)
	tailNode := list.PushBack(0)

	type args struct {
		value int
		at    *Node[int]
	}
	tests := []struct {
		name string
		l    List[int]
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
				t.Errorf("LinkedList.InsertAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PushFront(t *testing.T) {
	emptyList := new(LinkedList[int])
	list := new(LinkedList[int])
	headNode := list.PushBack(0)

	type args struct {
		value int
	}
	tests := []struct {
		name string
		l    List[int]
		args args
		want *Node[int]
	}{
		{"EmptyList", emptyList, args{0}, &Node[int]{Value: 0, prev: nil, next: nil, list: emptyList}},
		{"List", list, args{0}, &Node[int]{Value: 0, prev: nil, next: headNode, list: list}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PushFront(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.PushFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PushBack(t *testing.T) {
	emptyList := new(LinkedList[int])
	list := new(LinkedList[int])
	headNode := list.PushFront(0)

	type args struct {
		value int
	}
	tests := []struct {
		name string
		l    List[int]
		args args
		want *Node[int]
	}{
		{"EmptyList", emptyList, args{0}, &Node[int]{Value: 0, prev: nil, next: nil, list: emptyList}},
		{"List", list, args{0}, &Node[int]{Value: 0, prev: headNode, next: nil, list: list}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PushBack(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.PushBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_Len(t *testing.T) {
	list := new(CircularLinkedList[int])
	list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want int
	}{
		{"EmptyList", new(CircularLinkedList[int]), 0},
		{"List", list, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("CircularLinkedList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_Front(t *testing.T) {
	list := new(CircularLinkedList[int])
	node := list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want *Node[int]
	}{
		{"EmptyList", new(CircularLinkedList[int]), nil},
		{"List", list, node},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Front(); got != tt.want {
				t.Errorf("CircularLinkedList.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_Back(t *testing.T) {
	list := new(CircularLinkedList[int])
	node := list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want *Node[int]
	}{
		{"EmptyList", new(CircularLinkedList[int]), nil},
		{"List", list, node},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Back(); got != tt.want {
				t.Errorf("CircularLinkedList.Back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_Pop(t *testing.T) {
	list := new(CircularLinkedList[int])
	headNode := list.PushBack(0)
	innerNode := list.PushBack(0)
	tailNode := list.PushBack(0)
	emptyNode := new(Node[int])

	type args struct {
		node *Node[int]
	}
	tests := []struct {
		name string
		l    List[int]
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
				t.Errorf("CircularLinkedList.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_PopFront(t *testing.T) {
	list := new(CircularLinkedList[int])
	head := list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want *Node[int]
	}{
		{"EmptyList", new(CircularLinkedList[int]), nil},
		{"List", list, head},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PopFront(); got != tt.want {
				t.Errorf("CircularLinkedList.PopFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_PopBack(t *testing.T) {
	list := new(CircularLinkedList[int])
	tail := list.PushBack(0)

	tests := []struct {
		name string
		l    List[int]
		want *Node[int]
	}{
		{"EmptyList", new(CircularLinkedList[int]), nil},
		{"List", list, tail},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PopBack(); got != tt.want {
				t.Errorf("CircularLinkedList.PopBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_InsertBefore(t *testing.T) {
	list := new(CircularLinkedList[int])
	headNode := list.PushBack(0)
	tailNode := list.PushBack(0)

	type args struct {
		value int
		at    *Node[int]
	}
	tests := []struct {
		name string
		l    List[int]
		args args
		want *Node[int]
	}{
		{"HeadNode", list, args{0, headNode}, &Node[int]{Value: 0, prev: tailNode, next: headNode, list: list}},
		{"TailNode", list, args{0, tailNode}, &Node[int]{Value: 0, prev: headNode, next: tailNode, list: list}},
		{"NilNode", list, args{0, nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.InsertBefore(tt.args.value, tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CircularLinkedList.InsertBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_InsertAfter(t *testing.T) {
	list := new(CircularLinkedList[int])
	headNode := list.PushBack(0)
	tailNode := list.PushBack(0)

	type args struct {
		value int
		at    *Node[int]
	}
	tests := []struct {
		name string
		l    List[int]
		args args
		want *Node[int]
	}{
		{"TailNode", list, args{0, tailNode}, &Node[int]{Value: 0, prev: tailNode, next: headNode, list: list}},
		{"HeadNode", list, args{0, headNode}, &Node[int]{Value: 0, prev: headNode, next: tailNode, list: list}},
		{"NilNode", list, args{0, nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.InsertAfter(tt.args.value, tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CircularLinkedList.InsertAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_PushFront(t *testing.T) {
	emptyList := new(CircularLinkedList[int])
	list := new(CircularLinkedList[int])
	tailNode := list.PushBack(0)

	headNode := &Node[int]{Value: 0, list: emptyList}
	headNode.prev = headNode
	headNode.next = headNode

	type args struct {
		value int
	}
	tests := []struct {
		name string
		l    List[int]
		args args
		want *Node[int]
	}{
		{"EmptyList", emptyList, args{0}, headNode},
		{"List", list, args{0}, &Node[int]{Value: 0, prev: tailNode, next: tailNode, list: list}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PushFront(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CircularLinkedList.PushFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircularLinkedList_PushBack(t *testing.T) {
	emptyList := new(CircularLinkedList[int])
	list := new(CircularLinkedList[int])
	headNode := list.PushFront(0)

	tailNode := &Node[int]{Value: 0, list: emptyList}
	tailNode.prev = tailNode
	tailNode.next = tailNode

	type args struct {
		value int
	}
	tests := []struct {
		name string
		l    List[int]
		args args
		want *Node[int]
	}{
		{"EmptyList", emptyList, args{0}, tailNode},
		{"List", list, args{0}, &Node[int]{Value: 0, prev: headNode, next: headNode, list: list}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.PushBack(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CircularLinkedList.PushBack() = %v, want %v", got, tt.want)
			}
		})
	}
}
