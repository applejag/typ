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

// ZeroOf returns the zero value for a given type. The first argument is unused,
// but using Go's type inference can be useful to create the zero value for an
// anonymous type without needing to declare the full type.
func ZeroOf[T any](T) T {
	var zero T
	return zero
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

// TernCast will cast the value if condition is true. Otherwise the last
// argument is returned.
func TernCast[T any](cond bool, value any, ifFalse T) T {
	if cond {
		return value.(T)
	}
	return ifFalse
}

// IsNil checks if the generic value is nil.
func IsNil[T any](value T) bool {
	var asAny any = value
	return asAny == nil
}

// NewOf returns the result of new() for the same type as the first argument.
// Useful when wanting to call new() on an anonymous type.
func NewOf[T any](*T) *T {
	return new(T)
}

// MakeSliceOf returns the result of make([]T) for the same type as the first
// argument.
//
// Useful when wanting to call make() on a slice of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeSliceOf[T any](_ T, size ...int) []T {
	if len(size) > 2 {
		panic("MakeSliceOf: max 2 size arguments")
	}
	return make([]T, SafeGet(size, 0), SafeGet(size, 1))
}

// MakeSliceOfSlice returns the result of make([]T) for the same type as the
// slice element type of the first argument.
//
// Useful when wanting to call make() on a slice of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeSliceOfSlice[T any](_ []T, size ...int) []T {
	if len(size) > 2 {
		panic("MakeSliceOfSlice: max 2 size arguments")
	}
	return make([]T, SafeGet(size, 0), SafeGet(size, 1))
}

// MakeSliceOfKey returns the result of make([]T) for the same type as the
// map key type of the first argument.
//
// Useful when wanting to call make() on a slice of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeSliceOfKey[K comparable, V any](_ map[K]V, size ...int) []K {
	if len(size) > 2 {
		panic("MakeSliceOfSlice: max 2 size arguments")
	}
	return make([]K, SafeGet(size, 0), SafeGet(size, 1))
}

// MakeSliceOfValue returns the result of make([]T) for the same type as the
// map value type of the first argument.
//
// Useful when wanting to call make() on a slice of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeSliceOfValue[K comparable, V any](_ map[K]V, size ...int) []V {
	if len(size) > 2 {
		panic("MakeSliceOfSlice: max 2 size arguments")
	}
	return make([]V, SafeGet(size, 0), SafeGet(size, 1))
}

// MakeMapOf returns the result of make(map[K]V) for the same type as the first
// arguments.
//
// Useful when wanting to call make() on a map of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeMapOf[K comparable, V any](_ K, _ V, size ...int) map[K]V {
	if len(size) > 1 {
		panic("MakeMapOf: max 1 size argument")
	}
	return make(map[K]V, SafeGet(size, 0))
}

// MakeMapOfMap returns the result of make(map[K]V) for the same type as the
// key and value types of the first argument.
//
// Useful when wanting to call make() on a map of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeMapOfMap[K comparable, V any](_ map[K]V, size ...int) map[K]V {
	if len(size) > 1 {
		panic("MakeMapOf: max 1 size argument")
	}
	return make(map[K]V, SafeGet(size, 0))
}

// MakeChanOf returns the result of make(chan T) for the same type as the first
// argument.
//
// Useful when wanting to call make() on a channel of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeChanOf[T any](_ T, size ...int) chan T {
	if len(size) > 1 {
		panic("MakeChanOf: max 1 size argument")
	}
	return make(chan T, SafeGet(size, 0))
}

// MakeChanOfChan returns the result of make(chan T) for the same type as the
// channel type in the first argument.
//
// Useful when wanting to call make() on a channel of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeChanOfChan[T any](_ chan T, size ...int) chan T {
	if len(size) > 1 {
		panic("MakeChanOf: max 1 size argument")
	}
	return make(chan T, SafeGet(size, 0))
}

// Ptr returns a pointer to the value. Useful when working with literals.
func Ptr[T any](value T) *T {
	return &value
}
