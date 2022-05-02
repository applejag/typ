// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package lists_test

import (
	"fmt"

	"gopkg.in/typ.v4/lists"
)

func ExampleStack() {
	var s lists.Stack[string]
	s.Push("tripp")
	s.Push("trapp")
	s.Push("trull")

	fmt.Println(len(s))  // 3
	fmt.Println(s.Pop()) // trull, true
	fmt.Println(s.Pop()) // trapp, true
	fmt.Println(s.Pop()) // tripp, true
	fmt.Println(s.Pop()) // "", false

	// Output:
	// 3
	// trull true
	// trapp true
	// tripp true
	//  false
}
