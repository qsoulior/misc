package slice

import (
	"math/rand"
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
		{"NonExistentItem", args{[]int{1, 3, 5, 5, 7, 9}, 10}, -1},
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

func TestInterpolationSearch(t *testing.T) {
	testSearch(t, InterpolationSearch)
}

func TestExponentialSearch(t *testing.T) {
	testSearch(t, ExponentialSearch)
}

func benchmarkSearch(b *testing.B, fn SearchFunc) {
	const n = 1e6
	s := randomSlice(n, 1e4)
	slices.Sort(s)
	target := s[4*n/5]
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

func BenchmarkInterpolationSearch(b *testing.B) {
	benchmarkSearch(b, InterpolationSearch)
}

func BenchmarkExponentialSearch(b *testing.B) {
	benchmarkSearch(b, ExponentialSearch)
}

func fuzzSearch(f *testing.F, fn SearchFunc) {
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
		slices.Sort(s)
		target := s[rand.Intn(1000)]
		i := fn(s, target, cmp)
		if i == -1 {
			t.Error("target element not found")
		}
		if s[i] != target {
			t.Error("found element is not equal to target")
		}
	})
}

func FuzzLinearSearch(f *testing.F) {
	fuzzSearch(f, LinearSearch)
}

func FuzzBinarySearch(f *testing.F) {
	fuzzSearch(f, BinarySearch)
}

func FuzzJumpSearch(f *testing.F) {
	fuzzSearch(f, JumpSearch)
}

func FuzzInterpolationSearch(f *testing.F) {
	fuzzSearch(f, InterpolationSearch)
}

func FuzzExponentialSearch(f *testing.F) {
	fuzzSearch(f, ExponentialSearch)
}
