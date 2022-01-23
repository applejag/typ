// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// SPDX-FileCopyrightText: 2009 The Go Authors
//
// SPDX-License-Identifier: LicenseRef-Go-BSD

package typ_test

import (
	"fmt"

	"gopkg.in/typ.v0"
)

func ExampleList() {
	// Create a new list and put some numbers in it.
	l := typ.NewList[int]()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
