// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// SPDX-FileCopyrightText: 2009 The Go Authors
//
// SPDX-License-Identifier: BSD-3-Clause

package lists

import (
	"fmt"
	"testing"
)

// For debugging - keep around.
func dumpRing[T any](r *Ring[T]) {
	if r == nil {
		fmt.Println("empty")
		return
	}
	i, n := 0, r.Len()
	for p := r; i < n; p = p.next {
		fmt.Printf("%4d: %p = {<- %p | %p ->}\n", i, p, p.prev, p.next)
		i++
	}
	fmt.Println()
}

func verifyRing(t *testing.T, r *Ring[int], N int, sum int) {
	// Len
	n := r.Len()
	if n != N {
		t.Errorf("r.Len() == %d; expected %d", n, N)
	}

	// iteration
	n = 0
	s := 0
	r.Do(func(p int) {
		n++
		s += p
	})
	if n != N {
		t.Errorf("number of forward iterations == %d; expected %d", n, N)
	}
	if sum >= 0 && s != sum {
		t.Errorf("forward ring sum = %d; expected %d", s, sum)
	}

	if r == nil {
		return
	}

	// connections
	if r.next != nil {
		var p *Ring[int] // previous element
		for q := r; p == nil || q != r; q = q.next {
			if p != nil && p != q.prev {
				t.Errorf("prev = %p, expected q.prev = %p\n", p, q.prev)
			}
			p = q
		}
		if p != r.prev {
			t.Errorf("prev = %p, expected r.prev = %p\n", p, r.prev)
		}
	}

	// Next, Prev
	if r.Next() != r.next {
		t.Errorf("r.Next() != r.next")
	}
	if r.Prev() != r.prev {
		t.Errorf("r.Prev() != r.prev")
	}

	// Move
	if r.Move(0) != r {
		t.Errorf("r.Move(0) != r")
	}
	if r.Move(N) != r {
		t.Errorf("r.Move(%d) != r", N)
	}
	if r.Move(-N) != r {
		t.Errorf("r.Move(%d) != r", -N)
	}
	for i := 0; i < 10; i++ {
		ni := N + i
		mi := ni % N
		if r.Move(ni) != r.Move(mi) {
			t.Errorf("r.Move(%d) != r.Move(%d)", ni, mi)
		}
		if r.Move(-ni) != r.Move(-mi) {
			t.Errorf("r.Move(%d) != r.Move(%d)", -ni, -mi)
		}
	}
}

func TestCornerCases(t *testing.T) {
	var (
		r0 *Ring[int]
		r1 Ring[int]
	)
	// Basics
	verifyRing(t, r0, 0, 0)
	verifyRing(t, &r1, 1, 0)
	// Insert
	r1.Link(r0)
	verifyRing(t, r0, 0, 0)
	verifyRing(t, &r1, 1, 0)
	// Insert
	r1.Link(r0)
	verifyRing(t, r0, 0, 0)
	verifyRing(t, &r1, 1, 0)
	// Unlink
	r1.Unlink(0)
	verifyRing(t, &r1, 1, 0)
}

func makeRingN(n int) *Ring[int] {
	r := NewRing[int](n)
	for i := 1; i <= n; i++ {
		r.Value = i
		r = r.Next()
	}
	return r
}

func sumRingN(n int) int { return (n*n + n) / 2 }

func TestNew(t *testing.T) {
	for i := 0; i < 10; i++ {
		r := NewRing[int](i)
		verifyRing(t, r, i, -1)
	}
	for i := 0; i < 10; i++ {
		r := makeRingN(i)
		verifyRing(t, r, i, sumRingN(i))
	}
}

func TestLink1(t *testing.T) {
	r1a := makeRingN(1)
	var r1b Ring[int]
	r2a := r1a.Link(&r1b)
	verifyRing(t, r2a, 2, 1)
	if r2a != r1a {
		t.Errorf("a) 2-element link failed")
	}

	r2b := r2a.Link(r2a.Next())
	verifyRing(t, r2b, 2, 1)
	if r2b != r2a.Next() {
		t.Errorf("b) 2-element link failed")
	}

	r1c := r2b.Link(r2b)
	verifyRing(t, r1c, 1, 1)
	verifyRing(t, r2b, 1, 0)
}

func TestLink2(t *testing.T) {
	var r0 *Ring[int]
	r1a := &Ring[int]{Value: 42}
	r1b := &Ring[int]{Value: 77}
	r10 := makeRingN(10)

	r1a.Link(r0)
	verifyRing(t, r1a, 1, 42)

	r1a.Link(r1b)
	verifyRing(t, r1a, 2, 42+77)

	r10.Link(r0)
	verifyRing(t, r10, 10, sumRingN(10))

	r10.Link(r1a)
	verifyRing(t, r10, 12, sumRingN(10)+42+77)
}

func TestLink3(t *testing.T) {
	var r Ring[int]
	n := 1
	for i := 1; i < 10; i++ {
		n += i
		verifyRing(t, r.Link(NewRing[int](i)), n, -1)
	}
}

func TestUnlink(t *testing.T) {
	r10 := makeRingN(10)
	s10 := r10.Move(6)

	sum10 := sumRingN(10)

	verifyRing(t, r10, 10, sum10)
	verifyRing(t, s10, 10, sum10)

	r0 := r10.Unlink(0)
	verifyRing(t, r0, 0, 0)

	r1 := r10.Unlink(1)
	verifyRing(t, r1, 1, 2)
	verifyRing(t, r10, 9, sum10-2)

	r9 := r10.Unlink(9)
	verifyRing(t, r9, 9, sum10-2)
	verifyRing(t, r10, 9, sum10-2)
}

func TestLinkUnlink(t *testing.T) {
	for i := 1; i < 4; i++ {
		ri := NewRing[int](i)
		for j := 0; j < i; j++ {
			rj := ri.Unlink(j)
			verifyRing(t, rj, j, -1)
			verifyRing(t, ri, i-j, -1)
			ri.Link(rj)
			verifyRing(t, ri, i, -1)
		}
	}
}

// Test that calling Move() on an empty Ring initializes it.
func TestMoveEmptyRing(t *testing.T) {
	var r Ring[int]

	r.Move(1)
	verifyRing(t, &r, 1, 0)
}
