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
	"time"
)

var (
	timeString   = "2012-12-21T21:21:21Z"
	timeJSON     = []byte(`"` + timeString + `"`)
	nullTimeJSON = []byte(`null`)
	timeValue, _ = time.Parse(time.RFC3339, timeString)
	badObject    = []byte(`{"hello": "world"}`)
)

func TestNullTimeUnmarshalJSON(t *testing.T) {
	var ti Null[time.Time]
	err := json.Unmarshal(timeJSON, &ti)
	maybePanic(err)
	assertTime(t, ti, "UnmarshalJSON() json")

	var null Null[time.Time]
	err = json.Unmarshal(nullTimeJSON, &null)
	maybePanic(err)
	assertNull(t, null, "null time json")
	if !null.Set {
		t.Error("should be Set")
	}

	var invalid Null[time.Time]
	err = invalid.UnmarshalJSON(invalidJSON)
	if _, ok := err.(*json.SyntaxError); !ok {
		t.Errorf("expected json.SyntaxError, not %T", err)
	}
	assertNull(t, invalid, "invalid from object json")

	var bad Null[time.Time]
	err = json.Unmarshal(badObject, &bad)
	if err == nil {
		t.Errorf("expected error: bad object")
	}
	assertNull(t, bad, "bad from object json")

	var wrongType Null[time.Time]
	err = json.Unmarshal(int32JSON, &wrongType)
	if err == nil {
		t.Errorf("expected error: wrong type JSON")
	}
	assertNull(t, wrongType, "wrong type object json")
}

func TestNullTimeUnmarshalText(t *testing.T) {
	ti := NullFrom(timeValue)
	txt, err := ti.MarshalText()
	maybePanic(err)
	assertJSONEquals(t, txt, timeString, "marshal text")

	var unmarshal Null[time.Time]
	err = unmarshal.UnmarshalText(txt)
	maybePanic(err)
	assertTime(t, unmarshal, "unmarshal text")

	var invalid Null[time.Time]
	err = invalid.UnmarshalText([]byte("hello world"))
	if err == nil {
		t.Error("expected error")
	}
	assertNull(t, invalid, "bad string")
}

func TestNullTimeMarshal(t *testing.T) {
	ti := NullFrom(timeValue)
	data, err := json.Marshal(ti)
	maybePanic(err)
	assertJSONEquals(t, data, string(timeJSON), "non-empty json marshal")

	ti.Valid = false
	data, err = json.Marshal(ti)
	maybePanic(err)
	assertJSONEquals(t, data, string(nullJSON), "null json marshal")
}

func TestNullTimeFrom(t *testing.T) {
	ti := NullFrom(timeValue)
	assertTime(t, ti, "TimeFrom() time.Time")
}

func TestNullTimeFromPtr(t *testing.T) {
	ti := NullFromPtr(&timeValue)
	assertTime(t, ti, "TimeFromPtr() time")

	null := NullFromPtr[time.Time](nil)
	assertNull(t, null, "TimeFromPtr(nil)")
}

func TestNullTimeSetValid(t *testing.T) {
	var ti time.Time
	change := NewNull(ti, false)
	assertNull(t, change, "SetValid()")
	change.SetValid(timeValue)
	assertTime(t, change, "SetValid()")
}

func TestNullTimePointer(t *testing.T) {
	ti := NullFrom(timeValue)
	ptr := ti.Ptr()
	if *ptr != timeValue {
		t.Errorf("bad %s time: %#v ≠ %v\n", "pointer", ptr, timeValue)
	}

	var nt time.Time
	null := NewNull(nt, false)
	ptr = null.Ptr()
	if ptr != nil {
		t.Errorf("bad %s time: %#v ≠ %s\n", "nil pointer", ptr, "nil")
	}
}

func TestNullTimeIsZero(t *testing.T) {
	ti := NullFrom(time.Now())
	if ti.IsZero() {
		t.Errorf("IsZero() should be false")
	}

	null := NullFromPtr[time.Time](nil)
	if !null.IsZero() {
		t.Errorf("IsZero() should be true")
	}
}

func TestNullTimeScanValue(t *testing.T) {
	var ti Null[time.Time]
	err := ti.Scan(timeValue)
	maybePanic(err)
	assertTime(t, ti, "scanned time")
	if v, err := ti.Value(); v != timeValue || err != nil {
		t.Error("bad value or err:", v, err)
	}

	var null Null[time.Time]
	err = null.Scan(nil)
	maybePanic(err)
	assertNull(t, null, "scanned null")
	if v, err := null.Value(); v != nil || err != nil {
		t.Error("bad value or err:", v, err)
	}

	var wrong Null[time.Time]
	err = wrong.Scan(int64(42))
	if err == nil {
		t.Error("expected error")
	}
	assertNull(t, wrong, "scanned wrong")
}

func assertTime(t *testing.T, ti Null[time.Time], from string) {
	if ti.Val != timeValue {
		t.Errorf("bad %v time: %v ≠ %v\n", from, ti.Val, timeValue)
	}
	if !ti.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}
