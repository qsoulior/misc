package queue

import (
	"reflect"
	"testing"
)

func emptyDeque() Deque[int] { return NewDeque[int]() }

func simpleDeque() Deque[int] {
	d := NewDeque[int]()
	d.PushBack(1)
	return d
}

func TestNewDeque(t *testing.T) {
	tests := []struct {
		name string
		want Deque[int]
	}{
		{"EmptyQueue", emptyDeque()},
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
	tests := []struct {
		name  string
		d     Deque[int]
		want  int
		want1 bool
	}{
		{"EmptyQueue", emptyDeque(), 0, false},
		{"SimpleQueue", simpleDeque(), 1, true},
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
		{"EmptyQueue", emptyDeque(), args{0}, 0},
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
