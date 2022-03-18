// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sync2

import (
	"sync/atomic"

	"gopkg.in/typ.v3"
)

// AtomicValue is a wrapper around sync/atomic.Value from the Go stdlib.
//
// An AtomicValue must not be copied after first use.
type AtomicValue[T any] struct {
	atom atomic.Value
}

// CompareAndSwap executes the compare-and-swap operation for the AtomicValue.
//
// Using nil as the new value will result in panic.
func (v *AtomicValue[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.atom.CompareAndSwap(old, new)
}

// Load returns the value set by the most recent call to Store, or the zero value
// for this type if there has been no call to Store.
func (v *AtomicValue[T]) Load() (val T) {
	x := v.atom.Load()
	if x == nil {
		return typ.Zero[T]()
	}
	return x.(T)
}

// Store sets the value of the AtomicValue.
//
// Using nil as the new value will result in panic.
func (v *AtomicValue[T]) Store(val T) {
	v.atom.Store(val)
}

// Swap stores the new value into the AtomicValue and returns the previous value.
// It returns a zero value if the AtomicValue is empty.
//
// Using nil as the new value will result in panic.
func (v *AtomicValue[T]) Swap(new T) (old T) {
	x := v.atom.Swap(new)
	if x == nil {
		return typ.Zero[T]()
	}
	return x.(T)
}
