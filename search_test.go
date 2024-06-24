package alg

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr  []int
		item int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ExistingItem", args{[]int{1, 3, 5, 5, 7, 9}, 7}, 4},
		{"NonExistentItem", args{[]int{1, 3, 5, 5, 7, 9}, 2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.item, func(e int, t int) int { return e - t }); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
