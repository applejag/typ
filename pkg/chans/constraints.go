// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package chans

// Receiver is a constraint that permits a receive-only chan or a send & receive
// channal.
type Receiver[T any] interface {
	~chan T | ~<-chan T
}

// Sender is a constraint that permits a send-only chan or a send & receive
// channal.
type Sender[T any] interface {
	~chan T | ~chan<- T
}

// Chan is a constraint that permits any type of channel, be it a receive-only,
// send-only, or unidirectional channel.
type Chan[T any] interface {
	~chan T | ~chan<- T | ~<-chan T
}
