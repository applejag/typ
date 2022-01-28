// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

// Index returns the index of a value, or -1 if none found.
//
// This differs from Search as Index doesn't require the slice to be sorted.
func Index[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// IndexFunc returns the index of the first occurence where the function returns
// true, or -1 if none found.
//
// This differs from Search as Index doesn't require the slice to be sorted.
func IndexFunc[T any](slice []T, f func(value T) bool) int {
	for i, v := range slice {
		if f(v) {
			return i
		}
	}
	return -1
}

// Repeat creates a new slice with the given value repeated across it.
func Repeat[T any](value T, count int) []T {
	result := make([]T, count)
	for i := 0; i < count; i++ {
		result[i] = value
	}
	return result
}

// Trim returns a slice of the slice that has had all unwanted values trimmed
// away from both the start and the end.
func Trim[T comparable](slice []T, unwanted []T) []T {
	return TrimLeft(TrimRight(slice, unwanted), unwanted)
}

// TrimFunc returns a slice of the slice that has had all unwanted values
// trimmed away from both the start and the end.
// Values are considered unwanted if the callback returns false.
func TrimFunc[T any](slice []T, unwanted func(value T) bool) []T {
	return TrimLeftFunc(TrimRightFunc(slice, unwanted), unwanted)
}

// TrimLeft returns a slice of the slice that has had all unwanted values
// trimmed away from the start of the slice.
func TrimLeft[T comparable](slice []T, unwanted []T) []T {
	for len(slice) > 0 && Contains(unwanted, slice[0]) {
		slice = slice[1:]
	}
	return slice
}

// TrimLeftFunc returns a slice of the slice that has had all unwanted values
// trimmed away from the start of the slice.
// Values are considered unwanted if the callback returns false.
func TrimLeftFunc[T any](slice []T, unwanted func(value T) bool) []T {
	for len(slice) > 0 && unwanted(slice[0]) {
		slice = slice[1:]
	}
	return slice
}

// TrimRight returns a slice of the slice that has had all unwanted values
// trimmed away from the end of the slice.
func TrimRight[T comparable](slice []T, unwanted []T) []T {
	for len(slice) > 0 && Contains(unwanted, slice[len(slice)-1]) {
		slice = slice[:len(slice)-1]
	}
	return slice
}

// TrimRightFunc returns a slice of the slice that has had all unwanted values
// trimmed away from the end of the slice.
// Values are considered unwanted if the callback returns false.
func TrimRightFunc[T any](slice []T, unwanted func(value T) bool) []T {
	for len(slice) > 0 && unwanted(slice[len(slice)-1]) {
		slice = slice[:len(slice)-1]
	}
	return slice
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

// DistinctFunc returns a new slice of only unique values.
func DistinctFunc[T any](slice []T, equals func(a, b T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !ContainsFunc(result, v, equals) {
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

// ContainsFunc checks if a value exists inside a slice of values with a custom
// equals operation.
func ContainsFunc[T any](slice []T, value T, equals func(a, b T) bool) bool {
	for _, v := range slice {
		if equals(v, value) {
			return true
		}
	}
	return false
}

// TryGet will get a value from a slice, or return false on the second return
// value if the index is outside the bounds of the slice. Passing a nil slice is
// equivalent to passing an empty slice.
func TryGet[T any](slice []T, index int) (T, bool) {
	if index < 0 || index >= len(slice) {
		return Zero[T](), false
	}
	return slice[index], true
}

// SafeGet will get a value from a slice, or the zero value for the type if
// the index is outside the bounds of the slice. Passing a nil slice is
// equivalent to passing an empty slice.
func SafeGet[T any](slice []T, index int) T {
	if index < 0 || index >= len(slice) {
		return Zero[T]()
	}
	return slice[index]
}

// SafeGetOr will get a value from a slice, or the fallback value for the type
// if the index is outside the bounds of the slice. Passing a nil slice is
// equivalent to passing an empty slice.
func SafeGetOr[T any](slice []T, index int, fallback T) T {
	if index < 0 || index >= len(slice) {
		return fallback
	}
	return slice[index]
}

// Last returns the last item in a slice, or the zero value if the slice is
// empty.
func Last[T any](slice []T) T {
	if len(slice) == 0 {
		return Zero[T]()
	}
	return slice[len(slice)-1]
}

// Any checks if any value matches the condition. Returns false if the slice is
// empty.
func Any[T any](slice []T, cond func(value T) bool) bool {
	for _, v := range slice {
		if cond(v) {
			return true
		}
	}
	return false
}

// All checks if all values matches the condition. Returns true if the slice is
// empty.
func All[T any](slice []T, cond func(value T) bool) bool {
	for _, v := range slice {
		if !cond(v) {
			return false
		}
	}
	return true
}

// Map will apply a conversion function to all elements in a slice and return
// the new slice with converted values.
func Map[TA any, TB any](slice []TA, conv func(value TA) TB) []TB {
	result := make([]TB, len(slice))
	for i, v := range slice {
		result[i] = conv(v)
	}
	return result
}

// MapErr will apply a conversion function to all elements in a slice and return
// the new slice with converted values. Will cancel the conversion on the first
// error occurrence.
func MapErr[TA any, TB any](slice []TA, conv func(value TA) (TB, error)) ([]TB, error) {
	result := make([]TB, len(slice))
	var err error
	for i, v := range slice {
		result[i], err = conv(v)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// Filter will return a new slice of all matching elements.
func Filter[T any](slice []T, match func(value T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if match(v) {
			result = append(result, v)
		}
	}
	return result
}

// Fold will accumulate an answer based on all values in a slice. Returns the
// seed value as-is if the slice is empty.
func Fold[TState, T any](slice []T, seed TState, acc func(state TState, value T) TState) TState {
	state := seed
	for _, v := range slice {
		seed = acc(state, v)
	}
	return state
}

// FoldReverse will accumulate an answer based on all values in a slice,
// starting with the last element and accumulating backwards. Returns the
// seed value as-is if the slice is empty.
func FoldReverse[TState, T any](slice []T, seed TState, acc func(state TState, value T) TState) TState {
	state := seed
	for i := len(slice) - 1; i >= 0; i++ {
		seed = acc(state, slice[i])
	}
	return state
}

// Concat returns a new slice with the values from the two slices concatenated.
func Concat[T any](a, b []T) []T {
	result := make([]T, len(a)+len(b))
	copy(result[:len(a)], a)
	copy(result[len(a):], b)
	return result
}