package collection

// Iterator is a lazy iterator over generic data types.
// It can be called several times to produce values.
// When the second returned value is `true` means the value is valid and it can be consumed.
// When the second returned value is `false` means the value is not valid.
// In this case the zero value of the type `T` is returned and it should not be consumed.
// Consecutive calls after the first time `false` is returned, should return the same values.
type Iterator[T any] func() (T, bool)

// Any returns true as soon as a value satisfies the test, false otherwise.
func (it Iterator[T]) Any(test func(item T) bool) bool {
	for i, ok := it(); ok; i, ok = it() {
		if test(i) {
			return true
		}
	}
	return false
}

// Collect will consume the iterator and return a `Vec` with all the values.
func (it Iterator[T]) Collect() Vec[T] {
	var vec Vec[T]
	for i, ok := it(); ok; i, ok = it() {
		vec = append(vec, i)
	}
	return vec
}

// Count will consume the iterator and return the number of values iterated.
func (it Iterator[T]) Count() int {
	var c int
	for _, ok := it(); ok; _, ok = it() {
		c++
	}
	return c
}

// Every will return false as soon as a value will fail the test, true otherwise.
func (it Iterator[T]) Every(test func(item T) bool) bool {
	for i, ok := it(); ok; i, ok = it() {
		if !test(i) {
			return false
		}
	}
	return true
}

// Filter will pass only the values that satisfy the test.
func (it Iterator[T]) Filter(test func(item T) bool) Iterator[T] {
	return func() (T, bool) {
		for i, ok := it(); ok; i, ok = it() {
			if test(i) {
				return i, ok
			}
		}
		return *new(T), false
	}
}

// Find will try to find a value that satisfies the test.
// The second returned value is true if a value was found, false otherwise.
func (it Iterator[T]) Find(test func(item T) bool) (T, bool) {
	for i, ok := it(); ok; i, ok = it() {
		if test(i) {
			return i, ok
		}
	}
	return *new(T), false
}

// FollowedBy will yield first the values of `it` followed by the values of `other`.
func (it Iterator[T]) FollowedBy(other Iterator[T]) Iterator[T] {
	other_turn := false
	return func() (T, bool) {
		if other_turn {
			return other()
		}
		if i, ok := it(); ok {
			return i, ok
		}
		other_turn = true
		return other()
	}
}

// ForEach will consume the iterator and run the action with every value.
func (it Iterator[T]) ForEach(action func(item T)) {
	for i, ok := it(); ok; i, ok = it() {
		action(i)
	}
}

// Skip will skip the first `count` values
func (it Iterator[T]) Skip(count int) Iterator[T] {
	skipped := false
	return func() (T, bool) {
		if skipped {
			return it()
		}
		for i := 0; i < count; i++ {
			if _, ok := it(); !ok {
				skipped = true
				return *new(T), false
			}
		}
		skipped = true
		return it()
	}
}

// SkipWhile will skip the first elements that satisfy the test.
func (it Iterator[T]) SkipWhile(test func(item T) bool) Iterator[T] {
	skipped := false
	return func() (T, bool) {
		if !skipped {
			skipped = true
			return it.Find(func(item T) bool { return !test(item) })
		}
		return it()
	}
}

// Take will yield at most the first `count` values.
func (it Iterator[T]) Take(count int) Iterator[T] {
	taken := 0
	return func() (T, bool) {
		if taken >= count {
			return *new(T), false
		}
		if i, ok := it(); ok {
			taken++
			return i, ok
		}
		taken = count
		return *new(T), false
	}
}

// TakeWhile will stop at the first value that does not satisfy the test.
func (it Iterator[T]) TakeWhile(test func(item T) bool) Iterator[T] {
	stopped := false
	return func() (T, bool) {
		if stopped {
			return *new(T), false
		}
		if i, ok := it(); ok && test(i) {
			return i, ok
		}
		stopped = true
		return *new(T), false
	}
}

// Tap will run `action` with every value that will pass through the iterator.
func (it Iterator[T]) Tap(action func(item T)) Iterator[T] {
	return func() (T, bool) {
		i, ok := it()
		if ok {
			action(i)
		}
		return i, ok
	}
}

// Reverse will consume the iterator, collect the values in a `Vec` and iterate in reverse those values.
// Since `Reverse` will consume the iterator and allocate a `Vec`, when possible use `Vec.ReverseIter`.
func (it Iterator[T]) Reverse() Iterator[T] {
	return it.Collect().ReverseIter()
}
