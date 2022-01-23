// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"constraints"
)

// compare checks if either value is greater or equal to the other.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func compare[T constraints.Ordered](a, b T) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

// Distinct returns a new slice of only unique values.
func Distinct[T comparable](slice []T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// Contains checks if a value exists inside a slice of values.
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsValue checks if a value exists inside a map.
func ContainsValue[K comparable, V comparable](m map[K]V, value V) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}
