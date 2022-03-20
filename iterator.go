package collection

type Iterator[T any] func() (T, bool)

func (it Iterator[T]) Any(test func(item T) bool) bool {
	for i, ok := it(); ok; i, ok = it() {
		if test(i) {
			return true
		}
	}
	return false
}

func (it Iterator[T]) Collect() Vec[T] {
	var vec Vec[T]
	for i, ok := it(); ok; i, ok = it() {
		vec = append(vec, i)
	}
	return vec
}

func (it Iterator[T]) Every(test func(item T) bool) bool {
	for i, ok := it(); ok; i, ok = it() {
		if !test(i) {
			return false
		}
	}
	return true
}

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

func (it Iterator[T]) Find(test func(item T) bool) (T, bool) {
	for i, ok := it(); ok; i, ok = it() {
		if test(i) {
			return i, ok
		}
	}
	return *new(T), false
}

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

func (it Iterator[T]) ForEach(action func(item T)) {
	for i, ok := it(); ok; i, ok = it() {
		action(i)
	}
}

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

func (it Iterator[T]) Reverse() Iterator[T] {
	return it.Collect().ReverseIter()
}
