package slice

import (
	"math/rand"
	"reflect"
	"slices"
	"testing"
)

type SortFunc func(x []int, cmp func(a, b int) int)

func cmp(a, b int) int { return a - b }

// Tests
func TestBubbleSort(t *testing.T) {
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
			BubbleSort(got, cmp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCocktailSort(t *testing.T) {
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
			CocktailSort(got, cmp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CocktailSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombSort(t *testing.T) {
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
			CombSort(got, cmp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			SelectionSort(got, cmp)
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
			QuickSort(got, cmp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmarks
func randomSlice(n, max int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(max)
	}

	return s
}

func benchmarkSort(b *testing.B, fn SortFunc) {
	ns := make([][]int, b.N)
	for i := 0; i < b.N; i++ {
		ns[i] = randomSlice(10000, 1000)
	}
	b.ResetTimer()
	for _, s := range ns {
		fn(s, cmp)
	}
}

func BenchmarkSort(b *testing.B) {
	benchmarkSort(b, slices.SortFunc)
}

func BenchmarkBubbleSort(b *testing.B) {
	benchmarkSort(b, BubbleSort)
}

func BenchmarkCocktailSort(b *testing.B) {
	benchmarkSort(b, CocktailSort)
}

func BenchmarkCombSort(b *testing.B) {
	benchmarkSort(b, CombSort)
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmarkSort(b, SelectionSort)
}

func BenchmarkQuickSort(b *testing.B) {
	benchmarkSort(b, QuickSort)
}
