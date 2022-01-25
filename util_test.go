// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"io"
	"testing"
)

func TestIsNil(t *testing.T) {
	assertIsFalse(t, "0", IsNil(0))
	assertIsFalse(t, `""`, IsNil(""))
	assertIsFalse(t, "any(0)", IsNil(any(0)))
	assertIsTrue(t, "any(nil)", IsNil(any(nil)))
	assertIsFalse(t, "error(io.EOF)", IsNil(error(io.EOF)))
	assertIsTrue(t, "error(nil)", IsNil(error(nil)))
	var x error
	var xAsAny any = x
	assertIsTrue(t, "any(error(nil))", IsNil(xAsAny))
}

func assertIsTrue(t *testing.T, name string, b bool) {
	if !b {
		t.Errorf("%s: want true, got false", name)
	}
}

func assertIsFalse(t *testing.T, name string, b bool) {
	if b {
		t.Errorf("%s: want false, got true", name)
	}
}
