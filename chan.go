// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"context"
	"time"
)

// SendTimeout sends a value to a channel, or cancels after a given duration.
func SendTimeout[T any](ch chan<- T, value T, timeout time.Duration) bool {
	if timeout <= 0 {
		ch <- value
		return true
	}
	timer := time.NewTimer(timeout)
	select {
	case ch <- value:
		timer.Stop()
		return true
	case <-timer.C:
		return false
	}
}

// SendContext receives a value from a channel, or cancels when the given
// context is cancelled.
func SendContext[T any](ctx context.Context, ch chan<- T, value T) bool {
	select {
	case ch <- value:
		return true
	case <-ctx.Done():
		return false
	}
}

// RecvTimeout receives a value from a channel, or cancels after a given timeout.
// If the timeout duration is zero or negative, then no limit is used.
func RecvTimeout[T any](ch <-chan T, timeout time.Duration) (T, bool) {
	if timeout <= 0 {
		value, ok := <-ch
		return value, ok
	}
	timer := time.NewTimer(timeout)
	select {
	case value, ok := <-ch:
		timer.Stop()
		return value, ok
	case <-timer.C:
		return Zero[T](), false
	}
}

// RecvContext receives a value from a channel, or cancels when the given
// context is cancelled.
func RecvContext[T any](ctx context.Context, ch <-chan T) (T, bool) {
	select {
	case value, ok := <-ch:
		return value, ok
	case <-ctx.Done():
		return Zero[T](), false
	}
}

// RecvQueued will receive all values from a channel until either there's no
// more values in the channel's queue buffer, or it has received maxValues
// values, or until the channel is closed, whichever comes first.
func RecvQueued[T any](ch <-chan T, maxValues int) []T {
	var buffer []T
	for len(buffer) < maxValues {
		select {
		case v := <-ch:
			buffer = append(buffer, v)
		default:
			break
		}
	}
	return buffer
}

// RecvQueuedFull will receive all values from a channel until either there's no
// more values in the channel's queue buffer, or it has filled buf with
// values, or until the channel is closed, whichever comes first, and then
// returns the number of values that was received.
func RecvQueuedFull[T any](ch <-chan T, buf []T) int {
	var index int
	for index < len(buf) {
		select {
		case v := <-ch:
			buf[index] = v
		default:
			break
		}
	}
	return index
}
