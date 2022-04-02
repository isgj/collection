// collection implements generic data structures in GO.
// collection.Vec and collection.Map are the same as native `slice` and `map` but with some methods.
// collection.Set is a set of comparable values. It is implemented as a map where the values are an empty struct.
// collection.DLList is a doubly linked list which can be used as a stack or a queue. It has fast O(1) operations in both ends.
// collection.Iterator is a lazy generic iterator. It is not limited to the structures defined in this package.
package collection
