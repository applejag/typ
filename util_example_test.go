// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ_test

import (
	"fmt"
	"path/filepath"
	"time"

	"gopkg.in/typ.v1"
)

func ExampleTrim() {
	x := []int{0, 1, 2, 3, 1, 2, 1}
	fmt.Printf("All: %v\n", x)
	fmt.Printf("Trimmed: %v\n", typ.Trim(x, []int{0, 1}))

	// Output:
	// All: [0 1 2 3 1 2 1]
	// Trimmed: [2 3 1 2]
}

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

func ExampleNewOf() {
	myVector := new(struct {
		X int
		Y int
		Z int
	})

	otherVector := typ.NewOf(myVector)

	fmt.Println("myVector:", myVector)
	fmt.Println("otherVector:", otherVector)
	fmt.Println("same?:", myVector == otherVector) // different pointers

	// Output:
	// myVector: &{0 0 0}
	// otherVector: &{0 0 0}
	// same?: false
}

func ExampleMakeSliceOfSlice() {
	var users = []struct {
		ID       int
		Username string
		Admin    bool
	}{
		{1, "foo", true},
		{2, "bar", false},
	}

	var admins = typ.MakeSliceOfSlice(users)
	for _, u := range users {
		if u.Admin {
			admins = append(admins, u)
		}
	}

	fmt.Println("users:", len(users))
	fmt.Println("admins:", len(admins))

	// Output:
	// users: 2
	// admins: 1
}

func ExampleMakeMapOfMap() {
	files := map[string]struct {
		SizeMB    int
		CreatedAt time.Time
		UpdatedAt time.Time
	}{
		"/root/gopher.png": {
			SizeMB:    10,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		"/root/meaningoflife.txt": {
			SizeMB:    0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	filteredFiles := typ.MakeMapOfMap(files)
	for path, f := range files {
		if ok, _ := filepath.Match("/root/*.txt", path); ok {
			filteredFiles[path] = f
		}
	}
	fmt.Println("Filtered:")
	for path, f := range filteredFiles {
		fmt.Println(path, ":", f.SizeMB, "MB")
	}
	// Output:
	// Filtered:
	// /root/meaningoflife.txt : 0 MB
}
