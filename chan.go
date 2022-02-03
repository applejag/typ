// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import "time"

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
