package set

import (
	"reflect"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
	}{
		{"EmptySet", make(HashSet[int]), args{0}},
		{"Set", HashSet[int]{0: struct{}{}}, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.value)
			if !tt.s.Contains(tt.args.value) {
				t.Errorf("HashSet.Contains() = %v, want %v", false, true)
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
		{"EmptySet", make(HashSet[int]), args{0}},
		{"Set", HashSet[int]{1: struct{}{}}, args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.value)
			if tt.s.Contains(tt.args.value) {
				t.Errorf("HashSet.Contains() = %v, want %v", true, false)
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
		{"EmptySet", make(HashSet[int]), args{0}, false},
		{"Set", HashSet[int]{1: struct{}{}}, args{1}, true},
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
	smallerSet := HashSet[int]{0: struct{}{}}
	largerSet := HashSet[int]{0: struct{}{}, 1: struct{}{}}
	unionSet := HashSet[int]{0: struct{}{}, 1: struct{}{}}

	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want HashSet[int]
	}{
		{"LargerSet", largerSet, args{smallerSet}, unionSet},
		{"SmallerSet", smallerSet, args{largerSet}, unionSet},
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
	smallerSet := HashSet[int]{0: struct{}{}}
	largerSet := HashSet[int]{0: struct{}{}, 1: struct{}{}}
	intersectSet := HashSet[int]{0: struct{}{}}

	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want HashSet[int]
	}{
		{"LargerSet", largerSet, args{smallerSet}, intersectSet},
		{"SmallerSet", smallerSet, args{largerSet}, intersectSet},
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
	smallerSet := HashSet[int]{0: struct{}{}}
	largerSet := HashSet[int]{0: struct{}{}, 1: struct{}{}}

	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want HashSet[int]
	}{
		{"LargerSet", largerSet, args{smallerSet}, HashSet[int]{1: struct{}{}}},
		{"SmallerSet", smallerSet, args{largerSet}, make(HashSet[int])},
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
		{"UnequalSets", HashSet[int]{0: struct{}{}, 1: struct{}{}}, args{HashSet[int]{1: struct{}{}, 2: struct{}{}}}, HashSet[int]{0: struct{}{}, 2: struct{}{}}},
		{"EqualSets", HashSet[int]{0: struct{}{}, 1: struct{}{}}, args{HashSet[int]{0: struct{}{}, 1: struct{}{}}}, make(HashSet[int])},
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
	set := HashSet[int]{0: struct{}{}, 1: struct{}{}}

	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want bool
	}{
		{"UnequalLengths", set, args{HashSet[int]{1: struct{}{}}}, false},
		{"UnequalSets", set, args{HashSet[int]{1: struct{}{}, 2: struct{}{}}}, false},
		{"EqualSets", set, args{HashSet[int]{1: struct{}{}, 0: struct{}{}}}, true},
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
	set := HashSet[int]{0: struct{}{}, 1: struct{}{}}

	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want bool
	}{
		{"LargerSet", HashSet[int]{0: struct{}{}, 1: struct{}{}, 2: struct{}{}}, args{set}, false},
		{"NotSubset", HashSet[int]{2: struct{}{}}, args{set}, false},
		{"Subset", HashSet[int]{0: struct{}{}}, args{set}, true},
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
	set := HashSet[int]{0: struct{}{}, 1: struct{}{}}

	type args struct {
		set HashSet[int]
	}
	tests := []struct {
		name string
		s    HashSet[int]
		args args
		want bool
	}{
		{"SmallerSet", HashSet[int]{0: struct{}{}}, args{set}, false},
		{"NotSuperset", set, args{HashSet[int]{2: struct{}{}}}, false},
		{"Superset", set, args{HashSet[int]{0: struct{}{}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Superset(tt.args.set); got != tt.want {
				t.Errorf("HashSet.Superset() = %v, want %v", got, tt.want)
			}
		})
	}
}
