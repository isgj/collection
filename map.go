package collection

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Keys() Vec[K] {
	v := make(Vec[K], 0, len(m))
	for k := range m {
		v = append(v, k)
	}
	return v
}

func (m Map[K, V]) Values() Vec[V] {
	v := make(Vec[V], 0, len(m))
	for _, val := range m {
		v = append(v, val)
	}
	return v
}
