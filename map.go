package orderedmap

type OrderedMap[K comparable, V any] struct {
	data map[K]V
	keys []K
}

// New[K comparable, V any] create a map[K]V that keeps order of items that are added
func New[K comparable, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{
		data: map[K]V{},
		keys: []K{},
	}
}

// Get returns value for a key
func (o *OrderedMap[K, V]) Len() int {
	return len(o.keys)
}

// Get returns value for a key
func (o *OrderedMap[K, V]) Get(key K) (V, bool) {
	val, exists := o.data[key]
	return val, exists
}

// Index returns index in ordered map of a key
func (o *OrderedMap[K, V]) Index(key K) (int, error) {
	for i, val := range o.keys {
		if val == key {
			return i, nil
		}
	}
	return 0, ErrKeyDoesNotExist
}

// Get returns value V on specific index
func (o *OrderedMap[K, V]) GetByIndex(index int) (V, error) {
	var result V
	if index < 0 || index > (len(o.data)-1) {
		return result, ErrIndexOutOfRange
	}

	key := o.keys[index]

	result, _ = o.data[key]
	return result, nil
}

// Set adds a key/value information to map.
//
// This adds key as the last one in list.
// If key already exist it overrides value
func (o *OrderedMap[K, V]) Set(key K, val V) {
	if _, exists := o.data[key]; !exists {
		o.keys = append(o.keys, key)
	}
	o.data[key] = val
}

// Insert can use index to inser specific key on other index than last
func (o *OrderedMap[K, V]) Insert(index int, key K, val V) error {
	_, exists := o.data[key]
	if exists {
		return ErrKeyAlreadyExist
	}

	if index > len(o.data) {
		return ErrIndexOutOfRange
	}
	var newValue K
	o.keys = append(o.keys, newValue)
	copy(o.keys[index+1:], o.keys[index:])

	o.keys[index] = key
	o.data[key] = val
	return nil
}

// Delete is same as regular delete on map, it does not return error if key does not exist
func (o *OrderedMap[K, V]) Delete(key K) {
	delete(o.data, key)

	index := -1

	for i, val := range o.keys {
		if val == key {
			index = i
			break
		}
	}
	if index != -1 {
		o.keys = append(o.keys[:index], o.keys[index+1:]...)
	}
}

// DeleteAtIndex deletes item on specific key index
func (o *OrderedMap[K, V]) DeleteAtIndex(index int) error {
	if index < 0 || index > (len(o.data)-1) {
		return ErrIndexOutOfRange
	}

	key := o.keys[index]
	delete(o.data, key)

	o.keys = append(o.keys[:index], o.keys[index+1:]...)
	return nil
}

// Map returns regular Go map of values
func (o *OrderedMap[K, V]) Map() map[K]V {
	return o.data
}

// Slice returns slice of values. Order of values is order of insertion
func (o *OrderedMap[K, V]) Slice() []V {
	c := o.Cursor()
	result := []V{}
	for c.Reset(); c.Valid(); c.Next() {
		_, _, val := c.Value()
		result = append(result, val)
	}
	return result
}

// Keys return keys of a map in slice (in same order as added)
func (o *OrderedMap[K, V]) Keys() []K {
	return o.keys
}

// Cursor return Cursor that can be used to iterate over values
//
//  c := m.Cursor()
//  for c.Reset(); c.Valid(); c.Next() {
//    index, key, value := c.Value()
//    ...
//  }
func (o *OrderedMap[K, V]) Cursor() Cursor[K, V] {
	cur := Cursor[K, V]{
		r: o,
	}
	return cur
}
