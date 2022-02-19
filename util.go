// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

// Compare checks if either value is greater or equal to the other.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
func Compare[T Ordered](a, b T) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

// Less returns true if the first argument is less than the second.
func Less[T Ordered](a, b T) bool {
	return a < b
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
func MakeSliceOf[E any](_ E, size ...int) []E {
	if len(size) > 2 {
		panic("MakeSliceOf: max 2 size arguments")
	}
	return make([]E, SafeGet(size, 0), SafeGet(size, 1))
}

// MakeSliceOfSlice returns the result of make([]E) for the same type as the
// slice element type of the first argument.
//
// Useful when wanting to call make() on a slice of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeSliceOfSlice[S ~[]E, E any](_ S, size ...int) S {
	if len(size) > 2 {
		panic("MakeSliceOfSlice: max 2 size arguments")
	}
	return make([]E, SafeGet(size, 0), SafeGet(size, 1))
}

// MakeSliceOfKey returns the result of make([]T) for the same type as the
// map key type of the first argument.
//
// Useful when wanting to call make() on a slice of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeSliceOfKey[M ~map[K]V, K comparable, V any](_ M, size ...int) []K {
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
func MakeSliceOfValue[M ~map[K]V, K comparable, V any](_ M, size ...int) []V {
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
func MakeMapOfMap[M ~map[K]V, K comparable, V any](_ M, size ...int) M {
	if len(size) > 1 {
		panic("MakeMapOf: max 1 size argument")
	}
	return make(M, SafeGet(size, 0))
}

// MakeChanOf returns the result of make(chan V) for the same type as the first
// argument.
//
// Useful when wanting to call make() on a channel of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeChanOf[V any](_ V, size ...int) chan V {
	if len(size) > 1 {
		panic("MakeChanOf: max 1 size argument")
	}
	return make(chan V, SafeGet(size, 0))
}

// MakeChanOfChan returns the result of make(chan V) for the same type as the
// channel type in the first argument.
//
// Useful when wanting to call make() on a channel of anonymous types by making
// use of type inference to skip having to declare the full anonymous type
// multiple times, which is quite common when writing tests.
func MakeChanOfChan[C ~chan V, V any](_ C, size ...int) C {
	if len(size) > 1 {
		panic("MakeChanOf: max 1 size argument")
	}
	return make(C, SafeGet(size, 0))
}

// Ref returns a pointer to the value. Useful when working with literals.
func Ref[T any](value T) *T {
	return &value
}

// DerefZero returns the dereferenced value, or zero if it was nil.
func DerefZero[P ~*V, V any](ptr P) V {
	if ptr == nil {
		return Zero[V]()
	}
	return *ptr
}
