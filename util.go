// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"constraints"
	"time"
)

// Compare checks if either value is greater or equal to the other.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func Compare[T constraints.Ordered](a, b T) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

// Zero returns the zero value for a given type.
func Zero[T any]() T {
	var zero T
	return zero
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

// SendTimeout sends a value to a channel, or cancels after a given duration.
func SendTimeout[T any](ch chan<- T, value T, timeout time.Duration) bool {
	if timeout <= 0 {
		ch <- value
		return true
	}
	timer := time.NewTimer(timeout)
	select {
	case ch <- value:
		timer.Stop()
		return true
	case <-timer.C:
		return false
	}
}

// RecvTimeout receives a value from a channel, or cancels after a given timeout.
// If the timeout duration is zero or negative, then no limit is used.
func RecvTimeout[T any](ch <-chan T, timeout time.Duration) (T, bool) {
	if timeout <= 0 {
		value, ok := <-ch
		return value, ok
	}
	timer := time.NewTimer(timeout)
	select {
	case value, ok := <-ch:
		timer.Stop()
		return value, ok
	case <-timer.C:
		return Zero[T](), false
	}
}

// Coal will return the first non-zero value. Equivalent to the "null coalescing"
// operator from other languages, or the SQL "COALESCE(...)" expression.
// 	var result = null ?? myDefaultValue;       // C#, JavaScript, PHP, etc
// 	var result = typ.Coal(nil, myDefaultValue) // Go
func Coal[T comparable](values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return zero
}

// Tern returns different values depending on the given conditional boolean.
// Equivalent to the "ternary" operator from other languages.
// 	var result = 1 > 2 ? "yes" : "no";        // C#, JavaScript, PHP, etc
// 	var result = typ.Tern(1 > 2, "yes", "no") // Go
func Tern[T any](cond bool, ifTrue, ifFalse T) T {
	if cond {
		return ifTrue
	}
	return ifFalse
}
