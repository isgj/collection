package collection

import "testing"

func TestAppendAll(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	if a.Len() != 3 {
		t.Errorf("expected 3, got %d", a.Len())
	}
}

func TestAppenIter(t *testing.T) {
	var a Vec[int]
	i := -1

	a = a.AppendIter(func() (int, bool) {
		i++
		return 1, i < 3
	})
	if a.Len() != 3 {
		t.Errorf("expected 3, got %d", a.Len())
	}
}

func TestIter(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	it := a.Iter()
	// consume the iterator
	for i := 0; i < 3; i++ {
		if v, ok := it(); !ok {
			t.Errorf("expected ok, got %v", ok)
		} else if v != i+1 {
			t.Errorf("expected %d, got %d", i+1, v)
		}
	}
	// returns zero value and false, end of iteration
	if v, ok := it(); ok {
		t.Errorf("expected !ok, got %v", ok)
	} else if v != 0 {
		t.Errorf("expected 0, got %d", v)
	}
}

func TestReverseIter(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	it := a.ReverseIter()
	// consume the iterator
	for i := 3; i > 0; i-- {
		if v, ok := it(); !ok {
			t.Errorf("expected ok, got %v", ok)
		} else if v != i {
			t.Errorf("expected %d, got %d", i, v)
		}
	}
	// returns zero value and false, end of iteration
	if v, ok := it(); ok {
		t.Errorf("expected !ok, got %v", ok)
	} else if v != 0 {
		t.Errorf("expected 0, got %d", v)
	}
}
