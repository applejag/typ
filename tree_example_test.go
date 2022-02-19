// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ_test

import (
	"fmt"

	"gopkg.in/typ.v1"
)

func ExampleNewAVLTreeOrdered() {
	tree := typ.NewAVLTreeOrdered[string]()

	// Unordered input
	tree.Add("E")
	tree.Add("B")
	tree.Add("D")
	tree.Add("C")
	tree.Add("A")

	// Sorted output
	fmt.Println(tree.Len(), tree)

	// Output:
	// 5 [A B C D E]
}

func ExampleNewAVLTree() {
	type Name struct {
		First string
		Last  string
	}

	// Sort first on first name, then on last name
	tree := typ.NewAVLTree(func(a, b Name) int {
		v := typ.Compare(a.First, b.First)
		if v == 0 {
			v = typ.Compare(a.Last, b.Last)
		}
		return v
	})

	// Unordered input
	tree.Add(Name{"Bert", "Screws"})
	tree.Add(Name{"John", "Doe"})
	tree.Add(Name{"Anchor", "Shippie"})
	tree.Add(Name{"Bert", "Horton"})
	tree.Add(Name{"Jane", "Doe"})

	// Sorted output
	fmt.Println(tree.Len(), tree)

	// Output:
	// 5 [{Anchor Shippie} {Bert Horton} {Bert Screws} {Jane Doe} {John Doe}]
}
