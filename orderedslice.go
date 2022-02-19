// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"fmt"
)

// NewOrderedSlice returns a new sorted slice based on a slice of values.
// Only ordered types are allowed. The values are sorted on insertion.
func NewOrderedSlice[T Ordered](values []T) OrderedSlice[T] {
	slice := make([]T, len(values))
	copy(slice, values)
	Sort(slice)
	return OrderedSlice[T]{slice}
}

// OrderedSlice is a slice of ordered values. The slice is always sorted thanks
// to only inserting values in a sorted order.
type OrderedSlice[T Ordered] struct {
	slice []T
}

func (s OrderedSlice[T]) String() string {
	return fmt.Sprint(s.slice)
}

func (s *OrderedSlice[T]) Get(index int) T {
	if index < 0 || index >= s.Len() {
		panic(fmt.Sprintf("orderedslice: index out of range [%d] with length %d", index, s.Len()))
	}
	return s.slice[index]
}

func (s *OrderedSlice[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.slice)
}

func (s *OrderedSlice[T]) Add(value T) int {
	if s == nil {
		panic("orderedslice: tried to add to nil orderedslice")
	}
	index := Search(s.slice, value)
	Insert(&s.slice, index, value)
	return index
}

func (s *OrderedSlice[T]) RemoveAt(index int) {
	if index < 0 || index >= s.Len() {
		panic(fmt.Sprintf("orderedslice: index out of range [%d] with length %d", index, s.Len()))
	}
	Remove(&s.slice, index)
}

func (s *OrderedSlice[T]) Remove(value T) int {
	index := s.Index(value)
	Remove(&s.slice, index)
	return index
}

func (s *OrderedSlice[T]) Contains(value T) bool {
	return s.Index(value) != -1
}

func (s *OrderedSlice[T]) Index(value T) int {
	index := Search(s.slice, value)
	if index < 0 || index >= s.Len() || s.slice[index] != value {
		return -1
	}
	return index
}
