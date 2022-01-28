// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"fmt"
	"strings"
)

// NewSetOfSlice returns a Set with all values from a slice added to it.
func NewSetOfSlice[T comparable](slice []T) Set[T] {
	var set Set[T]
	for _, v := range slice {
		set.Add(v)
	}
	return set
}

// NewSetOfKeys returns a Set with all keys from a map added to it.
func NewSetOfKeys[K comparable, V any](m map[K]V) Set[K] {
	var set Set[K]
	for k := range m {
		set.Add(k)
	}
	return set
}

// NewSetOfValues returns a Set with all values from a map added to it.
func NewSetOfValues[K comparable, V comparable](m map[K]V) Set[V] {
	var set Set[V]
	for _, v := range m {
		set.Add(v)
	}
	return set
}

// Set holds a collection of values with no duplicates. Its methods are based
// on the mathmatical branch of set theory, and its implementation is using a
// Go map[T]struct{}.
type Set[T comparable] map[T]struct{}

// String converts this set to its string representation.
func (s Set[T]) String() string {
	var sb strings.Builder
	sb.WriteByte('{')
	addDelim := false
	for v := range s {
		if addDelim {
			sb.WriteByte(' ')
		} else {
			addDelim = true
		}
		fmt.Fprint(&sb, v)
	}
	sb.WriteByte('}')
	return sb.String()
}

// Has returns true if the value exists in the set.
func (s Set[T]) Has(value T) bool {
	_, has := s[value]
	return has
}

// Add will add an element to the set, and return true if it was added
// or false if the value already existed in the set.
func (s Set[T]) Add(value T) bool {
	if s.Has(value) {
		return false
	}
	s[value] = struct{}{}
	return true
}

// AddSet will add all element found in specified set to this set, and
// return the number of values that was added.
func (s Set[T]) AddSet(set Set[T]) int {
	var added int
	for v := range set {
		if s.Add(v) {
			added++
		}
	}
	return added
}

// Remove will remove an element from the set, and return true if it was removed
// or false if no such value existed in the set.
func (s Set[T]) Remove(value T) bool {
	if !s.Has(value) {
		return false
	}
	delete(s, value)
	return true
}

// RemoveSet will remove all element found in specified set from this set, and
// return the number of values that was removed.
func (s Set[T]) RemoveSet(set Set[T]) int {
	var removed int
	for v := range set {
		if s.Remove(v) {
			removed++
		}
	}
	return removed
}

// Clone returns a copy of the set.
func (s Set[T]) Clone() Set[T] {
	clone := make(Set[T])
	for v := range s {
		clone.Add(v)
	}
	return clone
}

// Slice returns a new slice of all values in the set.
func (s Set[T]) Slice() []T {
	result := make([]T, 0, len(s))
	for v := range s {
		result = append(result, v)
	}
	return result
}

// Intersect performs an "intersection" on the sets and returns a new set.
// An intersection is a set of all elements that appear in both sets. In
// mathmatics it's denoted as:
// 	A ∩ B
// Example:
// 	{1 2 3} ∩ {3 4 5} = {3}
// This operation is commutative, meaning you will get the same result no matter
// the order of the operands. In other words:
// 	A.Intersect(B) == B.Intersect(A)
func (s Set[T]) Intersect(other Set[T]) Set[T] {
	result := make(Set[T])
	for v := range s {
		if other.Has(v) {
			result.Add(v)
		}
	}
	return result
}

// Union performs a "union" on the sets and returns a new set.
// A union is a set of all elements that appear in either set. In mathmatics
// it's denoted as:
// 	A ∪ B
// Example:
// 	{1 2 3} ∪ {3 4 5} = {1 2 3 4 5}
// This operation is commutative, meaning you will get the same result no matter
// the order of the operands. In other words:
// 	A.Union(B) == B.Union(A)
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := s.Clone()
	for v := range other {
		result.Add(v)
	}
	return result
}

// SetDiff performs a "set difference" on the sets and returns a new set.
// A set difference resembles a subtraction, where the result is a set of all
// elements that appears in the first set but not in the second. In mathmatics
// it's denoted as:
// 	A \ B
// Example:
// 	{1 2 3} \ {3 4 5} = {1 2}
// This operation is noncommutative, meaning you will get different results
// depending on the order of the operands. In other words:
// 	A.SetDiff(B) != B.SetDiff(A)
func (s Set[T]) SetDiff(other Set[T]) Set[T] {
	result := make(Set[T])
	for v := range s {
		if !other.Has(v) {
			result.Add(v)
		}
	}
	return result
}

// SymDiff performs a "symmetric difference" on the sets and returns a new set.
// A symmetric difference is the set of all elements that appear in either of
// the sets, but not both. In mathmatics it's commonly denoted as either:
// 	A △ B
// or
// 	A ⊖ B
// Example:
// 	{1 2 3} ⊖ {3 4 5} = {1 2 4 5}
// This operation is commutative, meaning you will get the same result no matter
// the order of the operands. In other words:
// 	A.SymDiff(B) == B.SymDiff(A)
func (s Set[T]) SymDiff(other Set[T]) Set[T] {
	result := s.SetDiff(other)
	for v := range other {
		if !s.Has(v) {
			result.Add(v)
		}
	}
	return result
}

// CartesianProduct performs a "Cartesian product" on two sets and returns a new
// set. A Cartesian product of two sets is a set of all possible combinations
// between two sets. In mathmatics it's denoted as:
// 	A × B
// Example:
// 	{1 2 3} × {a b c} = {1a 1b 1c 2a 2b 2c 3a 3b 3c}
// This operation is noncommutative, meaning you will get different results
// depending on the order of the operands. In other words:
// 	A.CartesianProduct(B) != B.CartesianProduct(A)
// This noncommutative attribute of the Cartesian product operation is due to
// the pairs being in reverse order if you reverse the order of the operands.
// Example:
// 	{1 2 3} × {a b c} = {1a 1b 1c 2a 2b 2c 3a 3b 3c}
// 	{a b c} × {1 2 3} = {a1 a2 a3 b1 b2 b3 c1 c2 c3}
// 	{1a 1b 1c 2a 2b 2c 3a 3b 3c} != {a1 a2 a3 b1 b2 b3 c1 c2 c3}
func CartesianProduct[TA comparable, TB comparable](a Set[TA], b Set[TB]) Set[SetProduct[TA, TB]] {
	result := make(Set[SetProduct[TA, TB]])
	for valueA := range a {
		for valueB := range b {
			result.Add(SetProduct[TA, TB]{valueA, valueB})
		}
	}
	return result
}

// SetProduct is the resulting type from a Cartesian product operation.
type SetProduct[TA comparable, TB comparable] struct {
	A TA
	B TB
}
