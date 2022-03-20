package collection

type Vec[T any] []T

func (v Vec[T]) Len() int {
	return len(v)
}

func (v Vec[T]) Cap() int {
	return cap(v)
}

func (v Vec[T]) Append(item T) Vec[T] {
	return append(v, item)
}

func (v Vec[T]) AppendAll(items ...T) Vec[T] {
	for _, i := range items {
		v = append(v, i)
	}
	return v
}

func (v Vec[T]) AppendIter(it Iterable[T]) Vec[T] {
	for i, ok := it(); ok; i, ok = it() {
		v = append(v, i)
	}
	return v
}

func (v Vec[T]) Iter() Iterable[T] {
	current := 0
	return func() (T, bool) {
		if current < len(v) {
			current++
			return v[current-1], true
		}
		return *new(T), false
	}
}

func (v Vec[T]) ReverseIter() Iterable[T] {
	current := len(v)
	return func() (T, bool) {
		if current > 0 {
			current--
			return v[current], true
		}
		return *new(T), false
	}
}
