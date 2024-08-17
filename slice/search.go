package slice

import "math"

// LinearSearch searches for target in slice with complexity O(n).
// It returns index of target or -1 if target is not found.
// cmp should return 0 if slice element matches target.
func LinearSearch[S ~[]E, E, T any](s S, target T, cmp func(E, T) int) int {
	for i, v := range s {
		if cmp(v, target) == 0 {
			return i
		}
	}
	return -1
}

// BinarySearch searches for target in sorted slice with complexity O(log(n)).
// It returns index of target or -1 if target is not found.
// cmp should return 0 if slice element matches target, a negative number
// if slice element precedes target, or a positive number if slice element follows target.
func BinarySearch[S ~[]E, E, T any](s S, target T, cmp func(E, T) int) int {
	left, right := 0, len(s)-1 // boundaries of slice part, in which search is performed

	for left <= right {
		mid := (left + right) / 2 // index of element in the middle of slice part
		c := cmp(s[mid], target)  // result of comparing middle element with target

		if c == 0 {
			// If middle element matches target, returns its index.
			return mid
		}

		if c < 0 {
			// If middle element precedes target,
			// then shift left boundary to the right of middle element.
			left = mid + 1
		} else {
			// If middle element follows target,
			// then shift right boundary to the left of middle element.
			right = mid - 1
		}
	}

	return -1
}

// JumpSearch searches for target in sorted slice with complexity O(sqrt(n)).
// It returns index of target or -1 if target is not found.
// cmp should return 0 if slice element matches target, a negative number
// if slice element precedes target, or a positive number if slice element follows target.
func JumpSearch[S ~[]E, E, T any](s S, target T, cmp func(E, T) int) int {
	n := len(s)
	step := int(math.Sqrt(float64(n)))

	// Search for range that possibly contains target.
	i := 0
	j := step
	for j < n && cmp(s[j], target) <= 0 {
		i = j
		j += step
	}

	// Fix right boundary if it is out of slice range.
	if j >= n {
		j = n - 1
	}

	// Search for target in range [i, j].
	for ; i <= j; i++ {
		if cmp(s[i], target) == 0 {
			return i
		}
	}

	return -1
}
