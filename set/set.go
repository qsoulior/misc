package set

// Обычное множество, основанное на хеш-таблице пустых структур.
type HashSet[T comparable] map[T]struct{}

func (s HashSet[T]) Add(value T) { s[value] = struct{}{} }

func (s HashSet[T]) Remove(value T) { delete(s, value) }

func (s HashSet[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

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

func (s HashSet[T]) Difference(set HashSet[T]) HashSet[T] {
	diff := make(HashSet[T])

	for value := range s {
		if !set.Contains(value) {
			diff.Add(value)
		}
	}

	return diff
}

func (s HashSet[T]) SymmetricDifference(set HashSet[T]) HashSet[T] {
	diff := s.Difference(set)

	for value := range set {
		if !s.Contains(value) {
			diff.Add(value)
		}
	}

	return diff
}

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

func (s HashSet[T]) Superset(set HashSet[T]) bool { return set.Subset(s) }
