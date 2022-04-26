package sync2

import (
	"sync"

	"gopkg.in/typ.v3/pkg/sets"
)

// NewSet creates a new empty sync set.
func NewSet[T comparable]() Set[T] {
	return set[T]{
		s:  make(sets.Set[T]),
		mu: &sync.RWMutex{},
	}
}

// WrapSet creates a new SyncSet that uses the provided set as its underlying
// set. I.e., modifications to the SyncSet/the underlying set will be reflected
// in the other.
func WrapSet[T comparable](toWrap sets.Set[T]) Set[T] {
	return set[T]{
		s:  toWrap,
		mu: &sync.RWMutex{},
	}
}

// Set is an interface for thread-safe sets.
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
	// mathmatics it's denoted as:
	// 	A ∩ B
	// Example:
	// 	{1 2 3} ∩ {3 4 5} = {3}
	// This operation is commutative, meaning you will get the same result no matter
	// the order of the operands. In other words:
	// 	A.Intersect(B) == B.Intersect(A)
	Intersect(set Set[T]) Set[T]
	// Union performs a "union" on the sets and returns a new set.
	// A union is a set of all elements that appear in either set. In mathmatics
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
	// elements that appears in the first set but not in the second. In mathmatics
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
	// the sets, but not both. In mathmatics it's commonly denoted as either:
	// 	A △ B
	// or
	// 	A ⊖ B
	// Example:
	// 	{1 2 3} ⊖ {3 4 5} = {1 2 4 5}
	// This operation is commutative, meaning you will get the same result no matter
	// the order of the operands. In other words:
	// 	A.SymDiff(B) == B.SymDiff(A)
	SymDiff(set Set[T]) Set[T]

	lock()
	rLock()
	unlock()
	rUnlock()
	underlying() sets.Set[T]
}

// Set is a wrapper for gopkg.in/typ.v3/pkg/sets.Set that is safe to use in
// a multithreaded environment.
type set[T comparable] struct {
	s  sets.Set[T]
	mu *sync.RWMutex
}

func (s set[T]) String() string {
	s.rLock()
	str := s.s.String()
	s.rUnlock()
	return str
}

func (s set[T]) Has(value T) bool {
	s.rLock()
	ok := s.s.Has(value)
	s.rUnlock()
	return ok
}

func (s set[T]) Add(value T) bool {
	s.lock()
	ok := s.s.Add(value)
	s.unlock()
	return ok
}

func (s set[T]) AddSet(set Set[T]) int {
	set.rLock()
	s.lock()
	numAdded := s.s.AddSet(set.underlying())
	s.unlock()
	set.rUnlock()
	return numAdded
}

func (s set[T]) Remove(value T) bool {
	s.lock()
	ok := s.s.Remove(value)
	s.unlock()
	return ok
}

func (s set[T]) RemoveSet(set Set[T]) int {
	set.rLock()
	s.lock()
	numRemoved := s.s.RemoveSet(set.underlying())
	s.unlock()
	set.rUnlock()
	return numRemoved
}

func (s set[T]) Clone() Set[T] {
	s.rLock()
	clone := WrapSet(s.s.Clone())
	s.rUnlock()
	return clone
}

func (s set[T]) Slice() []T {
	s.rLock()
	slice := s.s.Slice()
	s.rUnlock()
	return slice
}

func (s set[T]) Intersect(other Set[T]) Set[T] {
	s.rLock()
	other.rLock()
	intersection := s.s.Intersect(other.underlying())
	other.rUnlock()
	s.rUnlock()
	return set[T]{
		s:  intersection,
		mu: &sync.RWMutex{},
	}
}

func (s set[T]) Union(other Set[T]) Set[T] {
	s.rLock()
	other.rLock()
	union := s.s.Union(other.underlying())
	other.rUnlock()
	s.rUnlock()
	return set[T]{
		s:  union,
		mu: &sync.RWMutex{},
	}
}

func (s set[T]) SetDiff(other Set[T]) Set[T] {
	s.rLock()
	other.rLock()
	setDiff := s.s.SetDiff(other.underlying())
	other.rUnlock()
	s.rUnlock()
	return set[T]{
		s:  setDiff,
		mu: &sync.RWMutex{},
	}
}

func (s set[T]) SymDiff(other Set[T]) Set[T] {
	s.rLock()
	other.rLock()
	union := s.s.SymDiff(other.underlying())
	other.rUnlock()
	s.rUnlock()
	return &set[T]{
		s:  union,
		mu: &sync.RWMutex{},
	}
}

func (s set[T]) lock() {
	s.mu.Lock()
}

func (s set[T]) rLock() {
	s.mu.RLock()
}

func (s set[T]) unlock() {
	s.mu.Unlock()
}

func (s set[T]) rUnlock() {
	s.mu.RUnlock()
}

func (s set[T]) underlying() sets.Set[T] {
	return s.s
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
func CartesianProduct[TA comparable, TB comparable](a Set[TA], b Set[TB]) Set[Product[TA, TB]] {
	result := make(sets.Set[Product[TA, TB]])
	a.rLock()
	b.rLock()
	for valueA := range a.underlying() {
		for valueB := range b.underlying() {
			result.Add(Product[TA, TB]{valueA, valueB})
		}
	}
	b.rUnlock()
	a.rUnlock()
	return WrapSet(result)
}

// Product is the resulting type from a Cartesian product operation.
type Product[TA comparable, TB comparable] sets.Product[TA, TB]
