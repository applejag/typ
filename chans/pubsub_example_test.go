// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package chans_test

import (
	"fmt"
	"sync"

	"gopkg.in/typ.v4/chans"
)

func printMessages(prefix string, ch <-chan string, wg *sync.WaitGroup) {
	for msg := range ch {
		fmt.Println(prefix, msg)
	}
	wg.Done()
}

func ExamplePubSub() {
	var pub chans.PubSub[string]
	var wg sync.WaitGroup

	sub1 := pub.Sub()
	sub2 := pub.Sub()

	wg.Add(2)
	go printMessages("sub1:", sub1, &wg)
	go printMessages("sub2:", sub2, &wg)

	pub.PubWait("hello there")
	pub.UnsubAll()
	wg.Wait()

	// Unordered output:
	// sub1: hello there
	// sub2: hello there
}
