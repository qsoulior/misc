package slice

// BinarySearch searches for target in sorted slice with complexity O(log(n)).
// It returns index of target or -1 if target is not found.
// cmp should return 0 if slice element matches target, a negative number
// if slice element precedes target, or a positive number if slice element follows target.
func BinarySearch[S ~[]E, E, T any](arr S, target T, cmp func(E, T) int) int {
	left, right := 0, len(arr)-1 // boundaries of slice part, in which search is performed

	for left <= right {
		mid := (left + right) / 2  // index of element in the middle of slice part
		c := cmp(arr[mid], target) // result of comparing middle element with target

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
