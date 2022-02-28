package orderedmap_test

import (
	"testing"

	orderedmap "github.com/oktalz/ordered-map"
)

type customType struct {
	A int
	B string
}

func TestStringStruct(t *testing.T) {
	m := orderedmap.New[string, customType]()

	m.Set("1", customType{1, "1"})
	m.Set("2", customType{2, "2"})
	m.Set("3", customType{3, "3"})
	m.Delete("2")

	m.Insert(1, "42", customType{42, "42"})

	c := m.Cursor()

	counter := 0
	for c.Reset(); c.Valid(); c.Next() {
		counter++
	}

	if counter != 3 {
		t.Fatalf("counter should be 3 and it it %d", counter)
	}

	sl := m.Slice()
	if len(sl) != 3 {
		t.Fatalf("len should be 3 and it it %d", len(sl))
	}

	v := sl[1]
	if v.B != "42" {
		t.Fatalf("value should be `42` but its [%s]", v.B)
	}

	mp := m.Map()
	v, ok := mp["42"]
	if !ok {
		t.Fatalf("value not found")
	}
	if v.B != "42" {
		t.Fatalf("value should be `42` but its [%s]", v.B)
	}

	v, ok = m.Get("3")
	if !ok {
		t.Fatal("key 3 does not exist")
	}
	if v.B != "3" {
		t.Fatalf("value should be `3` but its [%s]", v.B)
	}
}
