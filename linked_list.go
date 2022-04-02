package collection

// DLList is a doubly linked list. It is based on the linked list implementation.
// It can be used as a stack and/or a queue.
// All the operations are O(1), even when the size of the list is large.
//
// The zero value for DLList is an empty list ready to use.
//    var queue collection.DLList[int]
//    queue.PushBack(1)
//    queue.PushBack(2)
//    queue.PushBack(3)
//    fmt.Println(queue.PopFront()) // 1, true
//    fmt.Println(queue.PopFront()) // 2, true
//    fmt.Println(queue.PopFront()) // 3, true
//    fmt.Println(queue.PopFront()) // 0, false
type DLList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// Back returns the last element of the list.
// If the list is empty, the zero value is returned and false.
func (ll *DLList[T]) Back() (T, bool) {
	if ll.size == 0 {
		return *new(T), false
	}
	return ll.tail.value, true
}

// Clear removes all elements from the list.
func (ll *DLList[T]) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

// Front returns the first element of the list.
// If the list is empty, the zero value is returned and false.
func (ll *DLList[T]) Front() (T, bool) {
	if ll.size == 0 {
		return *new(T), false
	}
	return ll.head.value, true
}

// IsEmpty returns true if the list is empty.
func (ll *DLList[T]) IsEmpty() bool {
	return ll.size == 0
}

// Iter returns a new iterator for the list, iterating the values from front to back.
func (ll *DLList[T]) Iter() Iterator[T] {
	cur_node := ll.head
	return func() (T, bool) {
		if cur_node == nil {
			return *new(T), false
		}
		v := cur_node.value
		cur_node = cur_node.next
		return v, true
	}
}

// PopBack removes the last element from the list.
// If the second return value is false, the list is empty and the zero value is returned.
func (ll *DLList[T]) PopBack() (T, bool) {
	if ll.size == 0 {
		return *new(T), false
	}
	n := ll.tail
	if ll.size == 1 {
		ll.head = nil
		ll.tail = nil
	} else {
		ll.tail = n.prev
		ll.tail.next = nil
	}
	ll.size--
	return n.value, true
}

// PopFront removes the first element from the list.
// If the second return value is false, the list is empty and the zero value is returned.
func (ll *DLList[T]) PopFront() (T, bool) {
	if ll.size == 0 {
		return *new(T), false
	}
	n := ll.head
	if ll.size == 1 {
		ll.head = nil
		ll.tail = nil
	} else {
		ll.head = n.next
		ll.head.prev = nil
	}
	ll.size--
	return n.value, true
}

// PushBack adds a new element at the back of the list.
func (ll *DLList[T]) PushBack(v T) {
	n := &node[T]{value: v}
	if ll.size == 0 {
		ll.head = n
		ll.tail = n
	} else {
		ll.tail.next = n
		n.prev = ll.tail
		ll.tail = n
	}
	ll.size++
}

// PushFront adds a new element at the front of the list.
func (ll *DLList[T]) PushFront(v T) {
	n := &node[T]{value: v}
	if ll.size == 0 {
		ll.head = n
		ll.tail = n
	} else {
		ll.head.prev = n
		n.next = ll.head
		ll.head = n
	}
	ll.size++
}

// ReverseIter returns a new iterator for the list, iterating the values from back to front.
func (ll *DLList[T]) ReverseIter() Iterator[T] {
	cur_node := ll.tail
	return func() (T, bool) {
		if cur_node == nil {
			return *new(T), false
		}
		v := cur_node.value
		cur_node = cur_node.prev
		return v, true
	}
}

// Size returns the number of elements in the list.
func (ll *DLList[T]) Size() int {
	return ll.size
}

// node is a helper struct that holds the value and the links to the next and previous nodes.
type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}
