// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package slices

import (
	"math/rand"
	"sort"

	"gopkg.in/typ.v4"
)

type sortOrdered[T typ.Ordered] []T

func (s sortOrdered[T]) Len() int {
	return len(s)
}

func (s sortOrdered[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortOrdered[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

type sortLess[T any] struct {
	slice []T
	less  func(a, b T) bool
}

func (s sortLess[T]) Len() int {
	return len(s.slice)
}

func (s sortLess[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func (s sortLess[T]) Less(i, j int) bool {
	return s.less(s.slice[i], s.slice[j])
}

// Sort will sort a slice using the default less-than operator.
func Sort[S ~[]E, E typ.Ordered](slice S) {
	sort.Sort(sortOrdered[E](slice))
}

// SortFunc will sort a slice using the given less function.
func SortFunc[S ~[]E, E any](slice S, less func(a, b E) bool) {
	sort.Sort(sortLess[E]{slice, less})
}

// SortDesc will sort a slice using the default less-than operator in
// descending order.
func SortDesc[S ~[]E, E typ.Ordered](slice S) {
	sort.Sort(sort.Reverse(sortOrdered[E](slice)))
}

// SortDescFunc will sort a slice using the given less function in
// descending order.
func SortDescFunc[S ~[]E, E any](slice S, less func(a, b E) bool) {
	sort.Sort(sort.Reverse(sortLess[E]{slice, less}))
}

// SortStableFunc will sort a slice using the given less function, while keeping
// the original order of equal elements.
func SortStableFunc[S ~[]E, E any](slice S, less func(a, b E) bool) {
	sort.Stable(sortLess[E]{slice, less})
}

// SortStableDescFunc will sort a slice using the given less function in
// descending order, while keeping the original order of equal elements.
func SortStableDescFunc[S ~[]E, E any](slice S, less func(a, b E) bool) {
	sort.Stable(sort.Reverse(sortLess[E]{slice, less}))
}

// Reverse will reverse all elements inside a slice, in place.
func Reverse[S ~[]E, E any](slice S) {
	for i, j := 0, len(slice)-1; i < len(slice)/2; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle will randomize the order of all elements inside a slice. It uses the
// rand package for random number generation, so you are expected to have called
// rand.Seed beforehand.
func Shuffle[S ~[]E, E any](slice S) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// ShuffleRand will randomize the order of all elements inside a slice using the
// Fisher-Yates shuffle algoritm. It uses the rand argument for random number
// generation.
func ShuffleRand[S ~[]E, E any](slice S, rand *rand.Rand) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// BinarySearch performs a binary search to find the index of a value in a
// sorted slice of ordered values. The index of the first match is returned, or
// the index where it insert the value if the value is not present.
// The slice must be sorted in ascending order.
func BinarySearch[S ~[]E, E typ.Ordered](slice S, value E) int {
	return sort.Search(len(slice), func(i int) bool {
		return slice[i] >= value
	})
}

// BinarySearchFunc performs a binary search to find the index of a value in a
// sorted slice of ordered values. The index of the first match is returned, or
// the index where it insert the value if the value is not present.
// The slice must be sorted in ascending order.
//
// The less function should return true if the given value is less than the
// sought value.
func BinarySearchFunc[S ~[]E, E any](slice S, less func(a E) bool) int {
	return sort.Search(len(slice), func(i int) bool {
		return !less(slice[i])
	})
}
