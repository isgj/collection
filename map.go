package collection

// Map is the same as `map` but with some methods.
type Map[K comparable, V any] map[K]V

// Clear will delete all the key/value pairs in the map.
func (m Map[K, V]) Clear() {
	for k := range m {
		delete(m, k)
	}
}

// Has checks if the key is in the map.
func (m Map[K, V]) Has(key K) bool {
	_, ok := m[key]
	return ok
}

// Keys will return a Vec with the keys of the map.
func (m Map[K, V]) Keys() Vec[K] {
	v := make(Vec[K], 0, len(m))
	for k := range m {
		v = append(v, k)
	}
	return v
}

// Len is an alias to the `len` function.
func (m Map[K, V]) Len() int {
	return len(m)
}

// IsEmpty checks if the map has no entries
func (m Map[K, V]) IsEmpty() bool {
	return len(m) == 0
}

// Values will return a Vec with the values of the map.
func (m Map[K, V]) Values() Vec[V] {
	v := make(Vec[V], 0, len(m))
	for _, val := range m {
		v = append(v, val)
	}
	return v
}
