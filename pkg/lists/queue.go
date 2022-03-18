// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package lists

import "gopkg.in/typ.v3"

// Queue is a first-in-first-out collection.
//
// The implementation is done via a linked-list.
type Queue[T any] struct {
	list List[T]
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return q.list.Len()
}

// Enqueue adds a value to the start of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.list.PushFront(value)
}

// Dequeue removes and returns a value from the end of the queue.
func (q *Queue[T]) Dequeue() (T, bool) {
	elem := q.list.Back()
	if elem == nil {
		return typ.Zero[T](), false
	}
	return q.list.Remove(elem), true
}

// Peek returns (but does not remove) a value from the end of the queue.
func (q *Queue[T]) Peek() (T, bool) {
	elem := q.list.Back()
	if elem == nil {
		return typ.Zero[T](), false
	}
	return elem.Value, true
}
