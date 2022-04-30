// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sets

// Set is an interface for sets.
type Set[T comparable] interface {
	// String converts this set to its string representation.
	String() string
	// Has returns true if the value exists in the set.
	Has(value T) bool
	// Add will add an element to the set, and return true if it was added
	// or false if the value already existed in the set.
	Add(value T) bool
	// AddSet will add all element found in specified set to this set, and
	// return the number of values that was added.
	AddSet(set Set[T]) int
	// Remove will remove an element from the set, and return true if it was removed
	// or false if no such value existed in the set.
	Remove(value T) bool
	// RemoveSet will remove all element found in specified set from this set, and
	// return the number of values that was removed.
	RemoveSet(set Set[T]) int
	// Clone returns a copy of the set.
	Clone() Set[T]
	// Slice returns a new slice of all values in the set.
	Slice() []T
	// Intersect performs an "intersection" on the sets and returns a new set.
	// An intersection is a set of all elements that appear in both sets. In
	// mathematics it's denoted as:
	// 	A ∩ B
	// Example:
	// 	{1 2 3} ∩ {3 4 5} = {3}
	// This operation is commutative, meaning you will get the same result no matter
	// the order of the operands. In other words:
	// 	A.Intersect(B) == B.Intersect(A)
	Intersect(set Set[T]) Set[T]
	// Union performs a "union" on the sets and returns a new set.
	// A union is a set of all elements that appear in either set. In mathematics
	// it's denoted as:
	// 	A ∪ B
	// Example:
	// 	{1 2 3} ∪ {3 4 5} = {1 2 3 4 5}
	// This operation is commutative, meaning you will get the same result no matter
	// the order of the operands. In other words:
	// 	A.Union(B) == B.Union(A)
	Union(set Set[T]) Set[T]
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
	SetDiff(set Set[T]) Set[T]
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
	SymDiff(set Set[T]) Set[T]
	// Range calls f sequentially for each value present in the set.
	// If f returns false, range stops the iteration.
	//
	// Order is not guaranteed to be the same between executions.
	//
	// Methods that modify the set should not be used in the passed in function,
	// as it will cause a deadlock.
	Range(f func(value T) bool)
}

// CartesianProduct performs a "Cartesian product" on two sets and returns a
// new set. A Cartesian product of two sets is a set of all possible combinations
// between two sets. In mathematics it's denoted as:
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
func CartesianProduct[TA, TB comparable](a Set[TA], b Set[TB]) map[Product[TA, TB]]struct{} {
	result := make(map[Product[TA, TB]]struct{}, 0)
	a.Range(func(valueA TA) bool {
		b.Range(func(valueB TB) bool {
			result[Product[TA, TB]{valueA, valueB}] = struct{}{}
			return true
		})
		return true
	})
	return result
}

// Product is the resulting type from a Cartesian product operation.
type Product[TA, TB any] struct {
	A TA
	B TB
}
