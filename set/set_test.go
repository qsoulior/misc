package set

import (
	"reflect"
	"testing"
)

func emptySet() HashSet[int] { return make(HashSet[int]) }

func simpleSetA() HashSet[int] { return HashSet[int]{1: struct{}{}} }

func simpleSetB() HashSet[int] { return HashSet[int]{0: struct{}{}, 2: struct{}{}} }

func simpleSetC() HashSet[int] { return HashSet[int]{0: struct{}{}, 1: struct{}{}} }

func TestHashSet_Len(t *testing.T) {
	tests := []struct {
		name string
		s    HashSet[int]
		want int
	}{
		{"EmptySet", emptySet(), 0},
		{"SimpleSet", simpleSetB(), 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("HashSet.Len() = %v, want %v", false, true)
			}
		})
	}
}

func TestHashSet_Add(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
	}{
		{"EmptySet", emptySet(), args{0}},
		{"SimpleSet", simpleSetB(), args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.value)
			if !tt.s.Contains(tt.args.value) {
				t.Errorf("set does not contain %v after HashSet.Add()", tt.args.value)
			}
		})
	}
}

func TestHashSet_Remove(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
	}{
		{"EmptySet", emptySet(), args{0}},
		{"SimpleSet", simpleSetB(), args{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.value)
			if tt.s.Contains(tt.args.value) {
				t.Errorf("set contains %v after HashSet.Remove()", tt.args.value)
			}
		})
	}
}

func TestHashSet_Contains(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want bool
	}{
		{"EmptySet", emptySet(), args{0}, false},
		{"Set", simpleSetB(), args{2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.value); got != tt.want {
				t.Errorf("HashSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Union(t *testing.T) {
	unionSet := HashSet[int]{0: struct{}{}, 1: struct{}{}, 2: struct{}{}}
	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want HashSet[int]
	}{
		{"LargerSet", simpleSetB(), args{simpleSetA()}, unionSet},
		{"SmallerSet", simpleSetA(), args{simpleSetB()}, unionSet},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Union(tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashSet.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Intersection(t *testing.T) {
	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want HashSet[int]
	}{
		{"LargerSet", simpleSetC(), args{simpleSetA()}, simpleSetA()},
		{"SmallerSet", simpleSetA(), args{simpleSetC()}, simpleSetA()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashSet.Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Difference(t *testing.T) {
	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want HashSet[int]
	}{
		{"LargerSet", simpleSetB(), args{simpleSetC()}, HashSet[int]{2: struct{}{}}},
		{"SmallerSet", simpleSetC(), args{simpleSetB()}, HashSet[int]{1: struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Difference(tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashSet.Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_SymmetricDifference(t *testing.T) {
	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want HashSet[int]
	}{
		{"UnequalSets", simpleSetB(), args{simpleSetC()}, HashSet[int]{1: struct{}{}, 2: struct{}{}}},
		{"EqualSets", simpleSetB(), args{simpleSetB()}, emptySet()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SymmetricDifference(tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HashSet.SymmetricDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Equal(t *testing.T) {
	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want bool
	}{
		{"UnequalLengths", simpleSetB(), args{HashSet[int]{2: struct{}{}}}, false},
		{"UnequalSets", simpleSetB(), args{HashSet[int]{1: struct{}{}, 2: struct{}{}}}, false},
		{"EqualSets", simpleSetB(), args{simpleSetB()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.set); got != tt.want {
				t.Errorf("HashSet.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Subset(t *testing.T) {
	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want bool
	}{
		{"LargerSet", simpleSetB(), args{simpleSetA()}, false},
		{"NotSubset", simpleSetA(), args{simpleSetB()}, false},
		{"Subset", HashSet[int]{0: struct{}{}}, args{simpleSetB()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Subset(tt.args.set); got != tt.want {
				t.Errorf("HashSet.Subset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSet_Superset(t *testing.T) {
	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want bool
	}{
		{"SmallerSet", simpleSetA(), args{simpleSetB()}, false},
		{"NotSuperset", simpleSetB(), args{simpleSetA()}, false},
		{"Superset", simpleSetB(), args{HashSet[int]{0: struct{}{}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Superset(tt.args.set); got != tt.want {
				t.Errorf("HashSet.Superset() = %v, want %v", got, tt.want)
			}
		})
	}
}
