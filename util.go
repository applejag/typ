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

// CompareFuncFromKey returns a new function that can compare two objects
// by extracting some ordered key from the type.
func CompareFuncFromKey[T any, K Ordered](key func(a T) K) func(a, b T) int {
	return func(a, b T) int {
		return Compare(key(a), key(b))
	}
}

// CompareFuncFromComparable returns a new function that can compare two objects
// by using the type's comparison operation.
func CompareFuncFromComparable[T comparable](less func(a, b T) bool) func(a, b T) int {
	return func(a, b T) int {
		if a == b {
			return 0
		}
		if less(a, b) {
			return -1
		}
		return 1
	}
}

// CompareFuncFromLess returns a new function that can compare two objects
// by flipping the less function.
func CompareFuncFromLess[T any](less func(a, b T) bool) func(a, b T) int {
	return func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if less(b, a) {
			return 1
		}
		return 0
	}
}

// Equal returns true if the first argument is equal to the second.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Zero returns the zero value for a given type.
func Zero[T any]() T {
	var zero T
	return zero
}

// IsZero returns true if the value is zero. If the type implements
//  IsZero() bool
// then that method is also used.
func IsZero[T comparable](value T) bool {
	var zero T
	if value == zero {
		return true
	}
	var asAny any = value
	if isZeroer, ok := asAny.(interface{ IsZero() bool }); ok {
		return isZeroer.IsZero()
	}
	return false
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
