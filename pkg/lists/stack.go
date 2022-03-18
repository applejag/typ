// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package lists

import "gopkg.in/typ.v3"

// Stack is a first-in-last-out collection.
type Stack[T any] []T

// Peek returns (but does not remove) the value on top of the stack, or false
// if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if s == nil || len(*s) == 0 {
		return typ.Zero[T](), false
	}
	slice := *s
	lastVal := slice[len(slice)-1]
	return lastVal, true
}

// Pop removes and returns the value on top of the stack, or false if the stack
// is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s == nil || len(*s) == 0 {
		return typ.Zero[T](), false
	}
	slice := *s
	lastIdx := len(slice) - 1
	lastVal := slice[lastIdx]
	*s = slice[:lastIdx]
	return lastVal, true
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}
