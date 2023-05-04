// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// SPDX-FileCopyrightText: 2009 The Go Authors
//
// SPDX-License-Identifier: BSD-3-Clause

package sync2_test

import (
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"testing/quick"

	"gopkg.in/typ.v4/sync2"
)

type mapOp string

const (
	opLoad          = mapOp("Load")
	opStore         = mapOp("Store")
	opLoadOrStore   = mapOp("LoadOrStore")
	opLoadAndDelete = mapOp("LoadAndDelete")
	opDelete        = mapOp("Delete")
)

var mapOps = [...]mapOp{opLoad, opStore, opLoadOrStore, opLoadAndDelete, opDelete}

// mapCall is a quick.Generator for calls on mapInterface.
type mapCall struct {
	op mapOp
	k  string
	v  string
}

func (c mapCall) apply(m mapInterface[string, string]) (string, bool) {
	switch c.op {
	case opLoad:
		return m.Load(c.k)
	case opStore:
		m.Store(c.k, c.v)
		return "", false
	case opLoadOrStore:
		return m.LoadOrStore(c.k, c.v)
	case opLoadAndDelete:
		return m.LoadAndDelete(c.k)
	case opDelete:
		m.Delete(c.k)
		return "", false
	default:
		panic("invalid mapOp")
	}
}

type mapResult[V any] struct {
	value V
	ok    bool
}

func randValue(r *rand.Rand) string {
	b := make([]byte, r.Intn(4))
	for i := range b {
		b[i] = 'a' + byte(rand.Intn(26))
	}
	return string(b)
}

func (mapCall) Generate(r *rand.Rand, size int) reflect.Value {
	c := mapCall{op: mapOps[rand.Intn(len(mapOps))], k: randValue(r)}
	switch c.op {
	case opStore, opLoadOrStore:
		c.v = randValue(r)
	}
	return reflect.ValueOf(c)
}

func applyCalls(m mapInterface[string, string], calls []mapCall) (results []mapResult[string], final map[string]string) {
	for _, c := range calls {
		v, ok := c.apply(m)
		results = append(results, mapResult[string]{v, ok})
	}

	final = make(map[string]string)
	m.Range(func(k, v string) bool {
		final[k] = v
		return true
	})

	return results, final
}

func applyMap(calls []mapCall) ([]mapResult[string], map[string]string) {
	return applyCalls(new(sync2.Map[string, string]), calls)
}

func applyRWMutexMap(calls []mapCall) ([]mapResult[string], map[string]string) {
	return applyCalls(new(RWMutexMap[string, string]), calls)
}

func applyDeepCopyMap(calls []mapCall) ([]mapResult[string], map[string]string) {
	return applyCalls(new(DeepCopyMap[string, string]), calls)
}

func TestMapMatchesRWMutex(t *testing.T) {
	if err := quick.CheckEqual(applyMap, applyRWMutexMap, nil); err != nil {
		t.Error(err)
	}
}

func TestMapMatchesDeepCopy(t *testing.T) {
	if err := quick.CheckEqual(applyMap, applyDeepCopyMap, nil); err != nil {
		t.Error(err)
	}
}

func TestConcurrentRange(t *testing.T) {
	const mapSize = 1 << 10

	m := new(sync2.Map[int64, int64])
	for n := int64(1); n <= mapSize; n++ {
		m.Store(n, n)
	}

	done := make(chan struct{})
	var wg sync.WaitGroup
	defer func() {
		close(done)
		wg.Wait()
	}()
	for g := int64(runtime.GOMAXPROCS(0)); g > 0; g-- {
		r := rand.New(rand.NewSource(g))
		wg.Add(1)
		go func(g int64) {
			defer wg.Done()
			for i := int64(0); ; i++ {
				select {
				case <-done:
					return
				default:
				}
				for n := int64(1); n < mapSize; n++ {
					if r.Int63n(mapSize) == 0 {
						m.Store(n, n*i*g)
					} else {
						m.Load(n)
					}
				}
			}
		}(g)
	}

	iters := 1 << 10
	if testing.Short() {
		iters = 16
	}
	for n := iters; n > 0; n-- {
		seen := make(map[int64]bool, mapSize)

		m.Range(func(k, v int64) bool {
			if v%k != 0 {
				t.Fatalf("while Storing multiples of %v, Range saw value %v", k, v)
			}
			if seen[k] {
				t.Fatalf("Range visited key %v twice", k)
			}
			seen[k] = true
			return true
		})

		if len(seen) != mapSize {
			t.Fatalf("Range visited %v elements of %v-element Map", len(seen), mapSize)
		}
	}
}

func TestIssue40999(t *testing.T) {
	var m sync2.Map[*int, string]

	// Since the miss-counting in missLocked (via Delete)
	// compares the miss count with len(m.dirty),
	// add an initial entry to bias len(m.dirty) above the miss count.
	m.Store(nil, "")

	var finalized uint32

	// Set finalizers that count for collected keys. A non-zero count
	// indicates that keys have not been leaked.
	for atomic.LoadUint32(&finalized) == 0 {
		p := new(int)
		runtime.SetFinalizer(p, func(*int) {
			atomic.AddUint32(&finalized, 1)
		})
		m.Store(p, "foo")
		m.Delete(p)
		runtime.GC()
	}
}

func TestConcurrentLen(t *testing.T) {
	const mapSize = 1 << 10

	m := new(sync2.Map[int64, int64])
	for n := int64(1); n <= mapSize; n++ {
		m.Store(n, n)
	}

	done := make(chan struct{})
	var wg sync.WaitGroup
	defer func() {
		close(done)
		wg.Wait()
	}()
	for g := int64(runtime.GOMAXPROCS(0)); g > 0; g-- {
		r := rand.New(rand.NewSource(g))
		wg.Add(1)
		go func(g int64) {
			defer wg.Done()
			for i := int64(0); ; i++ {
				select {
				case <-done:
					return
				default:
				}
				for n := int64(1); n < mapSize; n++ {
					if r.Int63n(mapSize) == 0 {
						m.Store(n, n*i*g)
					} else {
						m.Load(n)
					}
				}
			}
		}(g)
	}

	iters := 1 << 10
	if testing.Short() {
		iters = 16
	}
	for n := iters; n > 0; n-- {
		length := m.Len()

		if length != mapSize {
			t.Fatalf("Len returned %v length of %v-element Map", length, mapSize)
		}
	}
}

func TestNestedLen(t *testing.T) {
	// Ensure the Map.Len does not cause deadlock, due to possible
	// nested mutex locking using the same mutex.
	var m sync2.Map[int64, int64]
	m.Store(1, 123)
	m.Store(2, 351)
	m.Store(3, 519)

	m.Range(func(key, value int64) bool {
		t.Logf("key: %d, value: %d, len: %d", key, value, m.Len())
		return true
	})
}
