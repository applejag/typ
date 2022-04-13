// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sync2

import "sync"

// Once1 is an object that will perform exactly one action and stores the
// return values of the initial invokation. Any subsequent calls will reuse
// the same return values from the initial invokation.
//
// The Once1 implementation is a wrapper around the built-in sync.Once.
//
// A Once1 must not be copied after first use.
type Once1[R1 any] struct {
	once sync.Once
	R1   R1
}

func (o *Once1[R1]) Do(f func() R1) R1 {
	o.once.Do(func() {
		o.R1 = f()
	})
	return o.R1
}

// Once2 is an object that will perform exactly one action and stores the
// return values of the initial invokation. Any subsequent calls will reuse
// the same return values from the initial invokation.
//
// The Once2 implementation is a wrapper around the built-in sync.Once.
//
// A Once2 must not be copied after first use.
type Once2[R1, R2 any] struct {
	once sync.Once
	R1   R1
	R2   R2
}

func (o *Once2[R1, R2]) Do(f func() (R1, R2)) (R1, R2) {
	o.once.Do(func() {
		o.R1, o.R2 = f()
	})
	return o.R1, o.R2
}

// Once3 is an object that will perform exactly one action and stores the
// return values of the initial invokation. Any subsequent calls will reuse
// the same return values from the initial invokation.
//
// The Once3 implementation is a wrapper around the built-in sync.Once.
//
// A Once3 must not be copied after first use.
type Once3[R1, R2, R3 any] struct {
	once sync.Once
	R1   R1
	R2   R2
	R3   R3
}

func (o *Once3[R1, R2, R3]) Do(f func() (R1, R2, R3)) (R1, R2, R3) {
	o.once.Do(func() {
		o.R1, o.R2, o.R3 = f()
	})
	return o.R1, o.R2, o.R3
}
