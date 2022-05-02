// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package slices_test

import (
	"fmt"

	"gopkg.in/typ.v4/slices"
)

func ExampleTrim() {
	x := []int{0, 1, 2, 3, 1, 2, 1}
	fmt.Printf("All: %v\n", x)
	fmt.Printf("Trimmed: %v\n", slices.Trim(x, []int{0, 1}))

	// Output:
	// All: [0 1 2 3 1 2 1]
	// Trimmed: [2 3 1 2]
}

func ExampleDistinct() {
	values := []string{"a", "b", "b", "a"}

	fmt.Printf("All: %v\n", values)
	fmt.Printf("Distinct: %v\n", slices.Distinct(values))

	// Output:
	// All: [a b b a]
	// Distinct: [a b]
}
