// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sets_test

import (
	"fmt"

	"gopkg.in/typ.v3/pkg/sets"
)

func ExampleSet() {
	set := make(sets.Set[string])
	set.Add("A")
	set.Add("B")
	set.Add("C")

	for value := range set {
		fmt.Println("Value:", value)
	}

	// Unordered output:
	// Value: A
	// Value: B
	// Value: C
}

func ExampleSet_setOperations() {
	set1 := make(sets.Set[string])
	set1.Add("A")
	set1.Add("B")
	set1.Add("C")
	fmt.Println("set1:", set1) // {A B C}

	set2 := make(sets.Set[string])
	set2.Add("B")
	set2.Add("C")
	set2.Add("D")
	fmt.Println("set2:", set2) // {B C D}

	fmt.Println("union:", set1.Union(set2))         // {A B C D}
	fmt.Println("intersect:", set1.Intersect(set2)) // {B C}
	fmt.Println("set diff:", set1.SetDiff(set2))    // {A}
	fmt.Println("sym diff:", set1.SymDiff(set2))    // {A D}

	// Please note: the Set.String() output is not sorted!
}
