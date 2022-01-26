// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ_test

import (
	"fmt"

	"gopkg.in/typ.v1"
)

func ExampleOrderedTree() {
	var tree typ.OrderedTree[string]

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
