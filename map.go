// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

// ContainsValue checks if a value exists inside a map.
func ContainsValue[K comparable, V comparable](m map[K]V, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// CloneMap returns a shallow copy of a map.
func CloneMap[K comparable, V any](m map[K]V) map[K]V {
	newMap := make(map[K]V, len(m))
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}
