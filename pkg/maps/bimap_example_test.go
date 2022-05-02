// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package maps_test

import (
	"fmt"

	"gopkg.in/typ.v3/pkg/maps"
)

func ExampleBimap() {
	var bimap maps.Bimap[int, string]

	bimap.Add(1, "foo")
	bimap.Add(2, "bar")
	bimap.Add(3, "moo")
	bimap.Add(4, "doo")

	fmt.Println(bimap.GetForward(4))
	fmt.Println(bimap.GetReverse("moo"))

	// Output:
	// doo true
	// 3 true
}
