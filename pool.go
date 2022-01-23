// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import "sync"

// Pool is a wrapper around sync.Pool to allow generic and type safe access.
type Pool[T any] struct {
	pool sync.Pool
	New  func() T
}

// Get attempts to obtain an arbitrary item from the Pool, remove it from the
// Pool, and return it to the caller. At any stage it may fail and will instead
// use the Pool.New function to create a new item and return that instead.
func (p *Pool[T]) Get() T {
	if p.New == nil {
		var x T
		return x
	}
	p.pool.New = func() any { return p.New() }
	return p.pool.Get().(T)
}

// Put adds x to the pool.
func (p *Pool[T]) Put(x T) {
	p.pool.Put(x)
}
