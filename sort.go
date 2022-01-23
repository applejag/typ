// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"constraints"
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
