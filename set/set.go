package set

// HashSet implements set based on hash table of empty structs.
type HashSet[T comparable] map[T]struct{}

// Len returns number of elements contained in set.
func (s HashSet[T]) Len() int { return len(s) }

// Add inserts value into set.
func (s HashSet[T]) Add(value T) { s[value] = struct{}{} }

// Remove removes value from set.
func (s HashSet[T]) Remove(value T) { delete(s, value) }

// Contains returns true if value is contained in set.
func (s HashSet[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

// Union returns new set with all elements from sets.
func (s HashSet[T]) Union(set HashSet[T]) HashSet[T] {
	length, l := len(set), len(s)
	if l > length {
		length = l
	}

	union := make(HashSet[T], length)
	for value := range s {
		union.Add(value)
	}

	for value := range set {
		union.Add(value)
	}

	return union
}

// Intersection returns new set with elements common to sets.
func (s HashSet[T]) Intersection(set HashSet[T]) HashSet[T] {
	length, l := len(set), len(s)
	if l < length {
		length = l
	}

	intersection := make(HashSet[T], length)

	if l == length {
		for value := range s {
			if set.Contains(value) {
				intersection.Add(value)
			}
		}
	} else {
		for value := range set {
			if s.Contains(value) {
				intersection.Add(value)
			}
		}
	}

	return intersection
}

// Difference returns new set with elements from s that are not in set.
func (s HashSet[T]) Difference(set HashSet[T]) HashSet[T] {
	diff := make(HashSet[T])

	for value := range s {
		if !set.Contains(value) {
			diff.Add(value)
		}
	}

	return diff
}

// SymmetricDifference returns new set with elements in either s or set but not both.
func (s HashSet[T]) SymmetricDifference(set HashSet[T]) HashSet[T] {
	diff := s.Difference(set)

	for value := range set {
		if !s.Contains(value) {
			diff.Add(value)
		}
	}

	return diff
}

// Equal returns true if s contains every element of set
// and their lengths are equal.
func (s HashSet[T]) Equal(set HashSet[T]) bool {
	if len(s) != len(set) {
		return false
	}

	for value := range s {
		if !set.Contains(value) {
			return false
		}
	}

	return true
}

// Subset returns true if set contains every element of s.
func (s HashSet[T]) Subset(set HashSet[T]) bool {
	if len(s) > len(set) {
		return false
	}

	for value := range s {
		if !set.Contains(value) {
			return false
		}
	}

	return true
}

// Superset returns true if s contains every element of set.
func (s HashSet[T]) Superset(set HashSet[T]) bool { return set.Subset(s) }
