package orderedmap_test

import (
	"fmt"

	orderedmap "github.com/oktalz/ordered-map"
)

func Example() {
	m := orderedmap.New[string, customType]()

	// add 3 values to map
	m.Set("1", customType{1, "1"}) // map[1:{1 1}]
	m.Set("2", customType{2, "2"}) // map[1:{1 1} 2:{2 2}]
	m.Set("3", customType{3, "3"}) // map[1:{1 1} 2:{2 2} 3:{3 3}]

	// delete one with key "2"
	m.Delete("2") // map[1:{1 1} 3:{3 3}]

	// insert to specific index
	m.Insert(1, "42", customType{42, "42"}) // map[1:{1 1} 3:{3 3} 42:{42 42}] order:[1 42 3]

	c := m.Cursor()
	for c.Reset(); c.Valid(); c.Next() { // c.Reset() is not needed on first run
		index, key, value := c.Value()
		fmt.Println(index, key, value)
	}

	// get by key
	v, ok := m.Get("3")
	if ok {
		fmt.Println(v)
	}

	// get by index
	v, err := m.GetByIndex(1)
	if err != nil {
		fmt.Println(v)
	}

	// get whole map
	fmt.Println(m.Map())

	// get in ordered way
	fmt.Println(m.Values())

	// print keys
	fmt.Println(m.Keys())
}

// This is a package-level example:
func ExampleCursor() {
	m := orderedmap.New[string, customType]()

	m.Set("1", customType{1, "1"})
	m.Set("2", customType{2, "2"})
	m.Set("3", customType{3, "3"})

	c := m.Cursor()
	for c.Reset(); c.Valid(); c.Next() {
		index, key, value := c.Value()
		fmt.Println(index, key, value)
	}
}
