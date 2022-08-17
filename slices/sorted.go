// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package slices

import (
	"fmt"
	"sort"

	"gopkg.in/typ.v4"
)

// NewSorted returns a new sorted slice based on a slice of values and a
// custom less function that is used to keep values sorted.
// The values are sorted on insertion.
//
// The less function is expected to return the same value for the same set of
// inputs for the lifetime of the sorted slice.
//
// Note that if the less function is used when finding values to remove. If the
// less function cannot properly distinguish between two elements, then any of
// the equivalent elements may be the one being removed. The SortedSlice does
// not keep track of collision detection.
func NewSorted[S ~[]E, E comparable](values S, less func(a, b E) bool) Sorted[E] {
	slice := makeSortedSlice(values, less)
	return Sorted[E]{slice, typ.CompareFuncFromComparable(less)}
}

// NewSortedOrdered returns a new sorted slice based on a slice of values.
// Only ordered types are allowed. The values are sorted on insertion.
func NewSortedOrdered[S ~[]E, E typ.Ordered](values S) Sorted[E] {
	return NewSorted(values, typ.Less[E])
}

// NewSortedCompare returns a new sorted slice based on a slice of values and
// a custom compare function.
func NewSortedCompare[S ~[]E, E any](values S, compare func(a, b E) int) Sorted[E] {
	slice := makeSortedSlice(values, func(a, b E) bool {
		return compare(a, b) < 0
	})
	return Sorted[E]{slice, compare}
}

// NewSortedFunc returns a new sorted slice based on a slice of values and a
// key extractor function.
func NewSortedFunc[S ~[]E, E any, K typ.Ordered](values S, key func(a E) K) Sorted[E] {
	slice := makeSortedSlice(values, func(a, b E) bool {
		return key(a) < key(b)
	})
	return Sorted[E]{slice, typ.CompareFuncFromKey(key)}
}

func makeSortedSlice[S ~[]E, E any](values S, less func(a, b E) bool) S {
	slice := make([]E, len(values))
	copy(slice, values)
	sort.SliceStable(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
	return slice
}

// Sorted is a slice of ordered values. The slice is always sorted thanks
// to only inserting values in a sorted order.
type Sorted[T any] struct {
	slice   []T
	compare func(a, b T) int
}

func (s Sorted[T]) String() string {
	return fmt.Sprint(s.slice)
}

func (s *Sorted[T]) Get(index int) T {
	if index < 0 || index >= s.Len() {
		panic(fmt.Sprintf("sortedslice: index out of range [%d] with length %d", index, s.Len()))
	}
	return s.slice[index]
}

func (s *Sorted[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.slice)
}

func (s *Sorted[T]) Add(value T) int {
	if s == nil {
		panic("sortedslice: tried to add to nil sortedslice")
	}
	index := s.search(value)
	Insert(&s.slice, index, value)
	return index
}

func (s *Sorted[T]) RemoveAt(index int) {
	if index < 0 || index >= s.Len() {
		panic(fmt.Sprintf("sortedslice: index out of range [%d] with length %d", index, s.Len()))
	}
	Remove(&s.slice, index)
}

func (s *Sorted[T]) Remove(value T) int {
	index := s.Index(value)
	Remove(&s.slice, index)
	return index
}

func (s *Sorted[T]) Contains(value T) bool {
	return s.Index(value) != -1
}

func (s *Sorted[T]) Index(value T) int {
	index := s.search(value)
	if index < 0 || index >= s.Len() || s.compare(s.slice[index], value) != 0 {
		return -1
	}
	return index
}

func (s *Sorted[T]) search(value T) int {
	if s.compare == nil {
		panic("sortedslice: not initialized")
	}
	return sort.Search(len(s.slice), func(i int) bool {
		return s.compare(s.slice[i], value) >= 0
	})
}
