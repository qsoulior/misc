package queue

import (
	"reflect"
	"testing"
)

func simplePriorityItem() *priorityItem[int] {
	return &priorityItem[int]{value: 0, priority: 2}
}

func emptyPrioritySlice() prioritySlice[int] { return make(prioritySlice[int], 0) }

func simplePrioritySlice() prioritySlice[int] {
	h := make(prioritySlice[int], 0)
	h.Push(&priorityItem[int]{value: 1, priority: 2})
	h.Push(&priorityItem[int]{value: 2, priority: 3})
	h.Push(&priorityItem[int]{value: 3, priority: 1})
	return h
}

func emptyPriorityQueue() PriorityQueue[int] { return NewPriorityQueue[int]() }

func simplePriorityQueue() PriorityQueue[int] {
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
		{"EmptyQueue", emptyPriorityQueue(), 0},
		{"SimpleQueue", simplePriorityQueue(), 3},
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
		want1 int
		want2 bool
	}{
		{"EmptyQueue", emptyPriorityQueue(), 0, 0, false},
		{"SimpleQueue", simplePriorityQueue(), 2, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.p.Front()
			if got != tt.want {
				t.Errorf("PriorityQueue.Front() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PriorityQueue.Front() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("PriorityQueue.Front() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestPriorityQueue_PopFront(t *testing.T) {
	tests := []struct {
		name  string
		p     PriorityQueue[int]
		want  int
		want1 int
		want2 bool
	}{
		{"EmptyQueue", emptyPriorityQueue(), 0, 0, false},
		{"SimpleQueue", simplePriorityQueue(), 2, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.p.PopFront()
			if got != tt.want {
				t.Errorf("PriorityQueue.PopFront() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PriorityQueue.PopFront() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("PriorityQueue.PopFront() got2 = %v, want %v", got2, tt.want2)
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
		{"EmptyQueue", emptyPriorityQueue(), args{4, 0}, 4},
		{"SimpleQueue", simplePriorityQueue(), args{4, 0}, 4},
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
		{"EmptySlice", emptyPrioritySlice(), 0},
		{"SimpleSlice", simplePrioritySlice(), 3},
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
		{"LowerPriority", simplePrioritySlice(), args{0, 1}, false},
		{"HigherPriority", simplePrioritySlice(), args{1, 2}, true},
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
		name string
		h    prioritySlice[int]
		args args
		want int
	}{
		{"SimpleSlice", simplePrioritySlice(), args{0, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
			got := tt.h[0].value
			if got != tt.want {
				t.Errorf("prioritySlice.Swap() got = %v, want %v", got, tt.want)
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
		{"EmptySlice", emptyPrioritySlice(), args{simplePriorityItem()}, 0},
		{"SimpleSlice", simplePrioritySlice(), args{simplePriorityItem()}, 1},
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
		{"SimpleSlice", simplePrioritySlice(), &priorityItem[int]{value: 3, priority: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prioritySlice.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
