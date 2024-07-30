package queue

import (
	"reflect"
	"testing"
)

func simplePriorityItem() *PriorityItem[int] { return &PriorityItem[int]{value: 0, priority: 2} }

func emptyMinPrioritySlice() minPrioritySlice[int] { return make(minPrioritySlice[int], 0) }

func simpleMinPrioritySlice() minPrioritySlice[int] {
	return minPrioritySlice[int]{
		&PriorityItem[int]{value: 3, priority: 1},
		&PriorityItem[int]{value: 1, priority: 2},
		&PriorityItem[int]{value: 2, priority: 3},
	}
}

func emptyMaxPrioritySlice() maxPrioritySlice[int] { return maxPrioritySlice[int]{} }

func simpleMaxPrioritySlice() maxPrioritySlice[int] {
	return maxPrioritySlice[int]{
		minPrioritySlice[int]{
			&PriorityItem[int]{value: 2, priority: 3},
			&PriorityItem[int]{value: 1, priority: 2},
			&PriorityItem[int]{value: 3, priority: 1},
		},
	}
}

func emptyPriorityQueue() PriorityQueue[int] { return NewMaxPriorityQueue[int]() }

func simplePriorityQueue() PriorityQueue[int] {
	pq := NewMaxPriorityQueue[int]()
	pq.Push(1, 2)
	pq.Push(2, 3)
	pq.Push(3, 1)
	return pq
}

func TestNewMinPriorityQueue(t *testing.T) {
	tests := []struct {
		name string
		want PriorityQueue[int]
	}{
		{"EmptyQueue", &priorityQueue[int]{new(minPrioritySlice[int])}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMinPriorityQueue[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMinPriorityQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMaxPriorityQueue(t *testing.T) {
	tests := []struct {
		name string
		want PriorityQueue[int]
	}{
		{"EmptyQueue", &priorityQueue[int]{new(maxPrioritySlice[int])}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaxPriorityQueue[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMaxPriorityQueue() = %v, want %v", got, tt.want)
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
		name  string
		p     PriorityQueue[int]
		args  args
		want  int
		want1 int
	}{
		{"EmptyQueue", emptyPriorityQueue(), args{4, 0}, 4, 0},
		{"SimpleQueue", simplePriorityQueue(), args{4, 0}, 4, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.Push(tt.args.value, tt.args.priority)
			if got != tt.want {
				t.Errorf("PriorityQueue.PushBack() = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PriorityQueue.PushBack() = %v, want1 %v", got1, tt.want1)
			}
		})
	}
}

func Test_minPrioritySlice_Len(t *testing.T) {
	tests := []struct {
		name string
		h    minPrioritySlice[int]
		want int
	}{
		{"EmptySlice", emptyMinPrioritySlice(), 0},
		{"SimpleSlice", simpleMinPrioritySlice(), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("minPrioritySlice.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPrioritySlice_First(t *testing.T) {
	tests := []struct {
		name string
		h    minPrioritySlice[int]
		want *PriorityItem[int]
	}{
		{"SimpleSlice", simpleMinPrioritySlice(), &PriorityItem[int]{value: 3, priority: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.First(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minPrioritySlice.First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPrioritySlice_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    minPrioritySlice[int]
		args args
		want *PriorityItem[int]
	}{
		{"SimpleSlice", simpleMinPrioritySlice(), args{0, 2}, &PriorityItem[int]{value: 2, priority: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
			if got := tt.h.First(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minPrioritySlice.Swap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPrioritySlice_Push(t *testing.T) {
	type args struct {
		x any
	}
	tests := []struct {
		name string
		h    minPrioritySlice[int]
		args args
		want *PriorityItem[int]
	}{
		{"EmptySlice", emptyMinPrioritySlice(), args{simplePriorityItem()}, simplePriorityItem()},
		{"SimpleSlice", simpleMinPrioritySlice(), args{simplePriorityItem()}, &PriorityItem[int]{value: 3, priority: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
			if got := tt.h.First(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minPrioritySlice.Push() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPrioritySlice_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    minPrioritySlice[int]
		want any
	}{
		{"SimpleSlice", simpleMinPrioritySlice(), &PriorityItem[int]{value: 2, priority: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minPrioritySlice.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPrioritySlice_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    minPrioritySlice[int]
		args args
		want bool
	}{
		{"LessPriority", simpleMinPrioritySlice(), args{0, 1}, true},
		{"GreaterPriority", simpleMinPrioritySlice(), args{2, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("minPrioritySlice.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxPrioritySlice_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    maxPrioritySlice[int]
		args args
		want bool
	}{
		{"LessPriority", simpleMaxPrioritySlice(), args{2, 1}, false},
		{"GreaterPriority", simpleMaxPrioritySlice(), args{0, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("maxPrioritySlice.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
