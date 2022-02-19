// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"fmt"
	"sort"
)

// NewSortedSlice returns a new sorted slice based on a slice of values and a
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
func NewSortedSlice[S ~[]E, E comparable](values S, less func(a, b E) bool) SortedSlice[E] {
	slice := make([]E, len(values))
	copy(slice, values)
	sort.SliceStable(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
	return SortedSlice[E]{slice, less}
}

// NewSortedSliceOrdered returns a new sorted slice based on a slice of values.
// Only ordered types are allowed. The values are sorted on insertion.
func NewSortedSliceOrdered[T Ordered](values ...T) SortedSlice[T] {
	return NewSortedSlice(values, Less[T])
}

// SortedSlice is a slice of ordered values. The slice is always sorted thanks
// to only inserting values in a sorted order.
type SortedSlice[T comparable] struct {
	slice []T
	less  func(a, b T) bool
}

func (s SortedSlice[T]) String() string {
	return fmt.Sprint(s.slice)
}

func (s *SortedSlice[T]) Get(index int) T {
	if index < 0 || index >= s.Len() {
		panic(fmt.Sprintf("sortedslice: index out of range [%d] with length %d", index, s.Len()))
	}
	return s.slice[index]
}

func (s *SortedSlice[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.slice)
}

func (s *SortedSlice[T]) Add(value T) int {
	if s == nil {
		panic("sortedslice: tried to add to nil sortedslice")
	}
	index := s.search(value)
	Insert(&s.slice, index, value)
	return index
}

func (s *SortedSlice[T]) RemoveAt(index int) {
	if index < 0 || index >= s.Len() {
		panic(fmt.Sprintf("sortedslice: index out of range [%d] with length %d", index, s.Len()))
	}
	Remove(&s.slice, index)
}

func (s *SortedSlice[T]) Remove(value T) int {
	index := s.Index(value)
	Remove(&s.slice, index)
	return index
}

func (s *SortedSlice[T]) Contains(value T) bool {
	return s.Index(value) != -1
}

func (s *SortedSlice[T]) Index(value T) int {
	index := s.search(value)
	if index < 0 || index >= s.Len() || s.slice[index] != value {
		return -1
	}
	return index
}

func (s *SortedSlice[T]) search(value T) int {
	if s.less == nil {
		panic("sortedslice: not initialized")
	}
	return sort.Search(len(s.slice), func(i int) bool {
		return !s.less(s.slice[i], value)
	})
}
