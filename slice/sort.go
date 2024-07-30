package slice

import "math/rand"

// SelectionSort sorts slice s in order as determined by cmp function.
// It uses selection sort algorithm with complexity O(n^2).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func SelectionSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	n := len(s)
	for i := 0; i < n; i++ {
		// Search for index of the smallest element in unsorted part of slice.
		min := i
		for j := i + 1; j < n; j++ {
			if cmp(s[j], s[min]) < 0 {
				min = j
			}
		}

		// Swap current element and the smallest element of unsorted part.
		s[i], s[min] = s[min], s[i]
	}
}

// QuickSort sorts slice s in order as determined by cmp function.
// It uses recursive quick sort algorithm with complexity O(n*log(n)).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func QuickSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	n := len(s)
	if n < 2 {
		// If slice contains 0 or 1 element, it is sorted.
		return
	}

	left, right := 0, n-1 // slice boundaries (indexes of first and last elements)
	pivot := rand.Intn(n) // index of pivot element

	// Swap pivot and last elements.
	s[pivot], s[right] = s[right], s[pivot]

	for i := 0; i < right; i++ {
		// If element precedes pivot, then move it
		// to the beginning after other moved elements.
		if cmp(s[i], s[right]) < 0 {
			s[i], s[left] = s[left], s[i]
			left++
		}
	}

	// Swap pivot element and element following last moved element.
	s[left], s[right] = s[right], s[left]

	QuickSort(s[:left], cmp)
	QuickSort(s[left+1:], cmp)
}
