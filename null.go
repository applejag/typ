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
	"bytes"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/typ.v0/internal/convert"
)

var nullBytes = []byte("null")
var trueBytes = []byte("true")
var falseBytes = []byte("false")

// Null is a nullable value.
type Null[T any] struct {
	Val   T
	Valid bool
	Set   bool
}

// NewNull creates a new nullable value.
func NewNull[T any](value T, valid bool) Null[T] {
	return Null[T]{
		Val:   value,
		Valid: valid,
		Set:   true,
	}
}

// NewNullFrom creates a new nullable value that will be marked invalid if nil,
// and will always be valid if type is not nullable.
func NewNullFrom[T any](value T) Null[T] {
	return NewNull(value, !IsNil(value))
}

// NewNullFromPtr creates a new nullable value that will be marked invalid if nil.
func NewNullFromPtr[T any](ptr *T) Null[T] {
	if ptr == nil {
		return NewNull(Zero[T](), false)
	}
	return NewNull(*ptr, true)
}

// IsValid returns true if this carries an explicit value and is not null.
func (n Null[T]) IsValid() bool {
	return n.Set && n.Valid
}

// IsSet returns true if this carries an explicit value (null inclusive).
func (n Null[T]) IsSet() bool {
	return n.Set
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *Null[T]) UnmarshalJSON(data []byte) error {
	n.Set = true
	if bytes.Equal(data, nullBytes) {
		n.Val = Zero[T]()
		n.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &n.Val); err != nil {
		return err
	}
	n.Valid = true
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *Null[T]) UnmarshalText(text []byte) error {
	n.Set = true
	if len(text) == 0 {
		n.Valid = false
		return nil
	}
	v, err := parseText[T](text)
	if err != nil {
		return err
	}
	n.Valid = true
	n.Val = v.(T)
	return nil
}

func parseText[T any](text []byte) (any, error) {
	var zero any = Zero[T]() // even for nil values, this holds the type data
	switch val := zero.(type) {
	case encoding.TextUnmarshaler:
		err := val.UnmarshalText(text)
		return val, err
	case bool:
		switch {
		case bytes.Equal(text, trueBytes):
			return true, nil
		case bytes.Equal(text, falseBytes):
			return false, nil
		default:
			return nil, errors.New("invalid input: " + string(text))
		}
	case byte:
		if len(text) > 1 {
			return nil, errors.New("text: cannot convert to byte, text len is greater than one")
		}
		return text[0], nil
	case []byte:
		b := make([]byte, len(text))
		copy(text, b)
		return b, nil
	case float32:
		res, err := strconv.ParseFloat(string(text), 32)
		return TernCast[float32](err != nil, res, 0), err
	case float64:
		res, err := strconv.ParseFloat(string(text), 64)
		return res, err
	case int:
		res, err := strconv.ParseInt(string(text), 10, 0)
		return TernCast(err != nil, res, int(0)), err
	case int8:
		res, err := strconv.ParseInt(string(text), 10, 8)
		return TernCast(err != nil, res, int8(0)), err
	case int16:
		res, err := strconv.ParseInt(string(text), 10, 16)
		return TernCast(err != nil, res, int16(0)), err
	case int32:
		res, err := strconv.ParseInt(string(text), 10, 32)
		return TernCast(err != nil, res, int32(0)), err
	case int64:
		res, err := strconv.ParseInt(string(text), 10, 64)
		return TernCast(err != nil, res, int64(0)), err
	case uint:
		res, err := strconv.ParseInt(string(text), 10, 0)
		return TernCast(err != nil, res, uint(0)), err
	// Collides with byte.
	//case uint8:
	//	res, err := strconv.ParseInt(string(text), 10, 8)
	//	return TernCast(err != nil, res, uint8(0)), err
	case uint16:
		res, err := strconv.ParseInt(string(text), 10, 16)
		return TernCast(err != nil, res, uint16(0)), err
	case uint32:
		res, err := strconv.ParseInt(string(text), 10, 32)
		return TernCast(err != nil, res, uint32(0)), err
	case uint64:
		res, err := strconv.ParseInt(string(text), 10, 64)
		return TernCast(err != nil, res, uint64(0)), err
	case string:
		return string(text), nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", val)
	}
}

// MarshalJSON implements json.Marshaler.
func (n *Null[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return nullBytes, nil
	}
	var asAny any = n.Val
	switch val := asAny.(type) {
	case json.Marshaler:
		return val.MarshalJSON()
	case bool:
		return Tern(val, trueBytes, falseBytes), nil
	case byte:
		return []byte{'"', val, '"'}, nil
	// Skipping uint8 as it collides with byte
	case float32, float64, int, int8, int16, int32, int64, uint, uint16, uint32, uint64:
		return formatNumber(asAny)
	default:
		return json.Marshal(val)
	}
}

// MarshalText implements encoding.TextMarshaler.
func (n *Null[T]) MarshalText() ([]byte, error) {
	if !n.Valid {
		return nullBytes, nil
	}
	var asAny any = n.Val
	switch val := asAny.(type) {
	case encoding.TextMarshaler:
		return val.MarshalText()
	case bool:
		return Tern(val, trueBytes, falseBytes), nil
	case byte:
		return []byte{val}, nil
	case []byte:
		return val, nil
	// Skipping uint8 as it collides with byte
	case float32, float64, int, int8, int16, int32, int64, uint, uint16, uint32, uint64:
		return formatNumber(asAny)
	case string:
		return []byte(val), nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", val)
	}
}

func formatNumber(val any) ([]byte, error) {
	switch num := val.(type) {
	case float32:
		return []byte(strconv.FormatFloat(float64(num), 'f', -1, 32)), nil
	case float64:
		return []byte(strconv.FormatFloat(num, 'f', -1, 64)), nil
	case int:
		return []byte(strconv.FormatInt(int64(num), 10)), nil
	case int8:
		return []byte(strconv.FormatInt(int64(num), 10)), nil
	case int16:
		return []byte(strconv.FormatInt(int64(num), 10)), nil
	case int32:
		return []byte(strconv.FormatInt(int64(num), 10)), nil
	case int64:
		return []byte(strconv.FormatInt(num, 10)), nil
	case uint:
		return []byte(strconv.FormatUint(uint64(num), 10)), nil
	case uint8:
		return []byte(strconv.FormatUint(uint64(num), 10)), nil
	case uint16:
		return []byte(strconv.FormatUint(uint64(num), 10)), nil
	case uint32:
		return []byte(strconv.FormatUint(uint64(num), 10)), nil
	case uint64:
		return []byte(strconv.FormatUint(num, 10)), nil
	default:
		return nil, fmt.Errorf("format number: unsupported type: %T", val)
	}
}

// SetValid changes the inner value and also sets it to be non-null.
func (n *Null[T]) SetValid(value T) {
	n.Val = value
	n.Valid = true
	n.Set = true
}

// Ptr returns a pointer to the inner value, or a nil pointer if this value
// is null.
func (n Null[T]) Ptr() *T {
	if !n.Valid {
		return nil
	}
	return &n.Val
}

// IsZero returns true for null or zero values.
func (n Null[T]) IsZero() bool {
	return !n.Valid
}

// Scan implements sql.Scanner.
func (n *Null[T]) Scan(value any) error {
	if value == nil {
		n.Val = Zero[T]()
		n.Valid = false
		n.Set = false
		return nil
	}
	var asAny any = n.Val
	switch asAny.(type) {
	case byte:
		str := value.(string)
		if len(str) == 0 {
			n.Val = Zero[T]()
			n.Valid = false
			n.Set = false
			return nil
		}
		n.Val = any(str[0]).(T)
		n.Valid = true
		n.Set = true
		return nil
	case time.Time:
		switch value.(type) {
		case time.Time:
			n.Valid = true
			n.Set = true
			n.Val = value.(T)
			return nil
		default:
			return fmt.Errorf("null: cannot scan type %T into null time: %v", value, value)
		}
	// default includes: bool, []byte, float32, float64,
	//                   int, int8, int16, int32, int64,
	//                   string, uint, uint16, uint32, uint64:
	default:
		n.Valid = true
		n.Set = true
		return convert.Assign(&n.Val, value)
	}
}

// Value implements driver.Value.
func (n Null[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	var asAny any = n.Val
	switch val := asAny.(type) {
	case driver.Valuer:
		return val.Value()
	case byte:
		return []byte{val}, nil
	case float32:
		return float64(val), nil
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	// Skipping uint8 as it collides with byte
	case uint:
		return uint64(val), nil
	case uint16:
		return uint64(val), nil
	case uint32:
		return uint64(val), nil
	// default includes: bool, []byte, float64, int64, string, time.Time, uint64
	default:
		return val, nil
	}
}
