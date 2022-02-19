// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"context"
	"time"
)

// SendTimeout sends a value to a channel, or cancels after a given duration.
func SendTimeout[C SendChan[V], V any](ch C, value V, timeout time.Duration) bool {
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
func SendContext[C SendChan[V], V any](ctx context.Context, ch C, value V) bool {
	select {
	case ch <- value:
		return true
	case <-ctx.Done():
		return false
	}
}

// RecvTimeout receives a value from a channel, or cancels after a given timeout.
// If the timeout duration is zero or negative, then no limit is used.
func RecvTimeout[C RecvChan[V], V any](ch C, timeout time.Duration) (V, bool) {
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
		return Zero[V](), false
	}
}

// RecvContext receives a value from a channel, or cancels when the given
// context is cancelled.
func RecvContext[C ~<-chan V, V any](ctx context.Context, ch C) (V, bool) {
	select {
	case value, ok := <-ch:
		return value, ok
	case <-ctx.Done():
		return Zero[V](), false
	}
}

// RecvQueued will receive all values from a channel until either there's no
// more values in the channel's queue buffer, or it has received maxValues
// values, or until the channel is closed, whichever comes first.
func RecvQueued[C RecvChan[V], V any](ch C, maxValues int) []V {
	var buffer []V
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
func RecvQueuedFull[C RecvChan[V], B ~[]V, V any](ch C, buf B) int {
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
