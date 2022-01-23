// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"constraints"
)

// Min returns the smallest value.
func Min[T constraints.Ordered](v ...T) T {
	switch len(v) {
	case 0:
		panic("typ.Min: at least one argument is required")
	case 1:
		return v[0]
	default:
		min := v[0]
		for _, v := range v[1:] {
			if v < min {
				min = v
			}
		}
		return min
	}
}

// Max returns the largest value.
func Max[T constraints.Ordered](v ...T) T {
	switch len(v) {
	case 0:
		panic("typ.Max: at least one argument is required")
	case 1:
		return v[0]
	default:
		max := v[0]
		for _, v := range v[1:] {
			if v > max {
				max = v
			}
		}
		return max
	}
}

// Clamp returns the value clamped between the minimum and maximum values.
func Clamp[T constraints.Ordered](v, min, max T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Clamp01 returns the value clamped between 0 (zero) and 1 (one).
func Clamp01[T constraints.Integer | constraints.Float](v T) T {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

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
