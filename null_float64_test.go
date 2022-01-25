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
	float64JSON = []byte(`1.2345`)
)

func TestNullFloat64From(t *testing.T) {
	f := NullFrom(1.2345)
	assertFloat64(t, f, "NullFrom[float64]()")

	zero := NullFrom[float64](0)
	if !zero.Valid {
		t.Error("NullFrom[float64](0)", "is invalid, but should be valid")
	}
}

func TestNullFloat64FromPtr(t *testing.T) {
	n := float64(1.2345)
	iptr := &n
	f := NullFromPtr(iptr)
	assertFloat64(t, f, "NullFrom[float64]Ptr()")

	null := NullFromPtr[float64](nil)
	assertNull(t, null, "NullFrom[float64]Ptr(nil)")
}

func TestNullFloat64Unmarshal(t *testing.T) {
	var f Null[float64]
	err := json.Unmarshal(float64JSON, &f)
	maybePanic(err)
	assertFloat64(t, f, "float64 json")

	var null Null[float64]
	err = json.Unmarshal(nullJSON, &null)
	maybePanic(err)
	assertNull(t, null, "null json")
	if !null.Set {
		t.Error("should be Set")
	}

	var badType Null[float64]
	err = json.Unmarshal(boolJSON, &badType)
	if err == nil {
		panic("err should not be nil")
	}
	assertNull(t, badType, "wrong type json")

	var invalid Null[float64]
	err = invalid.UnmarshalJSON(invalidJSON)
	if _, ok := err.(*json.SyntaxError); !ok {
		t.Errorf("expected json.SyntaxError, not %T", err)
	}
}

func TestNullFloat64TextUnmarshal(t *testing.T) {
	var f Null[float64]
	err := f.UnmarshalText([]byte("1.2345"))
	maybePanic(err)
	assertFloat64(t, f, "UnmarshalText() float64")

	var blank Null[float64]
	err = blank.UnmarshalText([]byte(""))
	maybePanic(err)
	assertNull(t, blank, "UnmarshalText() empty float64")
}

func TestNullFloat64Marshal(t *testing.T) {
	f := NullFrom(1.2345)
	data, err := json.Marshal(f)
	maybePanic(err)
	assertJSONEquals(t, data, "1.2345", "non-empty json marshal")

	// invalid values should be encoded as null
	null := NewNull[float64](0, false)
	data, err = json.Marshal(null)
	maybePanic(err)
	assertJSONEquals(t, data, "null", "null json marshal")
}

func TestNullFloat64MarshalText(t *testing.T) {
	f := NullFrom(1.2345)
	data, err := f.MarshalText()
	maybePanic(err)
	assertJSONEquals(t, data, "1.2345", "non-empty text marshal")

	// invalid values should be encoded as null
	null := NewNull[float64](0, false)
	data, err = null.MarshalText()
	maybePanic(err)
	assertJSONEquals(t, data, "", "null text marshal")
}

func TestNullFloat64Pointer(t *testing.T) {
	f := NullFrom(1.2345)
	ptr := f.Ptr()
	if *ptr != 1.2345 {
		t.Errorf("bad %s float64: %#v ≠ %v\n", "pointer", ptr, 1.2345)
	}

	null := NewNull[float64](0, false)
	ptr = null.Ptr()
	if ptr != nil {
		t.Errorf("bad %s float64: %#v ≠ %s\n", "nil pointer", ptr, "nil")
	}
}

func TestNullFloat64IsZero(t *testing.T) {
	f := NullFrom(1.2345)
	if f.IsZero() {
		t.Errorf("IsZero() should be false")
	}

	null := NewNull[float64](0, false)
	if !null.IsZero() {
		t.Errorf("IsZero() should be true")
	}

	zero := NewNull[float64](0, true)
	if zero.IsZero() {
		t.Errorf("IsZero() should be false")
	}
}

func TestFloat64SetValid(t *testing.T) {
	change := NewNull[float64](0, false)
	assertNull(t, change, "SetValid()")
	change.SetValid(1.2345)
	assertFloat64(t, change, "SetValid()")
}

func TestFloat64Scan(t *testing.T) {
	var f Null[float64]
	err := f.Scan(1.2345)
	maybePanic(err)
	assertFloat64(t, f, "scanned float64")

	var null Null[float64]
	err = null.Scan(nil)
	maybePanic(err)
	assertNull(t, null, "scanned null")
}

func assertFloat64(t *testing.T, f Null[float64], from string) {
	if f.Val != 1.2345 {
		t.Errorf("bad %s float64: %f ≠ %f\n", from, f.Val, 1.2345)
	}
	if !f.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}
