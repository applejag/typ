// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package maps

// Bimap is a bi-directional map where both the keys and values are indexed
// against each other, allowing performant lookup on both keys and values,
// at the cost of double the memory usage.
type Bimap[K, V comparable] struct {
	forward map[K]V
	reverse map[V]K
}

// Len returns the number of key-value pairs in this map.
func (b *Bimap[K, V]) Len() int {
	if b == nil {
		return 0
	}
	return len(b.forward)
}

// Add another key-value pair to be indexed inside this map. Both the key
// and the value is indexed, to allow performant lookups on both key and value.
//
// On collisions, the old values will be overwritten.
func (b *Bimap[K, V]) Add(key K, value V) {
	if oldVal, ok := b.GetForward(key); ok {
		delete(b.reverse, oldVal)
	}
	if oldKey, ok := b.GetReverse(value); ok {
		delete(b.forward, oldKey)
	}
	if b.forward == nil {
		b.forward = make(map[K]V)
		b.reverse = make(map[V]K)
	}
	b.forward[key] = value
	b.reverse[value] = key
}

// RemoveForward removes a key-value pair from this map based on the key.
func (b *Bimap[K, V]) RemoveForward(key K) {
	if value, ok := b.forward[key]; ok {
		delete(b.reverse, value)
		delete(b.forward, key)
	}
}

// RemoveReverse removes a key-value pair from this map based on the value.
func (b *Bimap[K, V]) RemoveReverse(value V) {
	if key, ok := b.reverse[value]; ok {
		delete(b.reverse, value)
		delete(b.forward, key)
	}
}

// Range loops over all the values in this map. The loop continues as long
// as the function f returns true.
func (b *Bimap[K, V]) Range(f func(key K, value V) bool) {
	for k, v := range b.forward {
		if !f(k, v) {
			return
		}
	}
}

// ContainsForward checks if the given key exists.
func (b *Bimap[K, V]) ContainsForward(key K) bool {
	_, ok := b.forward[key]
	return ok
}

// GetForward performs a lookup on the key to get the value.
func (b *Bimap[K, V]) GetForward(key K) (V, bool) {
	value, ok := b.forward[key]
	return value, ok
}

// ContainsReverse checks if the given value exists.
func (b *Bimap[K, V]) ContainsReverse(value V) bool {
	_, ok := b.reverse[value]
	return ok
}

// GetReverse performs a lookup on the value to get the key.
func (b *Bimap[K, V]) GetReverse(value V) (K, bool) {
	key, ok := b.reverse[value]
	return key, ok
}

// Clear empties this bidirectional map, removing all items.
func (b *Bimap[K, V]) Clear() {
	Clear(b.forward)
	Clear(b.reverse)
}

// Clone creates a shallow copy of this bidirectional map.
func (b *Bimap[K, V]) Clone() Bimap[K, V] {
	return Bimap[K, V]{
		forward: Clone(b.forward),
		reverse: Clone(b.reverse),
	}
}
