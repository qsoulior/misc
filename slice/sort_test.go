package slice

import (
	"math/rand"
	"reflect"
	"slices"
	"testing"
)

type SortFunc func(s []int, cmp func(a, b int) int)

func cmp(a, b int) int { return a - b }

func testSort(t *testing.T, fn SortFunc) {
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
		{"UnsortedSlice", args{[]int{6, 2, 3, 9, 1, 4, 1, 10, 5, 3, 7, 8, 1}}, []int{1, 1, 1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.arr
			fn(got, cmp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	testSort(t, BubbleSort)
}

func TestCocktailSort(t *testing.T) {
	testSort(t, CocktailSort)
}

func TestCombSort(t *testing.T) {
	testSort(t, CombSort)
}

func TestSelectionSort(t *testing.T) {
	testSort(t, SelectionSort)
}

func TestInsertionSort(t *testing.T) {
	testSort(t, InsertionSort)
}

func TestQuickSort(t *testing.T) {
	testSort(t, QuickSort)
}

func TestMergeSort(t *testing.T) {
	testSort(t, MergeSort)
}

func TestHeapSort(t *testing.T) {
	testSort(t, HeapSort)
}

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

func BenchmarkInsertionSort(b *testing.B) {
	benchmarkSort(b, InsertionSort)
}

func BenchmarkQuickSort(b *testing.B) {
	benchmarkSort(b, QuickSort)
}

func BenchmarkMergeSort(b *testing.B) {
	benchmarkSort(b, MergeSort)
}

func BenchmarkHeapSort(b *testing.B) {
	benchmarkSort(b, HeapSort)
}

func fuzzSort(f *testing.F, fn SortFunc) {
	for range 100 {
		s := randomSlice(1000, 1000)
		b := make([]byte, len(s))
		for i := range s {
			b[i] = byte(s[i])
		}
		f.Add(b)
	}
	f.Fuzz(func(t *testing.T, b []byte) {
		s := make([]int, len(b))
		for i := range s {
			s[i] = (int(b[i]) * rand.Intn(256)) % 1000
		}
		fn(s, cmp)
		if !slices.IsSortedFunc(s, cmp) {
			t.Error("s is not sorted")
		}
	})
}

func FuzzBubbleSort(f *testing.F) {
	fuzzSort(f, BubbleSort)
}

func FuzzCocktailSort(f *testing.F) {
	fuzzSort(f, CocktailSort)
}

func FuzzCombSort(f *testing.F) {
	fuzzSort(f, CombSort)
}

func FuzzSelectionSort(f *testing.F) {
	fuzzSort(f, SelectionSort)
}

func FuzzInsertionSort(f *testing.F) {
	fuzzSort(f, InsertionSort)
}

func FuzzQuickSort(f *testing.F) {
	fuzzSort(f, QuickSort)
}

func FuzzMergeSort(f *testing.F) {
	fuzzSort(f, MergeSort)
}

func FuzzHeapSort(f *testing.F) {
	fuzzSort(f, HeapSort)
}
