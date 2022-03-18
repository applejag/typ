// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package assert

import "testing"

// Comparable asserts two comparable values that they equal.
func Comparable[T comparable](t *testing.T, name string, want T, got T) {
	if want != got {
		t.Errorf(`%s: want "%v", got "%v"`, name, want, got)
	}
}
