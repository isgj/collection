[![Go Reference](https://pkg.go.dev/badge/github.com/isgj/collection.svg)](https://pkg.go.dev/github.com/isgj/collection)

# collection

Generic go structures

## Install

```
go get github.com/isgj/collection
```

## Usage

[collection.Vec[T]](https://pkg.go.dev/github.com/isgj/collection#Vec) implemented as a native go slice `[]T`. Because of this, other than the few currently
implemented methods you can use `Vec` also as a regular slice.

```go
package main

import (
	"fmt"

	"github.com/isgj/collection"
)

func main() {
	strings := collection.Vec[string]{"str1", "str2"}

	for i := 3; i < 6; i++ {
		strings = append(strings, fmt.Sprintf("str%d", i))
	}

	sliced := strings[2:4]
	indexed := strings[0]
	fmt.Printf("len=%d, cap=%d, indexed=%s, sliced=%v\n", len(strings), cap(strings), indexed, sliced)
}
// Output:
// len=5, cap=8, indexed=str1, sliced=[str3 str4]
```

The most noticable method of `Vec` is `Iter` (or `ReverseIter`) which returns a lazy iterator over the
elements of the slice. Check below.

---

[collection.Iterator[T]](https://pkg.go.dev/github.com/isgj/collection#Iterator) is a lazy iterator over a list of values of type `T`.

```go
package main

import (
	"fmt"

	"github.com/isgj/collection"
)

func main() {
	first_slice := collection.Vec[int]{1, 2, 3, 4, 5}
	second_slice := collection.Vec[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14}

	result := first_slice.
		Iter().
		FollowedBy(second_slice.ReverseIter()).
		Filter(func(item int) bool {
			fmt.Printf("will test item: %2d\n", item)
			return item%2 == 0
		}).
		Skip(2).
		Take(3).
		Collect()

	fmt.Printf("len=%d, cap=%d, vec=%v\n", result.Len(), result.Cap(), result)
}
// Output:
// will test item:  1
// will test item:  2
// will test item:  3
// will test item:  4
// will test item:  5
// will test item: 14
// will test item: 12
// will test item: 10
// len=3, cap=4, vec=[14 12 10]
```

> To note: because the iterator is lazy the test of `Filter` is not called on each element, but only as much as needed (no wasted calls).
> At the end you need to call `Collect` to get back a slice. If `Collect` is not called nothing gets executed.

---

[collection.Map](https://pkg.go.dev/github.com/isgj/collection#Map)

[collection.Set](https://pkg.go.dev/github.com/isgj/collection#Set)

[collection.DLList](https://pkg.go.dev/github.com/isgj/collection#DLList)

[collection.LRUCache](https://pkg.go.dev/github.com/isgj/collection#LRUCache)
