package orderedmap

type Cursor[K comparable, V any] struct {
	r        *OrderedMap[K, V]
	position int
}

func (o *Cursor[K, V]) Reset() {
	o.position = 0
}

// Position sets zero based index to specific position
func (o *Cursor[K, V]) Position(index int) {
	o.position = index
}

// Next moves cursor to next value
func (o *Cursor[K, V]) Next() {
	o.position++
}

// Valid checks if current position is valid
func (o *Cursor[K, V]) Valid() bool {
	return o.position < len(o.r.keys)
}

// Value returns index, key and value on current position
func (o *Cursor[K, V]) Value() (index int, key K, val V) {
	return o.position, o.r.keys[o.position], o.r.data[o.r.keys[o.position]]
}
