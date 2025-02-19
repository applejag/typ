// SPDX-FileCopyrightText: 2025 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sync2_test

import (
	"gopkg.in/typ.v4/sync2"
)

func ExampleMapAll() {
	m := new(sync2.Map[int, string])
	m.Store(1, "one")
	m.Store(2, "two")

	// In Go 1.23+ you can iterate the values like so:
	/*
		for k, v := range sync2.MapAll(m) {
			fmt.Printf("%s: %q\n", k, v)
		}
	*/
}

func ExampleMapKeys() {
	m := new(sync2.Map[int, string])
	m.Store(1, "one")
	m.Store(2, "two")

	// In Go 1.23+ you can iterate the values like so:
	/*
		for k := range sync2.MapKeys(m) {
			fmt.Printf("key: %s\n", k)
		}
	*/
}

func ExampleMapValues() {
	m := new(sync2.Map[int, string])
	m.Store(1, "one")
	m.Store(2, "two")

	// In Go 1.23+ you can iterate the values like so:
	/*
		for k := range sync2.MapValues(m) {
			fmt.Printf("value: %q\n", v)
		}
	*/

	// Alternatively you can get a slice of keys like so:
	/*
		keys := slices.Collect(sync2.MapValues(m))
	*/
}
