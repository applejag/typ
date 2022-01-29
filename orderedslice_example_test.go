// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ_test

import (
	"fmt"

	"gopkg.in/typ.v1"
)

func ExampleOrderedSlice() {
	var slice typ.OrderedSlice[string]
	slice.Add("f")
	slice.Add("b")
	slice.Add("e")
	slice.Add("a")
	slice.Add("d")
	slice.Add("c")
	slice.Add("g")

	fmt.Println(slice)
	fmt.Println("Contains a?", slice.Contains("a"))
	slice.Remove("d")
	fmt.Println(slice)

	// Output:
	// [a b c d e f g]
	// Contains a? true
	// [a b c e f g]
}
