package orderedmap_test

import (
	"testing"

	orderedmap "github.com/oktalz/ordered-map"
)

func TestIntString(t *testing.T) {
	m := orderedmap.New[int, string]()

	m.Set(1, "1")
	m.Set(2, "2")
	m.Set(3, "3")
	m.Delete(2)

	size := m.Len()
	if size != 2 {
		t.Fatalf("size should be 2 but its %d", size)
	}

	err := m.Insert(100, 42, "42")
	if err == nil {
		t.Fatal("error is nil, but index is to big")
	}

	err = m.Insert(1, 42, "42")
	if err != nil {
		t.Fatalf("error is not nil, but %v", err)
	}

	err = m.Insert(1, 42, "42")
	if err == nil {
		t.Fatal("error is not thrown")
	}

	c := m.Cursor()

	counter := 0
	for c.Reset(); c.Valid(); c.Next() {
		counter++
	}

	if counter != 3 {
		t.Fatalf("counter should be 3 and it it %d", counter)
	}

	c.Position(1)
	if !c.Valid() {
		t.Fatal("cursor position should be valid")
	}

	index, key, value := c.Value()
	if index != 1 {
		t.Fatalf("cursor index should be 2 but it is %d", index)
	}
	if key != 42 {
		t.Fatalf("key should be `42` but it is %d", key)
	}
	if value != "42" {
		t.Fatalf("key should be `42` but it is %s", value)
	}

	sl := m.Values()
	if len(sl) != 3 {
		t.Fatalf("len should be 3 and it it %d", len(sl))
	}

	v := sl[1]
	if v != "42" {
		t.Fatalf("value should be `42` but its [%s]", v)
	}

	mp := m.Map()
	v, ok := mp[42]
	if !ok {
		t.Fatalf("value not found")
	}
	if v != "42" {
		t.Fatalf("value should be `42` but its [%s]", v)
	}

	err = m.DeleteAtIndex(100)
	if err == nil {
		t.Fatal("error is nil, but index does not exist")
	}

	err = m.DeleteAtIndex(0)
	if err != nil {
		t.Fatalf("error is not nil, but %v", err)
	}

	v, ok = m.Get(3)
	if !ok {
		t.Fatal("key 3 does not exist")
	}
	if v != "3" {
		t.Fatalf("value should be `3` but its [%s]", v)
	}

	_, err = m.GetByIndex(1000)
	if err == nil {
		t.Fatal("error not thrown")
	}

	v, err = m.GetByIndex(1)
	if err != nil {
		t.Fatalf("error is not nil, but %v", err)
	}
	if v != "3" {
		t.Fatalf("value should be `3` but its [%s]", v)
	}

	_, err = m.Index(3000)
	if err == nil {
		t.Fatal("error not thrown")
	}

	index, err = m.Index(3)
	if err != nil {
		t.Fatalf("error is not nil, but %v", err)
	}
	if index != 1 {
		t.Fatalf("value should be 1 but its [%d]", index)
	}

	keys := m.Keys()
	if len(keys) != 2 {
		t.Fatalf("len(keys) != 3 len=%d", len(keys))
	}

	if keys[0] != 42 || keys[1] != 3 {
		t.Fatalf("keys not correct %v", keys)
	}
}
