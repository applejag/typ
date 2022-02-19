// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// SPDX-FileCopyrightText: 2009 The Go Authors
//
// SPDX-License-Identifier: BSD-3-Clause

package typ

// Real is a type constraint for any real numbers. That being integers or floats.
type Real interface {
	Integer | Float
}

// Number is a type constraint for any Go numbers, including complex numbers.
type Number interface {
	Integer | Float | Complex
}

// Signed is a constraint that permits any signed integer type.
// If future releases of Go add new predeclared signed integer types,
// this constraint will be modified to include them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
// If future releases of Go add new predeclared integer types,
// this constraint will be modified to include them.
type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
// If future releases of Go add new predeclared floating-point types,
// this constraint will be modified to include them.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type.
// If future releases of Go add new predeclared complex numeric types,
// this constraint will be modified to include them.
type Complex interface {
	~complex64 | ~complex128
}

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	Integer | Float | ~string
}

// RecvChan is a constraint that permits a receive-only chan or a send & receive
// channal.
type RecvChan[T any] interface {
	~chan T | ~<-chan T
}

// SendChan is a constraint that permits a send-only chan or a send & receive
// channal.
type SendChan[T any] interface {
	~chan T | ~chan<- T
}
