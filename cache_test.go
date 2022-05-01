package collection

import "testing"

func TestNewCacheWithLimit(t *testing.T) {
	cache := NewCache[int, int](3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	if cache.Len() != 3 {
		t.Errorf("cache.Size() = %d, want %d", cache.Len(), 3)
	}
}

func TestLeastUsedIsEvicted(t *testing.T) {
	cache := NewCache[int, int](3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Get(1) // 2 becomes the least used
	cache.Put(4, 4)
	if v, ok := cache.Get(2); ok {
		t.Errorf("cache.Get(1) = %d, want %d", v, 0)
	}
}

func TestCacheClear(t *testing.T) {
	cache := NewCache[int, int](3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Clear()
	if cache.Len() != 0 {
		t.Errorf("cache.Size() = %d, want %d", cache.Len(), 0)
	}
}

func TestCacheGet(t *testing.T) {
	cache := NewCache[int, int](3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4) // 1 is evicted
	if v, ok := cache.Get(1); ok {
		t.Errorf("cache.Get(1) = %d, %t, want %d, %t", v, ok, 0, false)
	}
	if v, ok := cache.Get(2); !ok || v != 2 {
		t.Errorf("cache.Get(2) = %d, %t, want %d, %t", v, ok, 2, true)
	}
	if v, ok := cache.Get(3); !ok || v != 3 {
		t.Errorf("cache.Get(3) = %d, %t, want %d, %t", v, ok, 3, true)
	}
	if v, ok := cache.Get(4); !ok || v != 4 {
		t.Errorf("cache.Get(4) = %d, %t, want %d, %t", v, ok, 4, true)
	}
}

func TestCacheGetOrAdd(t *testing.T) {
	cache := NewCache[int, int](3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	if v, ok := cache.Get(1); !ok || v != 1 {
		t.Errorf("cache.Get(1) = %d, %t, want %d, %t", v, ok, 1, true)
	}
	if v := cache.GetOrAdd(1, func() int { return 10 }); v != 1 {
		t.Errorf("cache.GetOrAdd(1) = %d, want %d", v, 1)
	}
	if v, ok := cache.Get(1); !ok || v != 1 {
		t.Errorf("cache.Get(1) = %d, %t, want %d, %t", v, ok, 1, true)
	}
	if v := cache.GetOrAdd(4, func() int { return 10 }); v != 10 {
		t.Errorf("cache.GetOrAdd(4) = %d, want %d", v, 10)
	}
	if v, ok := cache.Get(4); !ok || v != 10 {
		t.Errorf("cache.Get(4) = %d, %t, want %d, %t", v, ok, 10, false)
	}
}

func TestCacheLen(t *testing.T) {
	cache := NewCache[int, int](3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	if cache.Len() != 3 {
		t.Errorf("cache.Len() = %d, want %d", cache.Len(), 3)
	}
	cache.Put(4, 4) // 1 is evicted
	if cache.Len() != 3 {
		t.Errorf("cache.Len() = %d, want %d", cache.Len(), 3)
	}
}

func TestCachePutLeastUsed(t *testing.T) {
	cache := NewCache[int, int](3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(1, 10) // value of key 1 should be 10, and it should be the head
	cache.Put(4, 4)  // 2 is evicted
	cache.Put(5, 5)  // 3 is evicted
	if v, ok := cache.Get(1); !ok || v != 10 {
		t.Errorf("cache.Get(1) = %d, %t, want %d, %t", v, ok, 10, true)
	}
	if v, ok := cache.Get(2); ok {
		t.Errorf("cache.Get(2) = %d, %t, want %d, %t", v, ok, 0, false)
	}
	if v, ok := cache.Get(3); ok {
		t.Errorf("cache.Get(3) = %d, %t, want %d, %t", v, ok, 0, false)
	}
	if v, ok := cache.Get(4); !ok || v != 4 {
		t.Errorf("cache.Get(4) = %d, %t, want %d, %t", v, ok, 4, true)
	}
	if v, ok := cache.Get(5); !ok || v != 5 {
		t.Errorf("cache.Get(4) = %d, %t, want %d, %t", v, ok, 5, true)
	}
	if cache.Len() != 3 {
		t.Errorf("cache.Len() = %d, want %d", cache.Len(), 3)
	}
}

func TestCachePutLastUsed(t *testing.T) {
	cache := NewCache[int, int](3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(3, 10)
	cache.Put(4, 4) // 1 is evicted
	cache.Put(5, 5) // 2 is evicted
	if v, ok := cache.Get(1); ok {
		t.Errorf("cache.Get(1) = %d, %t, want %d, %t", v, ok, 0, false)
	}
	if v, ok := cache.Get(2); ok {
		t.Errorf("cache.Get(2) = %d, %t, want %d, %t", v, ok, 0, false)
	}
	if v, ok := cache.Get(3); !ok || v != 10 {
		t.Errorf("cache.Get(3) = %d, %t, want %d, %t", v, ok, 10, true)
	}
	if v, ok := cache.Get(4); !ok || v != 4 {
		t.Errorf("cache.Get(4) = %d, %t, want %d, %t", v, ok, 4, true)
	}
	if v, ok := cache.Get(5); !ok || v != 5 {
		t.Errorf("cache.Get(4) = %d, %t, want %d, %t", v, ok, 5, true)
	}
	if cache.Len() != 3 {
		t.Errorf("cache.Len() = %d, want %d", cache.Len(), 3)
	}
}
