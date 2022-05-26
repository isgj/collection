package collection

// LRUCache implements a least recently used cache
type LRUCache[K comparable, V any] struct {
	size   int
	head   *cnode[K, V]
	tail   *cnode[K, V]
	cached map[K]*cnode[K, V]
}

// NewLRUCache creates a new LRUCache.
// If the size is 0 or negative, the cache is unbounded.
func NewCache[K comparable, V any](size int) *LRUCache[K, V] {
	return &LRUCache[K, V]{size: size, cached: make(map[K]*cnode[K, V])}
}

// Clear removes all items from the cache.
func (c *LRUCache[K, V]) Clear() {
	c.cached = make(map[K]*cnode[K, V])
	c.head = nil
	c.tail = nil
}

// Get returns the value for the given key if present in the cache.
func (c *LRUCache[K, V]) Get(key K) (val V, ok bool) {
	node, ok := c.cached[key]
	if !ok {
		return val, ok
	}
	c.moveToHead(node)
	return node.val, ok
}

// GetOrAdd returns the value for the given key if present in the cache.
// If not, it adds the value returned bu f and returns the given value.
func (c *LRUCache[K, V]) GetOrAdd(key K, f func() V) V {
	node, ok := c.Get(key)
	if ok {
		return node
	}
	val := f()
	c.Put(key, val)
	return val
}

// IterKeys returns an iterator over the keys in the cache.
// The keys are returned from the least recently used to the last one.
func (c *LRUCache[K, V]) IterKeys() Iterator[K] {
	cur_node := c.head
	return func() (k K, ok bool) {
		if cur_node == nil {
			return k, false
		}
		k, cur_node = cur_node.key, cur_node.next
		return k, true
	}
}

// IterVals returns an iterator over the values in the cache.
// The values are returned from the least recently used to the last one.
func (c *LRUCache[K, V]) IterVals() Iterator[V] {
	cur_node := c.head
	return func() (v V, ok bool) {
		if cur_node == nil {
			return v, false
		}
		v, cur_node = cur_node.val, cur_node.next
		return v, true
	}
}

// Len returns the number of items in the cache.
func (c *LRUCache[K, V]) Len() int {
	return len(c.cached)
}

// Put adds the given key-value pair to the cache.
func (c *LRUCache[K, V]) Put(key K, val V) {
	node, ok := c.cached[key]
	if ok {
		node.val = val
		c.moveToHead(node)
		return
	}
	if c.size > 0 && len(c.cached) >= c.size {
		c.removeTail()
	}
	node = &cnode[K, V]{key: key, val: val}
	c.cached[key] = node
	// Add the first node
	if c.head == nil {
		c.head, c.tail = node, node
		return
	}
	c.moveToHead(node)
}

// ReverseIterKeys returns an iterator over the keys in the cache.
// The keys are returned from the last used to the least recently one.
func (c *LRUCache[K, V]) ReverseIterKeys() Iterator[K] {
	cur_node := c.tail
	return func() (k K, ok bool) {
		if cur_node == nil {
			return k, false
		}
		k, cur_node = cur_node.key, cur_node.prev
		return k, true
	}
}

// IterVals returns an iterator over the values in the cache.
// The values are returned from the last used to the least recently one.
func (c *LRUCache[K, V]) ReverseIterVals() Iterator[V] {
	cur_node := c.tail
	return func() (v V, ok bool) {
		if cur_node == nil {
			return v, false
		}
		v, cur_node = cur_node.val, cur_node.prev
		return v, true
	}
}

func (c *LRUCache[K, V]) moveToHead(node *cnode[K, V]) {
	if node == c.head {
		return
	}
	if node == c.tail {
		c.tail = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev = nil
	node.next = c.head
	c.head.prev = node
	c.head = node
}

func (c *LRUCache[K, V]) removeTail() {
	if c.tail == nil {
		return
	}
	if c.tail.prev != nil {
		c.tail.prev.next = nil
	}
	delete(c.cached, c.tail.key)
	c.tail = c.tail.prev
}

type cnode[K comparable, V any] struct {
	key  K
	val  V
	prev *cnode[K, V]
	next *cnode[K, V]
}
