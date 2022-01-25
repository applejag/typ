// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"encoding/json"
	"fmt"
	"time"
)

func ExampleNull() {
	type User struct {
		FirstName   string          `json:"firstName"`
		MiddleName  Null[string]    `json:"middleName"`
		LastName    string          `json:"lastName"`
		DateOfBirth Null[time.Time] `json:"dob"`
		DateOfDeath Null[time.Time] `json:"dod"`
	}
	user := User{
		FirstName:   "John",
		MiddleName:  Null[string]{},
		LastName:    "Doe",
		DateOfBirth: NullFrom(time.Date(1980, 5, 13, 0, 0, 0, 0, time.Local)),
		DateOfDeath: Null[time.Time]{},
	}

	bytes, _ := json.MarshalIndent(&user, "", "  ")
	fmt.Println(string(bytes))

	// Output:
	// {
	//   "firstName": "John",
	//   "middleName": null,
	//   "lastName": "Doe",
	//   "dob": "1980-05-13T00:00:00+02:00",
	//   "dod": null
	// }
}
