// Utility iterator functions.
package iter

import (
	c "github.com/isgj/collection"
	"golang.org/x/exp/constraints"
)

// Map will map values of type `I` from to type `O` through the `to` mapper function.
//
// Example:
// 		v := collection.Vec[int]{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 		number_iterator := v.Iter().Filter(func(item int) string { return item%2 == 0})
// 		iter.Map(number_iterator, func(item int) string { return fmt.Sprint("number ", item) }).
// 			ForEach(func(item string) { fmt.Prinln(item) })
//
// This is a workaround to implement `Map` as currently methods cannot have type paramters.
func Map[I any, O any](it c.Iterator[I], to func(item I) O) c.Iterator[O] {
	return func() (O, bool) {
		if i, ok := it(); ok {
			return to(i), true
		}
		return *new(O), false
	}
}

// Reduce will reduce the values through the reducer
//
// This is a workaround to implement `Reduce` as currently methods cannot have type paramters.
func Reduce[I any, O any](it c.Iterator[I], start O, reducer func(acc O, item I) O) O {
	for i, ok := it(); ok; i, ok = it() {
		start = reducer(start, i)
	}
	return start
}

// Sum will return the sum of the values. If the iterated values are strings it will concatenate them
func Sum[T constraints.Ordered](it c.Iterator[T]) T {
	result, _ := it()
	for i, ok := it(); ok; i, ok = it() {
		result += i
	}
	return result
}

func FromSlice[T any](s []T) c.Iterator[T] {
	return c.Vec[T](s).Iter()
}

// Iterate from 0 to `end` by steps of 1
func Range(end int) c.Iterator[int] {
	return XRange(0, end, 1)
}

// Iterate from `start` to `end` by steps of `step`.
// Currently the range will only increase, the step will be added for each iteration so make sure that it will
// eventually reach the end, otherwise you enter an infinite loop.
func XRange(start, end, step int) c.Iterator[int] {
	return func() (int, bool) {
		if start < end {
			r := start
			start += step
			return r, true
		}
		return 0, false
	}
}
