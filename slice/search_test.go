package slice

import (
	"slices"
	"testing"
)

type SearchFunc func(s []int, target int, cmp func(e, t int) int) int

func testSearch(t *testing.T, fn SearchFunc) {
	type args struct {
		s      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"NilSlice", args{nil, 2}, -1},
		{"EmptySlice", args{[]int{}, 2}, -1},
		{"ExistingItem", args{[]int{1, 3, 5, 5, 7, 9}, 7}, 4},
		{"NonExistentItem", args{[]int{1, 3, 5, 5, 7, 9}, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn(tt.args.s, tt.args.target, cmp); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearSearch(t *testing.T) {
	testSearch(t, LinearSearch)
}

func TestBinarySearch(t *testing.T) {
	testSearch(t, BinarySearch)
}

func TestJumpSearch(t *testing.T) {
	testSearch(t, JumpSearch)
}

func benchmarkSearch(b *testing.B, fn SearchFunc) {
	const n = 1e6
	s := randomSlice(n, 1e4)
	slices.Sort(s)
	target := s[n/2]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(s, target, cmp)
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	benchmarkSearch(b, LinearSearch)
}

func BenchmarkBinarySearch(b *testing.B) {
	benchmarkSearch(b, BinarySearch)
}

func BenchmarkJumpSearch(b *testing.B) {
	benchmarkSearch(b, JumpSearch)
}
