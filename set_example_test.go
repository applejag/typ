// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ_test

import (
	"fmt"

	"gopkg.in/typ.v1"
)

func ExampleSet() {
	set1 := make(typ.Set[string])
	set1.Set("A")
	set1.Set("B")
	set1.Set("C")
	fmt.Println("set1:", set1) // {A B C}

	set2 := make(typ.Set[string])
	set2.Set("B")
	set2.Set("C")
	set2.Set("D")
	fmt.Println("set2:", set2) // {B C D}

	fmt.Println("union:", set1.Union(set2))         // {A B C D}
	fmt.Println("intersect:", set1.Intersect(set2)) // {B C}
	fmt.Println("set diff:", set1.SetDiff(set2))    // {A}
	fmt.Println("sym diff:", set1.SymDiff(set2))    // {A D}

	// Please note: the Set.String() output is not sorted!
}
