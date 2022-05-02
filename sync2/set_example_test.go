// SPDX-FileCopyrightText: 2022 Per Alexander Fougner
// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT
package sync2_test

import (
	"fmt"
)

func ExampleSet() {
	set := newSetABC()
	set.Range(func(value string) bool {
		fmt.Println("Value:", value)
		return true
	})

	// Unordered output:
	// Value: A
	// Value: B
	// Value: C
}

func ExampleSet_setOperations() {
	set1 := newSetABC()
	fmt.Println("set1:", set1) // {A B C}

	set2 := newSetBCD()
	fmt.Println("set2:", set2) // {B C D}

	fmt.Println("union:", set1.Union(set2))         // {A B C D}
	fmt.Println("intersect:", set1.Intersect(set2)) // {B C}
	fmt.Println("set diff:", set1.SetDiff(set2))    // {A}
	fmt.Println("sym diff:", set1.SymDiff(set2))    // {A D}

	// Please note: the Set.String() output is not sorted!
}
