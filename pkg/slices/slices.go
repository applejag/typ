// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

// Package slices contains utility functions for Go slices, as well as a Sorted slice
// type.
package slices

import (
	"gopkg.in/typ.v3"
	"gopkg.in/typ.v3/pkg/maps"
	"gopkg.in/typ.v3/pkg/sets"
)

// Fill populates a whole slice with the same value using exponential copy.
func Fill[S ~[]E, E any](slice S, value E) {
	if len(slice) == 0 {
		return
	}
	slice[0] = value
	for i := 1; i < len(slice); i += i {
		copy(slice[i:], slice[:i])
	}
}

// Insert inserts a value at a given index in the slice and shifts all following
// values to the right.
func Insert[S ~[]E, E any](slice *S, index int, value E) {
	*slice = append(*slice, value)
	copy((*slice)[index+1:], (*slice)[index:])
	(*slice)[index] = value
}

// InsertSlice inserts a slice of values at a given index in the slice and
// shifts all following values to the right.
func InsertSlice[S ~[]E, E any](slice *S, index int, values S) {
	*slice = append(*slice, values...)
	copy((*slice)[index+len(values):], (*slice)[index:])
	copy((*slice)[index:], values)
}

// Remove takes out a value at a given index and shifts all following values
// to the left.
func Remove[S ~[]E, E any](slice *S, index int) {
	copy((*slice)[index:], (*slice)[index+1:])
	*slice = (*slice)[:len(*slice)-1]
}

// RemoveSlice takes out a slice of values at a given index and length and
// shifts all following values to the left.
func RemoveSlice[S ~[]E, E any](slice *S, index int, length int) {
	copy((*slice)[index:], (*slice)[index+length:])
	*slice = (*slice)[:len(*slice)-length]
}

// Index returns the index of a value, or -1 if none found.
//
// This differs from Search as Index doesn't require the slice to be sorted.
func Index[S ~[]E, E comparable](slice S, value E) int {
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
func IndexFunc[S ~[]E, E any](slice S, f func(value E) bool) int {
	for i, v := range slice {
		if f(v) {
			return i
		}
	}
	return -1
}

// Repeat creates a new slice with the given value repeated across it.
func Repeat[E any](value E, count int) []E {
	result := make([]E, count)
	Fill(result, value)
	return result
}

// Trim returns a slice of the slice that has had all unwanted values trimmed
// away from both the start and the end.
func Trim[S ~[]E, E comparable](slice S, unwanted S) S {
	return TrimLeft[S, E](TrimRight[S, E](slice, unwanted), unwanted)
}

// TrimFunc returns a slice of the slice that has had all unwanted values
// trimmed away from both the start and the end.
// Values are considered unwanted if the callback returns false.
func TrimFunc[S ~[]E, E any](slice S, unwanted func(value E) bool) S {
	return TrimLeftFunc(TrimRightFunc(slice, unwanted), unwanted)
}

// TrimLeft returns a slice of the slice that has had all unwanted values
// trimmed away from the start of the slice.
func TrimLeft[S ~[]E, E comparable](slice S, unwanted S) S {
	for len(slice) > 0 && Contains(unwanted, slice[0]) {
		slice = slice[1:]
	}
	return slice
}

// TrimLeftFunc returns a slice of the slice that has had all unwanted values
// trimmed away from the start of the slice.
// Values are considered unwanted if the callback returns false.
func TrimLeftFunc[S ~[]E, E any](slice S, unwanted func(value E) bool) S {
	for len(slice) > 0 && unwanted(slice[0]) {
		slice = slice[1:]
	}
	return slice
}

// TrimRight returns a slice of the slice that has had all unwanted values
// trimmed away from the end of the slice.
func TrimRight[S ~[]E, E comparable](slice S, unwanted S) S {
	for len(slice) > 0 && Contains(unwanted, slice[len(slice)-1]) {
		slice = slice[:len(slice)-1]
	}
	return slice
}

// TrimRightFunc returns a slice of the slice that has had all unwanted values
// trimmed away from the end of the slice.
// Values are considered unwanted if the callback returns false.
func TrimRightFunc[S ~[]E, E any](slice S, unwanted func(value E) bool) S {
	for len(slice) > 0 && unwanted(slice[len(slice)-1]) {
		slice = slice[:len(slice)-1]
	}
	return slice
}

// Distinct returns a new slice of only unique values.
func Distinct[S ~[]E, E comparable](slice S) S {
	result := make(S, 0, len(slice))
	for _, v := range slice {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// DistinctFunc returns a new slice of only unique values.
func DistinctFunc[S ~[]E, E any](slice S, equals func(a, b E) bool) S {
	result := make(S, 0, len(slice))
	for _, v := range slice {
		if !ContainsFunc(result, v, equals) {
			result = append(result, v)
		}
	}
	return result
}

// Contains checks if a value exists inside a slice of values.
func Contains[S ~[]E, E comparable](slice S, value E) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsFunc checks if a value exists inside a slice of values with a custom
// equals operation.
func ContainsFunc[S ~[]E, E any](slice S, value E, equals func(a, b E) bool) bool {
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
func TryGet[S ~[]E, E any](slice S, index int) (E, bool) {
	if index < 0 || index >= len(slice) {
		return typ.Zero[E](), false
	}
	return slice[index], true
}

// SafeGet will get a value from a slice, or the zero value for the type if
// the index is outside the bounds of the slice. Passing a nil slice is
// equivalent to passing an empty slice.
func SafeGet[S ~[]E, E any](slice S, index int) E {
	if index < 0 || index >= len(slice) {
		return typ.Zero[E]()
	}
	return slice[index]
}

// SafeGetOr will get a value from a slice, or the fallback value for the type
// if the index is outside the bounds of the slice. Passing a nil slice is
// equivalent to passing an empty slice.
func SafeGetOr[S ~[]E, E any](slice S, index int, fallback E) E {
	if index < 0 || index >= len(slice) {
		return fallback
	}
	return slice[index]
}

// Any checks if any value matches the condition. Returns false if the slice is
// empty.
func Any[S ~[]E, E any](slice S, cond func(value E) bool) bool {
	for _, v := range slice {
		if cond(v) {
			return true
		}
	}
	return false
}

// All checks if all values matches the condition. Returns true if the slice is
// empty.
func All[S ~[]E, E any](slice S, cond func(value E) bool) bool {
	for _, v := range slice {
		if !cond(v) {
			return false
		}
	}
	return true
}

// Map will apply a conversion function to all elements in a slice and return
// the new slice with converted values.
func Map[S ~[]E, E, Result any](slice S, conv func(value E) Result) []Result {
	result := make([]Result, len(slice))
	for i, v := range slice {
		result[i] = conv(v)
	}
	return result
}

// MapErr will apply a conversion function to all elements in a slice and return
// the new slice with converted values. Will cancel the conversion on the first
// error occurrence.
func MapErr[S ~[]E, E, Result any](slice S, conv func(value E) (Result, error)) ([]Result, error) {
	result := make([]Result, len(slice))
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
func Filter[S ~[]E, E any](slice S, match func(value E) bool) S {
	result := make(S, 0, len(slice))
	for _, v := range slice {
		if match(v) {
			result = append(result, v)
		}
	}
	return result
}

// Fold will accumulate an answer based on all values in a slice. Returns the
// seed value as-is if the slice is empty.
func Fold[S ~[]E, State, E any](slice S, seed State, acc func(state State, value E) State) State {
	state := seed
	for _, v := range slice {
		seed = acc(state, v)
	}
	return state
}

// FoldReverse will accumulate an answer based on all values in a slice,
// starting with the last element and accumulating backwards. Returns the
// seed value as-is if the slice is empty.
func FoldReverse[S ~[]E, State, E any](slice S, seed State, acc func(state State, value E) State) State {
	state := seed
	for i := len(slice) - 1; i >= 0; i++ {
		seed = acc(state, slice[i])
	}
	return state
}

// Concat returns a new slice with the values from the two slices concatenated.
func Concat[S ~[]E, E any](a, b S) S {
	result := make(S, len(a)+len(b))
	copy(result[:len(a)], a)
	copy(result[len(a):], b)
	return result
}

// Grouping is a key-values store returned by the GroupBy functions.
type Grouping[K, V any] struct {
	Key    K
	Values []V
}

// GroupBy will group all elements in the slice and return a slice of groups,
// using the key from the function provided.
func GroupBy[S ~[]V, K comparable, V any](slice S, keyer func(value V) K) []Grouping[K, V] {
	m := map[K][]V{}
	var orderedKeys []K
	for _, v := range slice {
		key := keyer(v)
		values, ok := m[key]
		m[key] = append(values, v)
		if !ok {
			orderedKeys = append(orderedKeys, key)
		}
	}
	groups := make([]Grouping[K, V], len(orderedKeys))
	for i, key := range orderedKeys {
		groups[i] = Grouping[K, V]{
			Key:    key,
			Values: m[key],
		}
	}
	return groups
}

// Counting is a key-count store returned by the CountBy function.
type Counting[K any] struct {
	Key   K
	Count int
}

// CountBy will count the number of occurrences for each key, using the key
// from the function provided.
func CountBy[S ~[]V, K comparable, V any](slice S, keyer func(value V) K) []Counting[K] {
	m := map[K]int{}
	var orderedKeys []K
	for _, v := range slice {
		key := keyer(v)
		count, ok := m[key]
		m[key] = count + 1
		if !ok {
			orderedKeys = append(orderedKeys, key)
		}
	}
	groups := make([]Counting[K], len(orderedKeys))
	for i, key := range orderedKeys {
		groups[i] = Counting[K]{
			Key:   key,
			Count: m[key],
		}
	}
	return groups
}

// Pairs returns a slice of pairs for the given slice. If the slice has less
// than two items, then an empty slice is returned.
func Pairs[S ~[]E, E any](slice S) [][2]E {
	if len(slice) < 2 {
		return nil
	}
	lim := len(slice) - 1
	pairs := make([][2]E, lim)
	for i := 0; i < lim; i++ {
		pairs[i] = [2]E{slice[i], slice[i+1]}
	}
	return pairs
}

// PairsFunc invokes the provided callback for all pairs for the given slice.
// If the slice has less than two items, then no invokation is performed.
func PairsFunc[S ~[]E, E any](slice S, callback func(a, b E)) {
	if len(slice) < 2 {
		return
	}
	lim := len(slice) - 1
	for i := 0; i < lim; i++ {
		callback(slice[i], slice[i+1])
	}
}

// Windowed returns a slice of windows, where each window is a slice of the
// specified size from the specified slice.
func Windowed[S ~[]E, E any](slice S, size int) []S {
	if len(slice) < size {
		return nil
	}
	lim := len(slice) - size + 1
	windows := make([]S, lim)
	for i := 0; i < lim; i++ {
		windows[i] = slice[i : i+size]
	}
	return windows
}

// WindowedFunc invokes the provided callback for all windows, where each window
// is a slice of the specified size from the specified slice.
func WindowedFunc[S ~[]E, E any](slice S, size int, callback func(window S)) {
	if len(slice) < size {
		return
	}
	lim := len(slice) - size + 1
	for i := 0; i < lim; i++ {
		callback(slice[i : i+size])
	}
}

// Chunk divides the slice up into chunks with a size limit. The last chunk
// may be smaller than size if the slice is not evenly divisible.
func Chunk[S ~[]E, E any](slice S, size int) []S {
	if len(slice) == 0 {
		return nil
	}
	div := len(slice) / size
	rounded := div * size
	lim := div + (len(slice) - rounded)
	chunks := make([]S, lim)
	for i, j := 0, 0; j < rounded; i, j = i+1, j+size {
		chunks[i] = slice[j : j+size]
	}
	if div != lim {
		chunks[lim-1] = slice[rounded:]
	}
	return chunks
}

// ChunkFunc divides the slice up into chunks and invokes the callback on each
// chunk. The last chunk may be smaller than size if the slice is not evenly
// divisible.
func ChunkFunc[S ~[]E, E any](slice S, size int, callback func(chunk S)) {
	if len(slice) == 0 {
		return
	}
	div := len(slice) / size
	rounded := div * size
	for i, j := 0, 0; j < rounded; i, j = i+1, j+size {
		callback(slice[j : j+size])
	}
	if rounded != len(slice) {
		callback(slice[rounded:])
	}
}

// Except returns a new slice for all items that are not found in the slice of
// items to exclude.
func Except[S ~[]E, E comparable](slice S, exclude S) S {
	set := maps.NewSetFromSlice(exclude)
	return ExceptSet(slice, set)
}

// ExceptSet returns a new slice for all items that are not found in the set of
// items to exclude.
func ExceptSet[S ~[]E, E comparable](slice S, exclude sets.Set[E]) S {
	result := make(S, 0, len(slice))
	for _, v := range slice {
		if !exclude.Has(v) {
			result = append(result, v)
		}
	}
	return result
}

// Last returns the last item in a slice. Will panic with an out of bound error
// if the slice is empty.
func Last[S ~[]E, E any](slice S) E {
	return slice[len(slice)-1]
}

// Clone returns a shallow copy of a slice.
func Clone[S ~[]E, E any](slice S) S {
	newSlice := make(S, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// Grow will add n number of values to the end of the slice.
func Grow[S ~[]E, E any](slice S, n int) S {
	// Relies on the compiler optimization introduced in Go v1.11
	// https://go.dev/doc/go1.11#performance-compiler
	return append(slice, make(S, n)...)
}
