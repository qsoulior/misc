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
	left := 0
	right := step
	for right < n && cmp(s[right], target) <= 0 {
		left = right
		right += step
	}

	// Fix right boundary if it is out of slice range.
	if right >= n {
		right = n - 1
	}

	// Search for target in range [left, right].
	for ; left <= right; left++ {
		if cmp(s[left], target) == 0 {
			return left
		}
	}

	return -1
}

// InterpolationSearch searches for target in sorted slice with complexity O(log(log(n))).
// It returns index of target or -1 if target is not found.
// cmp should return 0 if slice element matches target, a negative number
// if slice element precedes target, or a positive number if slice element follows target.
// The value returned by cmp should reflect proximity of element to target.
func InterpolationSearch[S ~[]E, E any](s S, target E, cmp func(E, E) int) int {
	left, right := 0, len(s)-1 // boundaries of slice part, in which search is performed
	for left <= right {
		// If divisor is zero, then use middle element as in binary search.
		idx := (left + right) / 2
		if div := cmp(s[right], s[left]); div != 0 {
			idx = left + cmp(target, s[left])*(right-left)/div // probe index
		}

		c := cmp(s[idx], target) // result of comparing probe with target

		if c == 0 {
			// If probe matches target, returns its index.
			return idx
		}

		if c < 0 {
			// If probe precedes target,
			// then shift left boundary to the right of probe.
			left = idx + 1
		} else {
			// If probe follows target,
			// then shift right boundary to the left of probe.
			right = idx - 1
		}
	}

	return -1
}

// ExponentialSearch searches for target in sorted slice with complexity O(log(i)),
// where i is the index of target element.
// It returns index of target or -1 if target is not found.
// cmp should return 0 if slice element matches target, a negative number
// if slice element precedes target, or a positive number if slice element follows target.
func ExponentialSearch[S ~[]E, E, T any](s S, target T, cmp func(E, T) int) int {
	n := len(s)

	// Search for range that possibly contains target.
	right := 1
	for right < n && cmp(s[right], target) <= 0 {
		right *= 2
	}

	left := right / 2

	// Fix right boundary if it is out of slice range.
	if right > n {
		right = n
	}

	// Search for target in range [left, right) using binary search algorithm.
	if i := BinarySearch(s[left:right], target, cmp); i != -1 {
		return left + i
	}

	return -1
}
