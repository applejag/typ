// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package slices

import (
	"testing"

	"gopkg.in/typ.v4"
	"gopkg.in/typ.v4/internal/assert"
)

func TestSorted(t *testing.T) {
	slice := NewSorted([]string{}, func(a, b string) bool {
		return a < b
	})
	testSortedStrings(t, slice)
}

func TestSortedOrdered(t *testing.T) {
	slice := NewSortedOrdered([]string{})
	testSortedStrings(t, slice)
}

func TestSortedCompare(t *testing.T) {
	slice := NewSortedCompare([]string{}, typ.Compare[string])
	testSortedStrings(t, slice)
}

func TestSortedFunc(t *testing.T) {
	slice := NewSortedFunc([]string{}, func(a string) string {
		return a
	})
	testSortedStrings(t, slice)
}

func testSortedStrings(t *testing.T, slice Sorted[string]) {
	assert.Comparable(t, "a index", 0, slice.Add("a"))
	assert.Comparable(t, "d index", 1, slice.Add("d"))
	assert.Comparable(t, "b index", 1, slice.Add("b"))
	assert.Comparable(t, "c index", 2, slice.Add("c"))

	assert.Comparable(t, "contains a", true, slice.Contains("a"))
	assert.Comparable(t, "contains b", true, slice.Contains("b"))
	assert.Comparable(t, "contains c", true, slice.Contains("c"))
	assert.Comparable(t, "contains d", true, slice.Contains("d"))

	assert.Comparable(t, "contains e", false, slice.Contains("e"))
}
