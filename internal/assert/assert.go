// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package assert

import (
	"sort"
	"strings"
	"testing"
)

// Comparable asserts two comparable values that they equal.
func Comparable[T comparable](t *testing.T, name string, want T, got T) {
	if want != got {
		t.Errorf(`%s: want "%v", got "%v"`, name, want, got)
	}
}

// ElementsMatch asserts that two slices of strings contain the same values,
// with no regard to the ordering.
func ElementsMatch[S ~[]string](t *testing.T, want, got S) {
	var wantClone S
	copy(wantClone, want)
	sort.Strings(wantClone)
	var gotClone S
	copy(gotClone, got)
	sort.Strings(gotClone)

	wantStr := strings.Join(wantClone, ", ")
	gotStr := strings.Join(gotClone, ", ")
	if wantStr != gotStr {
		t.Errorf("want %q, got %q", wantStr, gotStr)
	}
}
