// SPDX-FileCopyrightText: 2025 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package go123_test

import (
	"maps"
	"testing"

	"gopkg.in/typ.v4/sync2"
)

func TestMapIterCompiles(t *testing.T) {
	m := new(sync2.Map[int, string])

	for range sync2.MapAll(m) {
		// do nothing
	}
	for range sync2.MapKeys(m) {
		// do nothing
	}
	for range sync2.MapValues(m) {
		// do nothing
	}

	otherMap := map[int]string{}
	sync2.MapInsert(m, maps.All(otherMap))
	_ = sync2.MapCollect(maps.All(otherMap))
}
