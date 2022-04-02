package collection

// Vec is a generic slice, the same rules of the native slice aplly also to `Vec`.
type Vec[T any] []T

// Len returns the length of the slice. It's an alias to the global `len`
func (v Vec[T]) Len() int {
	return len(v)
}

// Cap returns the length of the slice. It's an alias to the global `cap`
func (v Vec[T]) Cap() int {
	return cap(v)
}

// Append will add the item at the end of the slice. It's an alias to the global `append`.
// Remember to assign the returned value, as you should do with `append`.
func (v Vec[T]) Append(item T) Vec[T] {
	return append(v, item)
}

// AppendAll will add all the items at the end of the slice.
func (v Vec[T]) AppendAll(items ...T) Vec[T] {
	for _, i := range items {
		v = append(v, i)
	}
	return v
}

// AppendIter will consume the iterator and append the yield values.
func (v Vec[T]) AppendIter(it Iterator[T]) Vec[T] {
	for i, ok := it(); ok; i, ok = it() {
		v = append(v, i)
	}
	return v
}

// Iter will return a lazy iterator over the values of the slice.
func (v Vec[T]) Iter() Iterator[T] {
	current := 0
	return func() (T, bool) {
		if current < len(v) {
			current++
			return v[current-1], true
		}
		return *new(T), false
	}
}

// IsEmpty checks if the slice has no entries.
func (v Vec[T]) IsEmpty() bool {
	return len(v) == 0
}

// ReverseIter will return a lazy iterator over the values of the slice in reverse order.
func (v Vec[T]) ReverseIter() Iterator[T] {
	current := len(v)
	return func() (T, bool) {
		if current > 0 {
			current--
			return v[current], true
		}
		return *new(T), false
	}
}
