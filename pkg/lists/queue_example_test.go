// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package lists_test

import (
	"fmt"

	"gopkg.in/typ.v3/pkg/lists"
)

func ExampleQueue() {
	var q lists.Queue[string]
	q.Enqueue("tripp")
	q.Enqueue("trapp")
	q.Enqueue("trull")

	fmt.Println(q.Len())     // 3
	fmt.Println(q.Dequeue()) // tripp, true
	fmt.Println(q.Dequeue()) // trapp, true
	fmt.Println(q.Dequeue()) // trull, true
	fmt.Println(q.Dequeue()) // "", false

	// Output:
	// 3
	// tripp true
	// trapp true
	// trull true
	//  false
}
