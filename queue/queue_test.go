package queue

import (
	"reflect"
	"testing"

	"github.com/qsoulior/misc/list"
)

func emptyQueue() Queue[int] { return NewListQueue[int]() }

func simpleQueue() Queue[int] {
	q := NewListQueue[int]()
	q.PushBack(1)
	return q
}

func Test_nodeValue(t *testing.T) {
	type args struct {
		node *list.Node[int]
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"Node", args{&list.Node[int]{Value: 1}}, 1, true},
		{"NilNode", args{nil}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := nodeValue(tt.args.node)
			if got != tt.want {
				t.Errorf("nodeValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("nodeValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
		want Queue[int]
	}{
		{"EmptyQueue", &listQueue[int]{new(list.CircularLinkedList[int])}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListQueue[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		q    Queue[int]
		want int
	}{
		{"EmptyQueue", emptyQueue(), 0},
		{"SimpleQueue", simpleQueue(), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Len(); got != tt.want {
				t.Errorf("Queue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Front(t *testing.T) {
	tests := []struct {
		name  string
		q     Queue[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", emptyQueue(), 0, false},
		{"SimpleQueue", simpleQueue(), 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.Front()
			if got != tt.want {
				t.Errorf("Queue.Front() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Queue.Front() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueue_Back(t *testing.T) {
	tests := []struct {
		name  string
		q     Queue[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", emptyQueue(), 0, false},
		{"SimpleQueue", simpleQueue(), 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.Back()
			if got != tt.want {
				t.Errorf("Queue.Back() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Queue.Back() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueue_PopFront(t *testing.T) {
	tests := []struct {
		name  string
		q     Queue[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", emptyQueue(), 0, false},
		{"SimpleQueue", simpleQueue(), 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.PopFront()
			if got != tt.want {
				t.Errorf("Queue.PopFront() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Queue.PopFront() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueue_PushBack(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		q    Queue[int]
		args args
		want int
	}{
		{"EmptyQueue", emptyQueue(), args{1}, 1},
		{"SimpleQueue", simpleQueue(), args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.q.PushBack(tt.args.value)
			if got != tt.want {
				t.Errorf("Queue.PushBack() got = %v, want %v", got, tt.want)
			}
		})
	}
}
