// Package set implements set data structures.
// It provides hash set implementation with core operations.
package set

// HashSet implements set based on hash table of empty structs.
type HashSet[T comparable] map[T]struct{}

// Len returns number of elements contained in set, O(1).
func (s HashSet[T]) Len() int { return len(s) }

// Add inserts value into set, O(1).
func (s HashSet[T]) Add(value T) { s[value] = struct{}{} }

// Remove removes value from set, O(1).
func (s HashSet[T]) Remove(value T) { delete(s, value) }

// Contains returns true if value is contained in set, O(1).
func (s HashSet[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

// Union returns new set with all elements from sets,
// complexity is O(n+m) where n is length of s and m is length of t.
func (s HashSet[T]) Union(t HashSet[T]) HashSet[T] {
	lt, ls := len(t), len(s)
	if ls > lt {
		lt = ls
	}

	union := make(HashSet[T], lt)
	for value := range s {
		union.Add(value)
	}

	for value := range t {
		union.Add(value)
	}

	return union
}

// Intersection returns new set with elements common to sets,
// complexity is O(n) where n is length of smaller set.
func (s HashSet[T]) Intersection(t HashSet[T]) HashSet[T] {
	lt, ls := len(t), len(s)
	if ls < lt {
		lt = ls
	}

	intersection := make(HashSet[T], lt)

	if ls == lt {
		for value := range s {
			if t.Contains(value) {
				intersection.Add(value)
			}
		}
	} else {
		for value := range t {
			if s.Contains(value) {
				intersection.Add(value)
			}
		}
	}

	return intersection
}

// Difference returns new set with elements from s that are not in t,
// complexity is O(n) where n is length of s.
func (s HashSet[T]) Difference(t HashSet[T]) HashSet[T] {
	diff := make(HashSet[T])

	for value := range s {
		if !t.Contains(value) {
			diff.Add(value)
		}
	}

	return diff
}

// SymmetricDifference returns new set with elements in either s or t but not both,
// complexity is O(n) where n is length of s.
func (s HashSet[T]) SymmetricDifference(t HashSet[T]) HashSet[T] {
	diff := s.Difference(t)

	for value := range t {
		if !s.Contains(value) {
			diff.Add(value)
		}
	}

	return diff
}

// Equal returns true if s contains every element of t and their lengths are equal,
// complexity is O(n), where n is length of s.
func (s HashSet[T]) Equal(t HashSet[T]) bool {
	if len(s) != len(t) {
		return false
	}

	for value := range s {
		if !t.Contains(value) {
			return false
		}
	}

	return true
}

// Subset returns true if t contains every element of s,
// complexity is O(n) where n is length of s.
func (s HashSet[T]) Subset(t HashSet[T]) bool {
	if len(s) > len(t) {
		return false
	}

	for value := range s {
		if !t.Contains(value) {
			return false
		}
	}

	return true
}

// Superset returns true if s contains every element of set,
// complexity is O(n) where n is length of t.
func (s HashSet[T]) Superset(t HashSet[T]) bool { return t.Subset(s) }
