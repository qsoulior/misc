package slice

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"NilSlice", args{nil}, nil},
		{"EmptySlice", args{[]int{}}, []int{}},
		{"UnsortedSlice", args{[]int{6, 2, 3, 9, 1, 4, 1}}, []int{1, 1, 2, 3, 4, 6, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.arr
			SelectionSort(got, func(a int, b int) int { return a - b })
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"NilSlice", args{nil}, nil},
		{"EmptySlice", args{[]int{}}, []int{}},
		{"UnsortedSlice", args{[]int{6, 2, 3, 9, 1, 4, 1}}, []int{1, 1, 2, 3, 4, 6, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.arr
			QuickSort(got, func(a int, b int) int { return a - b })
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
