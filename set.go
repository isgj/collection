package collection

// Set is the classic `set` data structure
type Set[T comparable] Map[T, struct{}]

// Add will add the element to the set
func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

// Clear will delete all the set elements.
func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Delete will remove the element from the set
func (s Set[T]) Delete(elem T) {
	delete(s, elem)
}

// Difference will return a new set that will contain only the elements of the receiver that are not in the other
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := Set[T]{}
	for elem := range s {
		if !other.Has(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Has checks if the element is in the set.
func (m Set[T]) Has(element T) bool {
	_, ok := m[element]
	return ok
}

// Intersection will return a new set with all the elements that are part of both the sets
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := Set[T]{}
	if s.Len() > other.Len() {
		s, other = other, s
	}
	for elem := range s {
		if other.Has(elem) {
			result.Add(elem)
		}
	}
	return result
}

// IsEmpty checks if the set is empty
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// Len is an alias to `len`
func (s Set[T]) Len() int {
	return len(s)
}

// ToVec will collect the elements of the set to a `Vec`
func (s Set[T]) ToVec() Vec[T] {
	v := make(Vec[T], 0, s.Len())
	for e := range s {
		v = append(v, e)
	}
	return v
}

// Union returns a new set with all the the elements of both sets
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := Set[T]{}
	for e := range s {
		result.Add(e)
	}
	for e := range other {
		result.Add(e)
	}
	return result
}

// NewSet returns a new set from the list of values
func NewSet[T comparable](values ...T) Set[T] {
	set := Set[T]{}
	for _, v := range values {
		set.Add(v)
	}
	return set
}

// NewSetFromIter will collect all the values of the iterator to a set
func NewSetFromIter[T comparable](it Iterator[T]) Set[T] {
	set := Set[T]{}
	for v, ok := it(); ok; v, ok = it() {
		set.Add(v)
	}
	return set
}
