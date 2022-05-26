package collection

import "testing"

func TestBack(t *testing.T) {
	var queue DLList[int]
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	s := queue.Size()
	if v, ok := queue.Back(); !ok || v != 3 {
		t.Errorf("queue.Back() = %v, %v, want %v, %v", v, ok, 3, true)
	}
	if s != queue.Size() {
		t.Errorf("queue.size cahnged was = %d, now = %d", s, queue.Size())
	}
}

func TestClear(t *testing.T) {
	var queue DLList[int]
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	queue.Clear()
	if !queue.IsEmpty() {
		t.Errorf("queue.IsEmpty() = %v, want %v", queue.IsEmpty(), true)
	}
	if queue.head != nil {
		t.Errorf("queue.head = %v, want %v", queue.head, nil)
	}
	if queue.tail != nil {
		t.Errorf("queue.tail = %v, want %v", queue.tail, nil)
	}
}

func TestFront(t *testing.T) {
	var queue DLList[int]
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	s := queue.Size()
	if v, ok := queue.Front(); !ok || v != 1 {
		t.Errorf("queue.Front() = %v, %v, want %v, %v", v, ok, 1, true)
	}
	if s != queue.Size() {
		t.Errorf("queue.size cahnged was = %d, now = %d", s, queue.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	var queue DLList[int]
	if !queue.IsEmpty() {
		t.Errorf("queue.IsEmpty() = %v, want %v", queue.IsEmpty(), true)
	}
	queue.PushBack(1)
	if queue.IsEmpty() {
		t.Errorf("queue.IsEmpty() = %v, want %v", queue.IsEmpty(), false)
	}
}

func TestDLLIter(t *testing.T) {
	var queue DLList[int]
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	var i int
	queue.Iter().ForEach(func(v int) {
		if v != i+1 {
			t.Errorf("queue.Iter().ForEach() = %v, want %v", v, i+1)
		}
		i++
	})
	if i != 3 {
		t.Errorf("queue.Iter().ForEach() = %v, want %v", i, 3)
	}
	i = 3
	queue.ReverseIter().ForEach(func(v int) {
		if v != i {
			t.Errorf("queue.ReverseIter().ForEach() = %v, want %v", v, i)
		}
		i--
	})
	if i != 0 {
		t.Errorf("queue.ReverseIter().ForEach() = %v, want %v", i, 0)
	}
}

func TestPopBack(t *testing.T) {
	var queue DLList[int]
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	s := queue.Size()
	if v, ok := queue.PopBack(); !ok || v != 3 {
		t.Errorf("queue.PopBack() = %v, %v, want %v, %v", v, ok, 3, true)
	}
	if s != queue.Size()+1 {
		t.Errorf("queue.size did not change, before = %d, now = %d", s, queue.Size())
	}
}

func TestPopFront(t *testing.T) {
	var queue DLList[int]
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	s := queue.Size()
	if v, ok := queue.PopFront(); !ok || v != 1 {
		t.Errorf("queue.PopFront() = %v, %v, want %v, %v", v, ok, 1, true)
	}
	if s != queue.Size()+1 {
		t.Errorf("queue.size did not change, before = %d, now = %d", s, queue.Size())
	}
}

func TestPushBack(t *testing.T) {
	var queue DLList[int]
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	s := queue.Size()
	if s != 3 {
		t.Errorf("queue.size = %d, want %d", s, 3)
	}
	if queue.head.value != 1 {
		t.Errorf("queue.head.value = %d, want %d", queue.head.value, 1)
	}
	if queue.tail.value != 3 {
		t.Errorf("queue.tail.value = %d, want %d", queue.tail.value, 3)
	}
}

func TestPushFront(t *testing.T) {
	var queue DLList[int]
	queue.PushFront(1)
	queue.PushFront(2)
	queue.PushFront(3)
	s := queue.Size()
	if s != 3 {
		t.Errorf("queue.size = %d, want %d", s, 3)
	}
	if queue.head.value != 3 {
		t.Errorf("queue.head.value = %d, want %d", queue.head.value, 3)
	}
	if queue.tail.value != 1 {
		t.Errorf("queue.tail.value = %d, want %d", queue.tail.value, 1)
	}
}

func TestLen(t *testing.T) {
	var queue DLList[int]
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	s := queue.Len()
	if s != 3 {
		t.Errorf("queue.len = %d, want %d", s, 3)
	}
}
