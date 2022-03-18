// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

// Package maps contains utility functions for maps.
package maps

// ContainsValue checks if a value exists inside a map.
func ContainsValue[M ~map[K]V, K comparable, V comparable](m M, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
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
