// Package slice implements various operations with slices.
// It provides sort and search algorithms.
package slice

import (
	"math/rand"
)

// BubbleSort sorts slice s in order as determined by cmp function.
// It uses bubble sorting algorithm with complexity O(n^2).
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
// It uses cocktail sorting (bidirectional bubble sorting) algorithm with complexity O(n^2).
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
// It uses comb sorting (modification of bubble sorting) algorithm with complexity O(n^2),
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
// It uses selection sorting algorithm with complexity O(n^2).
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
// It uses insertion sorting algorithm with complexity O(n^2).
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
// It uses recursive quick sorting algorithm with complexity O(n*log(n)).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func QuickSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	// If slice contains 0 or 1 elements, it is sorted.
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

// MergeSort sorts slice s in order as determined by cmp function.
// It uses recursive merge sorting algorithm with complexity O(n*log(n)).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func MergeSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	n := len(s)
	// If slice contains 0 or 1 elements, it is sorted.
	if n < 2 {
		return
	}

	m := n / 2
	MergeSort(s[:m], cmp) // sort left subslice
	MergeSort(s[m:], cmp) // sort right subslice

	buf := make(S, n)
	k := 0 // index of buffer element
	i := 0 // index of element in left subslice
	j := m // index of element in right subslice

	// Merge two subslices into one buffer.
	for i < m || j < n {
		if j >= n || (i < m && cmp(s[i], s[j]) < 0) {
			buf[k] = s[i]
			i++
		} else {
			buf[k] = s[j]
			j++
		}
		k++
	}

	// Copy elements from buffer to result slice.
	copy(s, buf)
}

// siftDown moves element of slice s with index i to correct position in max heap.
// It compares and swaps element with one of its children.
// Used in heap sorting algorithm.
func siftDown[S ~[]E, E any](s S, i int, cmp func(a E, b E) int) {
	n := len(s)
	for 2*i+1 < n {
		j := 2*i + 1 // index of the largest child (left by default)
		if r := j + 1; r < n && cmp(s[r], s[j]) > 0 {
			j = r // index of right child
		}

		// If parent is greater than or equal to the largest child,
		// then it has correct position.
		if cmp(s[i], s[j]) >= 0 {
			break
		}

		// Swap parent with the largest child.
		s[i], s[j] = s[j], s[i]
		i = j
	}
}

// HeapSort sorts slice s in order as determined by cmp function.
// It uses heap sorting algorithm with complexity O(n*log(n)).
// cmp should return 0 if a is equal b, a negative number if a precedes b,
// or a positive number if a follows b.
func HeapSort[S ~[]E, E any](s S, cmp func(a E, b E) int) {
	n := len(s)
	// Build max heap from slice with complexity O(n).
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(s, i, cmp)
	}

	for i := n - 1; i >= 0; i-- {
		// Swap the largest element with the last one.
		s[0], s[i] = s[i], s[0]

		// Move the new first element to correct position in heap.
		siftDown(s[:i], 0, cmp)
	}
}
