package queue

import (
	"reflect"
	"testing"
)

func simpleItem() *priorityItem[int] { return &priorityItem[int]{value: 0, priority: 2, index: 0} }

func emptySlice() prioritySlice[int] { return make(prioritySlice[int], 0) }

func simpleSlice() prioritySlice[int] {
	h := make(prioritySlice[int], 0)
	h.Push(&priorityItem[int]{value: 1, priority: 2, index: 0})
	h.Push(&priorityItem[int]{value: 2, priority: 3, index: 1})
	h.Push(&priorityItem[int]{value: 3, priority: 1, index: 2})
	return h
}

func emptyQueue() PriorityQueue[int] { return NewPriorityQueue[int]() }

func simpleQueue() PriorityQueue[int] {
	pq := NewPriorityQueue[int]()
	pq.Push(1, 2)
	pq.Push(2, 3)
	pq.Push(3, 1)
	return pq
}

func TestNewPriorityQueue(t *testing.T) {
	tests := []struct {
		name string
		want PriorityQueue[int]
	}{
		{"EmptyQueue", &priorityQueue[int]{make(prioritySlice[int], 0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPriorityQueue[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPriorityQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		p    PriorityQueue[int]
		want int
	}{
		{"EmptyQueue", emptyQueue(), 0},
		{"SimpleQueue", simpleQueue(), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("PriorityQueue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Front(t *testing.T) {
	tests := []struct {
		name  string
		p     PriorityQueue[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", emptyQueue(), 0, false},
		{"SimpleQueue", simpleQueue(), 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.Front()
			if got != tt.want {
				t.Errorf("PriorityQueue.Front() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PriorityQueue.Front() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPriorityQueue_PopFront(t *testing.T) {
	tests := []struct {
		name  string
		p     PriorityQueue[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", emptyQueue(), 0, false},
		{"SimpleQueue", simpleQueue(), 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.PopFront()
			if got != tt.want {
				t.Errorf("PriorityQueue.PopFront() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PriorityQueue.PopFront() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	type args struct {
		value    int
		priority int
	}
	tests := []struct {
		name string
		p    PriorityQueue[int]
		args args
		want int
	}{
		{"EmptyQueue", emptyQueue(), args{4, 0}, 4},
		{"SimpleQueue", simpleQueue(), args{4, 0}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Push(tt.args.value, tt.args.priority); got != tt.want {
				t.Errorf("PriorityQueue.PushBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prioritySlice_Len(t *testing.T) {
	tests := []struct {
		name string
		h    prioritySlice[int]
		want int
	}{
		{"EmptySlice", emptySlice(), 0},
		{"SimpleSlice", simpleSlice(), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("prioritySlice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prioritySlice_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    prioritySlice[int]
		args args
		want bool
	}{
		{"LowerPriority", simpleSlice(), args{0, 1}, false},
		{"HigherPriority", simpleSlice(), args{1, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("prioritySlice.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prioritySlice_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name  string
		h     prioritySlice[int]
		args  args
		want  int
		want1 int
	}{
		{"SimpleSlice", simpleSlice(), args{0, 2}, 3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
			got, got1 := tt.h[0].value, tt.h[0].index
			if got != tt.want {
				t.Errorf("prioritySlice.Swap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("prioritySlice.Swap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_prioritySlice_Push(t *testing.T) {
	type args struct {
		x any
	}
	tests := []struct {
		name string
		h    prioritySlice[int]
		args args
		want int
	}{
		{"EmptySlice", emptySlice(), args{simpleItem()}, 0},
		{"SimpleSlice", simpleSlice(), args{simpleItem()}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
			if got := tt.h[0].value; got != tt.want {
				t.Errorf("prioritySlice.Push() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prioritySlice_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    prioritySlice[int]
		want any
	}{
		{"SimpleSlice", simpleSlice(), &priorityItem[int]{value: 3, priority: 1, index: -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prioritySlice.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
