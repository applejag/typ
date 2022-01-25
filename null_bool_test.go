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
	"testing"
)

var (
	boolJSON = []byte(`true`)
)

func TestNullBoolFrom(t *testing.T) {
	b := NullFrom(true)
	assertBool(t, b, "NullFrom()")

	zero := NullFrom(false)
	if !zero.Valid {
		t.Error("NullFrom(false)", "is invalid, but should be valid")
	}
}

func TestNullBoolFromPtr(t *testing.T) {
	n := true
	bptr := &n
	b := NullFromPtr(bptr)
	assertBool(t, b, "NullFromPtr()")

	null := NullFromPtr[bool](nil)
	assertNull(t, null, "NullFromPtr(nil)")
}

func TestNullBoolUnmarshal(t *testing.T) {
	var null Null[bool]
	err := json.Unmarshal(nullJSON, &null)
	maybePanic(err)
	assertNull(t, null, "null json")
	if !null.Set {
		t.Error("should be Set", err)
	}

	var badType Null[bool]
	err = json.Unmarshal(int32JSON, &badType)
	if err == nil {
		panic("err should not be nil")
	}
	assertNull(t, badType, "wrong type json")

	var invalid Null[bool]
	err = invalid.UnmarshalJSON(invalidJSON)
	if _, ok := err.(*json.SyntaxError); !ok {
		t.Errorf("expected json.SyntaxError, not %T", err)
	}
}

func TestNullBoolTextUnmarshal(t *testing.T) {
	var b Null[bool]
	err := b.UnmarshalText([]byte("true"))
	maybePanic(err)
	assertBool(t, b, "UnmarshalText() bool")

	var zero Null[bool]
	err = zero.UnmarshalText([]byte("false"))
	maybePanic(err)
	assertFalseNull(t, zero, "UnmarshalText() false")

	var blank Null[bool]
	err = blank.UnmarshalText([]byte(""))
	maybePanic(err)
	assertNull(t, blank, "UnmarshalText() empty bool")

	var invalid Null[bool]
	err = invalid.UnmarshalText([]byte(":D"))
	if err == nil {
		panic("err should not be nil")
	}
	assertNull(t, invalid, "invalid json")
}

func TestNullBoolMarshal(t *testing.T) {
	b := NullFrom(true)
	data, err := json.Marshal(b)
	maybePanic(err)
	assertJSONEquals(t, data, "true", "non-empty json marshal")

	zero := NewNull(false, true)
	data, err = json.Marshal(zero)
	maybePanic(err)
	assertJSONEquals(t, data, "false", "zero json marshal")

	// invalid values should be encoded as null
	null := NewNull(false, false)
	data, err = json.Marshal(null)
	maybePanic(err)
	assertJSONEquals(t, data, "null", "null json marshal")
}

func TestNullBoolMarshalText(t *testing.T) {
	b := NullFrom(true)
	data, err := b.MarshalText()
	maybePanic(err)
	assertJSONEquals(t, data, "true", "non-empty text marshal")

	zero := NewNull(false, true)
	data, err = zero.MarshalText()
	maybePanic(err)
	assertJSONEquals(t, data, "false", "zero text marshal")

	// invalid values should be encoded as null
	null := NewNull(false, false)
	data, err = null.MarshalText()
	maybePanic(err)
	assertJSONEquals(t, data, "", "null text marshal")
}

func TestNullBoolPointer(t *testing.T) {
	b := NullFrom(true)
	ptr := b.Ptr()
	if *ptr != true {
		t.Errorf("bad %s bool: %#v ≠ %v\n", "pointer", ptr, true)
	}

	null := NewNull(false, false)
	ptr = null.Ptr()
	if ptr != nil {
		t.Errorf("bad %s bool: %#v ≠ %s\n", "nil pointer", ptr, "nil")
	}
}

func TestNullBoolIsZero(t *testing.T) {
	b := NullFrom(true)
	if b.IsZero() {
		t.Errorf("IsZero() should be false")
	}

	null := NewNull(false, false)
	if !null.IsZero() {
		t.Errorf("IsZero() should be true")
	}

	zero := NewNull(false, true)
	if zero.IsZero() {
		t.Errorf("IsZero() should be false")
	}
}

func TestNullBoolSetValid(t *testing.T) {
	change := NewNull(false, false)
	assertNull(t, change, "SetValid()")
	change.SetValid(true)
	assertBool(t, change, "SetValid()")
}

func TestNullBoolScan(t *testing.T) {
	var b Null[bool]
	err := b.Scan(true)
	maybePanic(err)
	assertBool(t, b, "scanned bool")

	var null Null[bool]
	err = null.Scan(nil)
	maybePanic(err)
	assertNull(t, null, "scanned null")
}

func assertBool(t *testing.T, b Null[bool], from string) {
	if b.Val != true {
		t.Errorf("bad %s bool: %v ≠ %v\n", from, b.Val, true)
	}
	if !b.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}

func assertFalseNull(t *testing.T, b Null[bool], from string) {
	if b.Val != false {
		t.Errorf("bad %s bool: %v ≠ %v\n", from, b.Val, false)
	}
	if !b.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}
