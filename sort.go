// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"constraints"
	"math/rand"
	"sort"
)

// SortOrdered implements sort.Interface via the default less-than operator.
type SortOrdered[T constraints.Ordered] []T

func (s SortOrdered[T]) Len() int {
	return len(s)
}

func (s SortOrdered[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortOrdered[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

// Sort will sort a slice using the default less-than operator.
func Sort[T constraints.Ordered](slice []T) {
	sort.Sort(SortOrdered[T](slice))
}

// SortDesc will sort a slice using the default less-than operator in
// descending order.
func SortDesc[T constraints.Ordered](slice []T) {
	sort.Sort(sort.Reverse(SortOrdered[T](slice)))
}

// Reverse will reverse all elements inside a slice, in place.
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < len(slice)/2; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle will randomize the order of all elements inside a slice. It uses the
// rand package for random number generation, so you are expected to have called
// rand.Seed beforehand.
func Shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// ShuffleRand will randomize the order of all elements inside a slice using the
// Fisher-Yates shuffle algoritm. It uses the rand argument for random number
// generation.
func ShuffleRand[T any](slice []T, rand *rand.Rand) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// Search performs a binary search to find the index of a value in a sorted
// slice of ordered values. The index of the first match is returned, or the
// index where it insert the value if the value is not present.
// The slice must be sorted in ascending order.
func Search[T constraints.Ordered](slice []T, value T) int {
	return sort.Search(len(slice), func(i int) bool {
		return slice[i] >= value
	})
}

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

// Index returns the index of the first occurence where the function returns
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
