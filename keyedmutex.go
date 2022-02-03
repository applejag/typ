// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import "sync"

// KeyedLocker represents an object that can be locked and unlocked on a
// per-key basis.
type KeyedLocker[T comparable] interface {
	LockKey(key T)
	UnlockKey(key T)
}

// KeyedMutex is a keyed mutual exclusion lock.
// The zero value for a KeyedMutex is an unlocked mutex for any key.
//
// The KeyedMutex does not clear its own cache, so it will continue to grow
// for every new key that is used, unless you call the ClearKey method.
//
// A KeyedMutex must not be copied after first use.
type KeyedMutex[T comparable] struct {
	m SyncMap[T, *sync.Mutex]
}

func (km *KeyedMutex[T]) LockKey(key T) {
	m, _ := km.m.LoadOrStore(key, &sync.Mutex{})
	m.Lock()
}

func (km *KeyedMutex[T]) TryLockKey(key T) bool {
	m, _ := km.m.LoadOrStore(key, &sync.Mutex{})
	return m.TryLock()
}

func (km *KeyedMutex[T]) UnlockKey(key T) {
	m, _ := km.m.LoadOrStore(key, &sync.Mutex{})
	m.Unlock()
}

func (km *KeyedMutex[T]) ClearKey(key T) {
	km.m.Delete(key)
}

// KeyedRWMutex is a keyed reader/writer mutual exclusion lock. The lock can be
// held on a per-unique-key basis by an arbitrary number of readers or a single
// writer. The zero value for a KeyedRWMutex is an unlocked mutex for any key.
//
// The KeyedRWMutex does not clear its own cache, so it will continue to grow
// for every new key that is used, unless you call the ClearKey method.
//
// A KeyedRWMutex must not be copied after first use.
type KeyedRWMutex[T comparable] struct {
	m SyncMap[T, *sync.RWMutex]
}

func (km *KeyedRWMutex[T]) LockKey(key T) {
	m, _ := km.m.LoadOrStore(key, &sync.RWMutex{})
	m.Lock()
}

func (km *KeyedRWMutex[T]) TryLockKey(key T) bool {
	m, _ := km.m.LoadOrStore(key, &sync.RWMutex{})
	return m.TryLock()
}

func (km *KeyedRWMutex[T]) UnlockKey(key T) {
	m, _ := km.m.LoadOrStore(key, &sync.RWMutex{})
	m.Unlock()
}

func (km *KeyedRWMutex[T]) RLockKey(key T) {
	m, _ := km.m.LoadOrStore(key, &sync.RWMutex{})
	m.RLock()
}

func (km *KeyedRWMutex[T]) TryRLockKey(key T) bool {
	m, _ := km.m.LoadOrStore(key, &sync.RWMutex{})
	return m.TryRLock()
}

func (km *KeyedRWMutex[T]) RUnlockKey(key T) {
	m, _ := km.m.LoadOrStore(key, &sync.RWMutex{})
	m.RUnlock()
}

func (km *KeyedRWMutex[T]) ClearKey(key T) {
	km.m.Delete(key)
}
