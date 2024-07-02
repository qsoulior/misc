package queue

import (
	"reflect"
	"testing"

	"github.com/qsoulior/alg/list"
)

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
		{"EmptyQueue", &queue[int]{new(list.CircularLinkedList[int])}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Len(t *testing.T) {
	queue := NewQueue[int]()
	queue.PushBack(0)

	tests := []struct {
		name string
		q    Queue[int]
		want int
	}{
		{"EmptyQueue", NewQueue[int](), 0},
		{"Queue", queue, 1},
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
	queue := NewQueue[int]()
	queue.PushBack(1)

	tests := []struct {
		name  string
		q     Queue[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", NewQueue[int](), 0, false},
		{"Queue", queue, 1, true},
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
	queue := NewQueue[int]()
	queue.PushBack(1)

	tests := []struct {
		name  string
		q     Queue[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", NewQueue[int](), 0, false},
		{"Queue", queue, 1, true},
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
	queue := NewQueue[int]()
	queue.PushBack(1)

	tests := []struct {
		name  string
		q     Queue[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", NewQueue[int](), 0, false},
		{"Queue", queue, 1, true},
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
	queue := NewQueue[int]()
	queue.data.PushBack(1)

	type args struct {
		value int
	}
	tests := []struct {
		name string
		q    Queue[int]
		args args
		want int
	}{
		{"EmptyQueue", NewQueue[int](), args{0}, 0},
		{"Queue", queue, args{0}, 0},
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

func TestNewDeque(t *testing.T) {
	tests := []struct {
		name string
		want Deque[int]
	}{
		{"EmptyQueue", &deque[int]{NewQueue[int]()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeque[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeque() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_PopBack(t *testing.T) {
	queue := NewDeque[int]()
	queue.PushBack(1)

	tests := []struct {
		name  string
		d     Deque[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", NewDeque[int](), 0, false},
		{"Queue", queue, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.d.PopBack()
			if got != tt.want {
				t.Errorf("Deque.PopBack() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Deque.PopBack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeque_PushFront(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		d    Deque[int]
		args args
		want int
	}{
		{"EmptyDeque", NewDeque[int](), args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.PushFront(tt.args.value)
			if got != tt.want {
				t.Errorf("Deque.PushFront() got = %v, want %v", got, tt.want)
			}
		})
	}
}
