package collection

import "testing"

func TestAny(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	if !a.Iter().Any(func(i int) bool { return i == 1 }) {
		t.Errorf("expected true, got false")
	}
	if a.Iter().Any(func(i int) bool { return i == 4 }) {
		t.Errorf("expected false, got true")
	}
}

func TestCollect(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	b := a.Iter().Collect()
	if b.Len() != 3 {
		t.Errorf("expected 3, got %d", b.Len())
	}
}

func TestCount(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	if a.Iter().Count() != 3 {
		t.Errorf("expected 3, got %d", a.Iter().Count())
	}
}

func TestEvery(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	if !a.Iter().Every(func(i int) bool { return i > 0 }) {
		t.Errorf("expected true, got false")
	}
	if a.Iter().Every(func(i int) bool { return i == 0 }) {
		t.Errorf("expected false, got true")
	}
}

func TestFilter(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	b := a.Iter().Filter(func(i int) bool { return i > 1 }).Collect()
	if b.Len() != 2 {
		t.Errorf("expected 2, got %d", b.Len())
	}
}

func TestFind(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	i, ok := a.Iter().Find(func(i int) bool { return i == 2 })
	if !ok {
		t.Errorf("expected true, got false")
	}
	if i != 2 {
		t.Errorf("expected 2, got %d", i)
	}
	_, ok = a.Iter().Find(func(i int) bool { return i == 4 })
	if ok {
		t.Errorf("expected false, got true")
	}
}

func TestFollowedBy(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	b := a.Iter().FollowedBy(a.Iter()).Collect()
	if b.Len() != 6 {
		t.Errorf("expected 6, got %d", b.Len())
	}
}

func TestForEach(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	ind := -1
	a.Iter().ForEach(func(i int) {
		ind++
		if i != a[ind] {
			t.Errorf("expected %d, got %d", a[ind], i)
		}
	})
	if ind != 2 {
		t.Errorf("expected 2, got %d", ind)
	}
}

func TestSkip(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	b := a.Iter().Skip(1).Collect()
	if b.Len() != 2 {
		t.Errorf("expected 2, got %d", b.Len())
	}
}

func TestSkipWhile(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	b := a.Iter().SkipWhile(func(i int) bool { return i < 2 }).Collect()
	if b.Len() != 2 {
		t.Errorf("expected 2, got %d", b.Len())
	}
}

func TestTake(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	b := a.Iter().Take(2).Collect()
	if b.Len() != 2 {
		t.Errorf("expected 2, got %d", b.Len())
	}
}

func TestTakeWhile(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	b := a.Iter().TakeWhile(func(i int) bool { return i < 3 }).Collect()
	if b.Len() != 2 {
		t.Errorf("expected 2, got %d", b.Len())
	}
}

func TestTap(t *testing.T) {
	var a Vec[int]
	a = a.AppendAll(1, 2, 3)
	ind := -1
	a.Iter().Tap(func(i int) {
		ind++
		if i != a[ind] {
			t.Errorf("expected %d, got %d", a[ind], i)
		}
	}).ForEach(func(i int) {})
	if ind != 2 {
		t.Errorf("expected 2, got %d", ind)
	}
}
