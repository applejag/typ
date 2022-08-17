// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package slices_test

import (
	"fmt"

	"gopkg.in/typ.v4/slices"
)

type User struct {
	Name    string
	IsAdmin bool
}

func (u User) String() string {
	if u.IsAdmin {
		return fmt.Sprintf("%s (admin)", u.Name)
	}
	return u.Name
}

func (u User) AsAdmin() User {
	u.IsAdmin = true
	return u
}

func ExampleNewSorted() {
	slice := slices.NewSorted([]User{}, func(a, b User) bool {
		return a.Name < b.Name
	})
	johnDoe := User{Name: "John"}
	slice.Add(johnDoe)
	slice.Add(User{Name: "Jane"})
	slice.Add(User{Name: "Ann"})
	slice.Add(User{Name: "Wayne"})

	fmt.Println(slice)
	fmt.Println("Contains John non-admin?", slice.Contains(johnDoe))
	fmt.Println("Contains John admin?", slice.Contains(johnDoe.AsAdmin()))
	slice.Remove(johnDoe)
	fmt.Println(slice)

	// Output:
	// [Ann Jane John Wayne]
	// Contains John non-admin? true
	// Contains John admin? false
	// [Ann Jane Wayne]
}
