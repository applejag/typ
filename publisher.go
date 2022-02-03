// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"errors"
	"sync"
	"time"
)

// Errors specific for the listener and subscriptions.
var (
	ErrAlreadyUnsubscribed       = errors.New("already unsubscribed")
	ErrSubscriptionNotInitalized = errors.New("subscription is not initialized")
)

// Publisher is a type that allows publishing an event which will be sent out
// to all subscribed channels. A sort of "fan-out message queue".
type Publisher[T any] struct {
	OnPubTimeout    func(ev T)    // called if Pub or PubWait times out
	PubTimeoutAfter time.Duration // times out Pub & PubWait, if positive

	subs  []chan T
	mutex sync.RWMutex
}

// Pub sends the event to all subscriptions in their own goroutines and returns
// immediately without waiting for any of the channels to finish sending.
func (o *Publisher[T]) Pub(ev T) {
	o.mutex.RLock()
	for _, sub := range o.subs {
		go o.send(ev, sub, o.PubTimeoutAfter, o.OnPubTimeout)
	}
	o.mutex.RUnlock()
}

// PubSlice sends a slice of events to all subscriptions in their own goroutines
// and returns immediately without waiting for any of the channels to finish
// sending.
func (o *Publisher[T]) PubSlice(evs []T) {
	o.mutex.RLock()
	for _, ev := range evs {
		for _, sub := range o.subs {
			go o.send(ev, sub, o.PubTimeoutAfter, o.OnPubTimeout)
		}
	}
	o.mutex.RUnlock()
}

// PubWait blocks while sending the event to all subscriptions in their own
// goroutines, and waits until all have received the message or timed out.
func (o *Publisher[T]) PubWait(ev T) {
	var wg sync.WaitGroup
	o.mutex.RLock()
	wg.Add(len(o.subs))
	for _, sub := range o.subs {
		go o.sendWaitGroup(ev, sub, o.PubTimeoutAfter, o.OnPubTimeout, &wg)
	}
	o.mutex.RUnlock()
	wg.Wait()
}

// PubSliceWait blocks while sending a slice of events to all subscriptions in
// their own goroutines, and waits until all have received the message or
// timed out.
func (o *Publisher[T]) PubSliceWait(evs []T) {
	var wg sync.WaitGroup
	o.mutex.RLock()
	wg.Add(len(o.subs) * len(evs))
	for _, ev := range evs {
		for _, sub := range o.subs {
			go o.sendWaitGroup(ev, sub, o.PubTimeoutAfter, o.OnPubTimeout, &wg)
		}
	}
	o.mutex.RUnlock()
	wg.Wait()
}

// PubSync blocks while sending the event syncronously to all subscriptions
// without starting a single goroutine. Useful in performance-critical use cases
// where there are a low expected number of subscribers (0-3).
func (o *Publisher[T]) PubSync(ev T) {
	o.mutex.RLock()
	for _, sub := range o.subs {
		o.send(ev, sub, o.PubTimeoutAfter, o.OnPubTimeout)
	}
	o.mutex.RUnlock()
}

// PubSliceSync blocks while sending a slice of events syncronously to all
// subscriptions without starting a single goroutine. Useful in
// performance-critical use cases where there are a low expected number of
// subscribers (0-3).
func (o *Publisher[T]) PubSliceSync(evs []T) {
	o.mutex.RLock()
	for _, ev := range evs {
		for _, sub := range o.subs {
			o.send(ev, sub, o.PubTimeoutAfter, o.OnPubTimeout)
		}
	}
	o.mutex.RUnlock()
}

func (o *Publisher[T]) send(ev T, sub chan T, timeout time.Duration, onTimeout func(T)) {
	if !SendTimeout(sub, ev, timeout) && onTimeout != nil {
		onTimeout(ev)
	}
}

func (o *Publisher[T]) sendWaitGroup(ev T, sub chan T, timeout time.Duration, onTimeout func(T), wg *sync.WaitGroup) {
	o.send(ev, sub, timeout, onTimeout)
	wg.Done()
}

// Sub subscribes to events in a newly created channel with no buffer.
func (o *Publisher[T]) Sub() <-chan T {
	o.mutex.Lock()
	sub := make(chan T)
	o.subs = append(o.subs, sub)
	o.mutex.Unlock()
	return sub
}

// SubBuf subscribes to events in a newly created channel with a specified
// buffer size.
func (o *Publisher[T]) SubBuf(size int) <-chan T {
	o.mutex.Lock()
	sub := make(chan T, size)
	o.subs = append(o.subs, sub)
	o.mutex.Unlock()
	return sub
}

// Unsub unsubscribes a previously subscribed channel.
func (o *Publisher[T]) Unsub(sub <-chan T) error {
	if sub == nil {
		return ErrSubscriptionNotInitalized
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()
	idx := o.subIndex(sub)
	if idx == -1 {
		return ErrAlreadyUnsubscribed
	}
	close(o.subs[idx])
	o.subs = append(o.subs[:idx], o.subs[idx+1:]...)
	return nil
}

// UnsubAll unsubscribes all subscription channels, rendering them all useless.
func (o *Publisher[T]) UnsubAll() error {
	o.mutex.Lock()
	for _, ch := range o.subs {
		close(ch)
	}
	o.subs = nil
	o.mutex.Unlock()
	return nil
}

func (o *Publisher[T]) subIndex(sub <-chan T) int {
	for i, ch := range o.subs {
		if ch == sub {
			return i
		}
	}
	return -1
}
