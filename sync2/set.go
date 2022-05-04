// SPDX-FileCopyrightText: 2022 Per Alexander Fougner
// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sync2

import (
	"fmt"
	"strings"

	"gopkg.in/typ.v4/sets"
)

// NewSetFromSlice returns a Set with all values from a slice added to it.
func NewSetFromSlice[S ~[]E, E comparable](slice S) sets.Set[E] {
	var set Set[E]
	for _, v := range slice {
		set.Add(v)
	}
	return &set
}

// NewSetFromKeys returns a Set with all keys from a map added to it.
func NewSetFromKeys[M ~map[K]V, K comparable, V any](m M) sets.Set[K] {
	var set Set[K]
	for k := range m {
		set.Add(k)
	}
	return &set
}

// NewSetFromValues returns a Set with all values from a map added to it.
func NewSetFromValues[M ~map[K]V, K comparable, V comparable](m M) sets.Set[V] {
	var set Set[V]
	for _, v := range m {
		set.Add(v)
	}
	return &set
}

// Set holds a collection of values with no duplicates. Its methods are based
// on the mathematical branch of set theory, and its implementation is using a
// Map[T, struct{}].
type Set[T comparable] struct {
	m Map[T, struct{}]
}

// assert that Set implements sets.Set interface.
var _ sets.Set[int] = &Set[int]{}

// String converts this set to its string representation.
func (s *Set[T]) String() string {
	var sb strings.Builder
	sb.WriteByte('{')
	addDelim := false
	s.Range(func(value T) bool {
		if addDelim {
			sb.WriteByte(' ')
		} else {
			addDelim = true
		}
		fmt.Fprint(&sb, value)
		return true
	})
	sb.WriteByte('}')
	return sb.String()
}

// Len returns the number of elements in this set.
func (s *Set[T]) Len() int {
	var count int
	s.Range(func(T) bool {
		count++
		return true
	})
	return count
}

// Has returns true if the value exists in the set.
func (s *Set[T]) Has(value T) bool {
	_, has := s.m.Load(value)
	return has
}

// Add will add an element to the set, and return true if it was added
// or false if the value already existed in the set.
func (s *Set[T]) Add(value T) bool {
	_, loaded := s.m.LoadOrStore(value, struct{}{})
	return !loaded
}

// AddSet will add all element found in specified set to this set, and
// return the number of values that was added.
func (s *Set[T]) AddSet(set sets.Set[T]) int {
	var added int
	set.Range(func(value T) bool {
		if s.Add(value) {
			added++
		}
		return true
	})
	return added
}

// Remove will remove an element from the set, and return true if it was removed
// or false if no such value existed in the set.
func (s *Set[T]) Remove(value T) bool {
	_, loaded := s.m.LoadAndDelete(value)
	return loaded
}

// RemoveSet will remove all element found in specified set from this set, and
// return the number of values that was removed.
func (s *Set[T]) RemoveSet(set sets.Set[T]) int {
	var removed int
	set.Range(func(value T) bool {
		if s.Remove(value) {
			removed++
		}
		return true
	})
	return removed
}

// Clone returns a copy of the set.
func (s *Set[T]) Clone() sets.Set[T] {
	var clone Set[T]
	clone.AddSet(s)
	return &clone
}

// Slice returns a new slice of all values in the set.
func (s *Set[T]) Slice() []T {
	var result []T
	s.m.Range(func(key T, _ struct{}) bool {
		result = append(result, key)
		return true
	})
	return result
}

// Intersect performs an "intersection" on the sets and returns a new set.
// An intersection is a set of all elements that appear in both sets. In
// mathematics it's denoted as:
// 	A ∩ B
// Example:
// 	{1 2 3} ∩ {3 4 5} = {3}
// This operation is commutative, meaning you will get the same result no matter
// the order of the operands. In other words:
// 	A.Intersect(B) == B.Intersect(A)
func (s *Set[T]) Intersect(other sets.Set[T]) sets.Set[T] {
	var result Set[T]
	s.Range(func(value T) bool {
		if other.Has(value) {
			result.Add(value)
		}
		return true
	})
	return &result
}

// Union performs a "union" on the sets and returns a new set.
// A union is a set of all elements that appear in either set. In mathematics
// it's denoted as:
// 	A ∪ B
// Example:
// 	{1 2 3} ∪ {3 4 5} = {1 2 3 4 5}
// This operation is commutative, meaning you will get the same result no matter
// the order of the operands. In other words:
// 	A.Union(B) == B.Union(A)
func (s *Set[T]) Union(other sets.Set[T]) sets.Set[T] {
	result := s.Clone()
	result.AddSet(other)
	return result
}

// SetDiff performs a "set difference" on the sets and returns a new set.
// A set difference resembles a subtraction, where the result is a set of all
// elements that appears in the first set but not in the second. In mathematics
// it's denoted as:
// 	A \ B
// Example:
// 	{1 2 3} \ {3 4 5} = {1 2}
// This operation is noncommutative, meaning you will get different results
// depending on the order of the operands. In other words:
// 	A.SetDiff(B) != B.SetDiff(A)
func (s *Set[T]) SetDiff(other sets.Set[T]) sets.Set[T] {
	var result Set[T]
	s.Range(func(v T) bool {
		if !other.Has(v) {
			result.Add(v)
		}
		return true
	})
	return &result
}

// SymDiff performs a "symmetric difference" on the sets and returns a new set.
// A symmetric difference is the set of all elements that appear in either of
// the sets, but not both. In mathematics it's commonly denoted as either:
// 	A △ B
// or
// 	A ⊖ B
// Example:
// 	{1 2 3} ⊖ {3 4 5} = {1 2 4 5}
// This operation is commutative, meaning you will get the same result no matter
// the order of the operands. In other words:
// 	A.SymDiff(B) == B.SymDiff(A)
func (s *Set[T]) SymDiff(other sets.Set[T]) sets.Set[T] {
	result := s.SetDiff(other)
	other.Range(func(v T) bool {
		if !s.Has(v) {
			result.Add(v)
		}
		return true
	})
	return result
}

// Range calls f sequentially for each value present in the set.
// If f returns false, range stops the iteration.
//
// Order is not guaranteed to be the same between executions.
//
// Methods that modify the set should not be used in the passed in function,
// as it will cause a deadlock.
func (s *Set[T]) Range(f func(value T) bool) {
	s.m.Range(func(v T, _ struct{}) bool {
		return f(v)
	})
}
