// SPDX-FileCopyrightText: 2025 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sync2

// MapAll returns an iterator over key-value pairs in m.
// The iteeration order is not specified and is not guaranteed to be the same
// from one call to the next.
func MapAll[K comparable, V any](m *Map[K, V]) func(yield func(K, V) bool) {
	return func(yield func(K, V) bool) {
		m.Range(func(key K, value V) bool {
			return yield(key, value)
		})
	}
}

// MapKeys returns an iterator over keys in m.
// The iteeration order is not specified and is not guaranteed to be the same
// from one call to the next.
func MapKeys[K comparable, V any](m *Map[K, V]) func(yield func(K) bool) {
	return func(yield func(K) bool) {
		m.Range(func(key K, _ V) bool {
			return yield(key)
		})
	}
}

// MapValues returns an iterator over values in m.
// The iteeration order is not specified and is not guaranteed to be the same
// from one call to the next.
func MapValues[K comparable, V any](m *Map[K, V]) func(yield func(V) bool) {
	return func(yield func(V) bool) {
		m.Range(func(_ K, value V) bool {
			return yield(value)
		})
	}
}

// MapInsert adds the key-value pairs from seq to m.
// If a key in seq already exists in m, its value will be overwritten.
func MapInsert[K comparable, V any](m *Map[K, V], iter func(yield func(K, V) bool)) {
	iter(func(k K, v V) bool {
		m.Store(k, v)
		return true
	})
}

// MapCollect collects key-value pairs from seq into a new map
// and returns it.
func MapCollect[K comparable, V any](iter func(yield func(K, V) bool)) *Map[K, V] {
	m := new(Map[K, V])
	iter(func(k K, v V) bool {
		m.Store(k, v)
		return true
	})
	return m
}
