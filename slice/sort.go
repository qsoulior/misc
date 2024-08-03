// Package slice implements various operations with slices.
// It provides sort and search algorithms.
package slice

import "math/rand"

// BubbleSort sorts slice s in order as determined by cmp function.
// It uses bubble sort algorithm with complexity O(n^2).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func BubbleSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	swapped := true
	for i := len(s) - 1; i > 0 && swapped; i-- {
		swapped = false
		// Move the largest element to the end of unsorted slice part.
		for j := 0; j < i; j++ {
			if cmp(s[j], s[j+1]) > 0 {
				s[j], s[j+1] = s[j+1], s[j]
				swapped = true
			}
		}
	}
}

// CocktailSort sorts slice s in order as determined by cmp function.
// It uses cocktail sort (bidirectional bubble sort) algorithm with complexity O(n^2).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func CocktailSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	left, right := 0, len(s)-1 // boundaries of unsorted slice part
	swapped := true            // if no elements have been swapped, then slice is sorted
	for left < right && swapped {
		// Move the largest element to the end of unsorted slice part.
		swapped = false
		for i := left; i < right; i++ {
			if cmp(s[i], s[i+1]) > 0 {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
		right-- // decrease right boundary of unsorted slice part

		// Move the smallest element to the beginning of unsorted slice part.
		swapped = false
		for i := right; i > left; i-- {
			if cmp(s[i], s[i-1]) < 0 {
				s[i], s[i-1] = s[i-1], s[i]
				swapped = true
			}
		}
		left++ // increase left boundary of unsorted slice part
	}
}

// CombSort sorts slice s in order as determined by cmp function.
// It uses comb sort (modification of bubble sort) algorithm with complexity O(n^2),
// where p is number of increments.
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func CombSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	const factor = 1.3 // optimal shrink factor suggested by Lacey and Box
	n := len(s)
	step := n - 1 // optimal initial step
	for step > 1 {
		for i := 0; i < n-step; i++ {
			if cmp(s[i], s[i+step]) > 0 {
				s[i], s[i+step] = s[i+step], s[i]
			}
		}

		// Update the step.
		step = int(float32(step) / factor)
		if step == 9 || step == 10 {
			step = 11 // rule of 11
		}
	}

	// When the step is 1, comb sort is equivalent to bubble sort.
	BubbleSort(s, cmp)
}

// SelectionSort sorts slice s in order as determined by cmp function.
// It uses selection sort algorithm with complexity O(n^2).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func SelectionSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		// Search for index of the smallest element in unsorted part of slice.
		min := i
		for j := i + 1; j < n; j++ {
			if cmp(s[j], s[min]) < 0 {
				min = j
			}
		}

		// Swap current element and the smallest element of unsorted part.
		if i != min {
			s[i], s[min] = s[min], s[i]
		}
	}
}

// InsertionSort sorts slice s in order as determined by cmp function.
// It uses insertion sort algorithm with complexity O(n^2).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func InsertionSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	for i := 1; i < len(s); i++ {
		e := s[i] // current element

		// Search for index of current element in sorted part of slice
		// and shift larger values to the right of it.
		j := i
		for j > 0 && cmp(s[j-1], e) > 0 {
			s[j] = s[j-1]
			j--
		}

		// Insert current element into the found position.
		s[j] = e
	}
}

// QuickSort sorts slice s in order as determined by cmp function.
// It uses recursive quick sort algorithm with complexity O(n*log(n)).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func QuickSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	// If slice contains 0 or 1 element, it is sorted.
	for n := len(s); n >= 2; n = len(s) {
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

		// Use Sedgewick's trick to limit recursive calls and reduce space complexity.
		if left < n-left-1 {
			QuickSort(s[:left], cmp)
			s = s[left+1:]
		} else {
			QuickSort(s[left+1:], cmp)
			s = s[:left]
		}
	}
}
