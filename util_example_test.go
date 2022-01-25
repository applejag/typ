// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ_test

import (
	"fmt"

	"gopkg.in/typ.v1"
)

func ExampleDistinct() {
	values := []string{"a", "b", "b", "a"}

	fmt.Printf("All: %v\n", values)
	fmt.Printf("Distinct: %v\n", typ.Distinct(values))

	// Output:
	// All: [a b b a]
	// Distinct: [a b]
}

func ExampleCoal() {
	bindAddressFromUser := ""
	bindAddressDefault := "localhost:8080"

	fmt.Println("Adress 1:", typ.Coal(bindAddressFromUser, bindAddressDefault))

	bindAddressFromUser = "192.168.1.10:80"
	fmt.Println("Adress 2:", typ.Coal(bindAddressFromUser, bindAddressDefault))

	// Output:
	// Adress 1: localhost:8080
	// Adress 2: 192.168.1.10:80
}

func ExampleTern() {
	age := 16
	fmt.Println("To drink I want a glass of", typ.Tern(age >= 18, "wine", "milk"))

	// Output:
	// To drink I want a glass of milk
}
