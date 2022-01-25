// Copyright for portions of project null-extended are held by *Greg Roseberry, 2014* as part of project null.
// All other copyright for project null-extended are held by *Patrick O'Brien, 2016*.
// All rights reserved.
//
// SPDX-FileCopyrightText: 2014 Greg Roseberry
// SPDX-FileCopyrightText: 2016 Patrick O'Brien
//
// SPDX-License-Identifier: BSD-2-Clause

package typ

import (
	"encoding/json"
	"reflect"
	"testing"
)

var (
	nullJSON    = []byte(`null`)
	invalidJSON = []byte(`:)`)
)

func TestNullImplementsMarshaler(t *testing.T) {
	marshalerType := reflect.TypeOf((*json.Marshaler)(nil)).Elem()
	null := NullFrom(0)
	assertIsTrue(t, "int", reflect.TypeOf(&null).Implements(marshalerType))
}

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func assertJSONEquals(t *testing.T, data []byte, cmp string, from string) {
	if string(data) != cmp {
		t.Errorf("bad %s data: %s ≠ %s\n", from, data, cmp)
	}
}

func assertNull[T any](t *testing.T, i Null[T], from string) {
	if i.Valid {
		t.Error(from, "is valid, but should be invalid")
	}
}
