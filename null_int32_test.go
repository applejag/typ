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
	"math"
	"strconv"
	"testing"
)

var (
	int32JSON = []byte(`2147483646`)
)

func TestNullInt32From(t *testing.T) {
	i := NullFrom[int32](2147483646)
	assertInt32(t, i, "NullFrom[int32]()")

	zero := NullFrom[int32](0)
	if !zero.Valid {
		t.Error("NullFrom[int32](0)", "is invalid, but should be valid")
	}
}

func TestNullInt32FromPtr(t *testing.T) {
	n := int32(2147483646)
	iptr := &n
	i := NullFromPtr[int32](iptr)
	assertInt32(t, i, "NullFromPtr[int32]()")

	null := NullFromPtr[int32](nil)
	assertNull(t, null, "NullFromPtr(nil)")
}

func TestNullInt32Unmarshal(t *testing.T) {
	var i Null[int32]
	err := json.Unmarshal(int32JSON, &i)
	maybePanic(err)
	assertInt32(t, i, "int32 json")

	var null Null[int32]
	err = json.Unmarshal(nullJSON, &null)
	maybePanic(err)
	assertNull(t, null, "null json")
	if !null.Set {
		t.Error("should be Set")
	}

	var badType Null[int32]
	err = json.Unmarshal(boolJSON, &badType)
	if err == nil {
		panic("err should not be nil")
	}
	assertNull(t, badType, "wrong type json")

	var invalid Null[int32]
	err = invalid.UnmarshalJSON(invalidJSON)
	if _, ok := err.(*json.SyntaxError); !ok {
		t.Errorf("expected json.SyntaxError, not %T", err)
	}
	assertNull(t, invalid, "invalid json")
}

func TestNullInt32UnmarshalNonIntegerNumber32(t *testing.T) {
	var i Null[int32]
	err := json.Unmarshal(float64JSON, &i)
	if err == nil {
		panic("err should be present; non-integer number coerced to int32")
	}
}

func TestNullInt32UnmarshalOverflow(t *testing.T) {
	int32Overflow := uint32(math.MaxInt32)

	// Max int32 should decode successfully
	var i Null[int32]
	err := json.Unmarshal([]byte(strconv.FormatUint(uint64(int32Overflow), 10)), &i)
	maybePanic(err)

	// Attempt to overflow
	int32Overflow++
	err = json.Unmarshal([]byte(strconv.FormatUint(uint64(int32Overflow), 10)), &i)
	if err == nil {
		panic("err should be present; decoded value overflows int32")
	}
}

func TestNullInt32TextUnmarshal(t *testing.T) {
	var i Null[int32]
	err := i.UnmarshalText([]byte("2147483646"))
	maybePanic(err)
	assertInt32(t, i, "UnmarshalText() int32")

	var blank Null[int32]
	err = blank.UnmarshalText([]byte(""))
	maybePanic(err)
	assertNull(t, blank, "UnmarshalText() empty int32")
}

func TestNullInt32Marshal(t *testing.T) {
	i := NullFrom[int32](2147483646)
	data, err := json.Marshal(i)
	maybePanic(err)
	assertJSONEquals(t, data, "2147483646", "non-empty json marshal")

	// invalid values should be encoded as null
	null := NewNull[int32](0, false)
	data, err = json.Marshal(null)
	maybePanic(err)
	assertJSONEquals(t, data, "null", "null json marshal")
}

func TestNullInt32MarshalText(t *testing.T) {
	i := NullFrom[int32](2147483646)
	data, err := i.MarshalText()
	maybePanic(err)
	assertJSONEquals(t, data, "2147483646", "non-empty text marshal")

	// invalid values should be encoded as null
	null := NewNull[int32](0, false)
	data, err = null.MarshalText()
	maybePanic(err)
	assertJSONEquals(t, data, "", "null text marshal")
}

func TestNullInt32Pointer(t *testing.T) {
	i := NullFrom[int32](2147483646)
	ptr := i.Ptr()
	if *ptr != 2147483646 {
		t.Errorf("bad %s int32: %#v ≠ %d\n", "pointer", ptr, 2147483646)
	}

	null := NewNull[int32](0, false)
	ptr = null.Ptr()
	if ptr != nil {
		t.Errorf("bad %s int32: %#v ≠ %s\n", "nil pointer", ptr, "nil")
	}
}

func TestNullInt32IsZero(t *testing.T) {
	i := NullFrom[int32](2147483646)
	if i.IsZero() {
		t.Errorf("IsZero() should be false")
	}

	null := NewNull[int32](0, false)
	if !null.IsZero() {
		t.Errorf("IsZero() should be true")
	}

	zero := NewNull[int32](0, true)
	if zero.IsZero() {
		t.Errorf("IsZero() should be false")
	}
}

func TestNullInt32SetValid(t *testing.T) {
	change := NewNull[int32](0, false)
	assertNull(t, change, "SetValid()")
	change.SetValid(2147483646)
	assertInt32(t, change, "SetValid()")
}

func TestNullInt32Scan(t *testing.T) {
	var i Null[int32]
	err := i.Scan(2147483646)
	maybePanic(err)
	assertInt32(t, i, "scanned int32")

	var null Null[int32]
	err = null.Scan(nil)
	maybePanic(err)
	assertNull(t, null, "scanned null")
}

func assertInt32(t *testing.T, i Null[int32], from string) {
	if i.Val != 2147483646 {
		t.Errorf("bad %s int32: %d ≠ %d\n", from, i.Val, 2147483646)
	}
	if !i.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}
