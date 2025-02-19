// SPDX-FileCopyrightText: 2025 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sync2_test

import (
	"fmt"

	"gopkg.in/typ.v4/sync2"
)

func ExampleMap() {
	m := new(sync2.Map[int, string])
	m.Store(1, "one")
	m.Store(2, "two")
	m.Store(3, "three")

	fmt.Println("Len:", m.Len())

	// Output:
	// Len: 3
}

func ExampleMap_Load() {
	m := new(sync2.Map[int, string])
	m.Store(1, "one")
	m.Store(2, "two")
	m.Store(3, "three")

	if value, ok := m.Load(0); ok {
		fmt.Println("Map[0]:", value)
	} else {
		fmt.Println("Map[0]: (no value)")
	}

	if value, ok := m.Load(1); ok {
		fmt.Println("Map[1]:", value)
	} else {
		fmt.Println("Map[1]: (no value)")
	}

	// Output:
	// Map[0]: (no value)
	// Map[1]: one
}
