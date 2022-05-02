// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

// Package maps contains utility functions for maps.
package maps

import "gopkg.in/typ.v4"

// ContainsValue checks if a value exists inside a map. It searches
// the map linearly, meaning it's slow, and a different map structure may be
// a wiser choice.
func ContainsValue[M ~map[K]V, K comparable, V comparable](m M, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// KeyOf returns the key of a value, or false if none is found. It searches
// the map linearly, meaning it's slow, and a different map structure may be
// a wiser choice.
func KeyOf[M ~map[K]V, K comparable, V comparable](m M, value V) (K, bool) {
	for k, v := range m {
		if v == value {
			return k, true
		}
	}
	return typ.Zero[K](), false
}

// Clone returns a shallow copy of a map.
func Clone[M ~map[K]V, K comparable, V any](m M) M {
	newMap := make(M, len(m))
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

// Clear will delete all key-value pairs from a map, rendering it empty.
func Clear[M ~map[K]V, K comparable, V any](m M) {
	// Relies on the compiler optimization introduced in Go v1.11
	// https://go.dev/doc/go1.11#performance-compiler
	for k := range m {
		delete(m, k)
	}
}

// HasKey returns true if the given map has a value on the given key.
func HasKey[M ~map[K]V, K comparable, V any](m M, key K) bool {
	_, ok := m[key]
	return ok
}

// Keys returns a slice of all the keys in this map. The order of the
// keys is arbitrary.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of all the values in this map. The order of the
// values is arbitrary.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

