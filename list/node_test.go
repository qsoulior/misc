package list

import "testing"

func TestNode_Prev(t *testing.T) {
	list := new(LinkedList[int])
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
	list := new(LinkedList[int])
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
