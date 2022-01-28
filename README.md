<!--
SPDX-FileCopyrightText: 2022 Kalle Fagerberg

SPDX-License-Identifier: CC-BY-4.0
-->

# go-typ

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/6b0289f204c044c2911a53c67a4833d9)](https://app.codacy.com/gh/go-typ/typ?utm_source=github.com&utm_medium=referral&utm_content=go-typ/typ&utm_campaign=Badge_Grade_Settings)
[![REUSE status](https://api.reuse.software/badge/github.com/go-typ/typ)](https://api.reuse.software/info/github.com/go-typ/typ)
[![Go Reference](https://pkg.go.dev/badge/gopkg.in/typ.v1.svg)](https://pkg.go.dev/gopkg.in/typ.v1)

Generic types and functions that are missing from Go, including sets, trees,
linked lists, etc.

All code is implemented with 0 dependencies and in pure Go code (no CGo).

## Background

Go v1.18 is about to be released now in February 2022, and with it comes some
features that has been talked about for a really long time. One of which being
**generics!** [(Go 1.18 beta release notes)](https://tip.golang.org/doc/go1.18)

They have moved generics from the Go v2.0 milestone over to Go v1.18, which
means they have to stay backwards compatible and cannot alter any existing
types. On top of this, they do not seem to plan on releasing any generic data
types in the Go standard library until Go v1.19. All in all, to use generic
data types with Go v1.18, you'll have to either write your own, or use a
third-party package, like this one :)

This repository includes those generic functions and types that I find are
missing from the release of Go v1.18-beta1, as well as a number of other
data structures and utility functions I think should've been included in the
standard library a long time ago. But now with generics, we can finally have
sensible implementations of sets, trees, stacks, etc without excessive casting.

## Compatibility

Requires Go v1.18beta1 or later as the code makes use of generics.

## Installation and usage

```sh
go get -u gopkg.in/typ.v1
```

```go
func UsingSets() {
	set1 := make(typ.Set[string])
	set1.Add("A")
	set1.Add("B")
	set1.Add("C")
	fmt.Println("set1:", set1) // {A B C}

	set2 := make(typ.Set[string])
	set2.Add("B")
	set2.Add("C")
	set2.Add("D")
	fmt.Println("set2:", set2) // {B C D}

	fmt.Println("union:", set1.Union(set2))         // {A B C D}
	fmt.Println("intersect:", set1.Intersect(set2)) // {B C}
	fmt.Println("set diff:", set1.SetDiff(set2))    // {A}
	fmt.Println("sym diff:", set1.SymDiff(set2))    // {A D}
}

func UsingOrderedTree() {
	var tree typ.OrderedTree[string]

	// Unordered input
	tree.Add("E")
	tree.Add("B")
	tree.Add("D")
	tree.Add("C")
	tree.Add("A")

	// Sorted output
	fmt.Println(tree.Len(), tree) // 5 [A B C D E]
}
```

## Features
### Types

- `typ.AtomicValue[T]`: Atomic value store, wrapper around [`sync/atomic.Value`](https://pkg.go.dev/sync/atomic#Value).
- `typ.List[T]`: Linked list, forked from [`container/list`](https://pkg.go.dev/container/list).
- `typ.Null[T]`: Nullable type without needing pointers, forked from [`github.com/volatiletech/null/v9`](https://github.com/volatiletech/null)
- `typ.Pool[T]`: Object pool, wrapper around [`sync.Pool`](https://pkg.go.dev/sync#Pool).
- `typ.Publisher[T]`: Publish-subscribe pattern (pubsub) using channels.
- `typ.Queue[T]`: First-in-first-out collection.
- `typ.Ring[T]`: Circular list, forked from [`container/ring`](https://pkg.go.dev/container/ring).
- `typ.Set[T]`: Set, based on set theory.
- `typ.Stack[T]`: First-in-last-out collection.
- `typ.SyncMap[K,V]`: Concurrent map, forked from [`sync.Map`](https://pkg.go.dev/sync#Map).
- `typ.Tree[T]`: AVL-tree (auto-balancing binary search tree) implementation.

> Explanation:
>
> - Forked type: Copied their code and modified it so it uses generic types down
>   to the backing struct layer. This benefits the most from generics support.
>
> - Wrapped type: Code depends on the underlying non-generic type, and adds
>   abstraction to hide the type casting. Less performant than full generic
>   support, but is done to reduce excessive complexity in this repository.
>
> - Neither forked nor wrapped: Original code written by yours truly.

### Constraints

- `typ.Number`: Type constraint for any number: integers, floats, & complex.
- `typ.Real`: Type constraint for real numbers: integers & floats.

### Utility functions

<!-- lint disable maximum-line-length -->

- `typ.All[T]([]T, func(T) bool) bool`: Does condition match all values?
- `typ.Any[T]([]T, func(T) bool) bool`: Does condition match any value?
- `typ.ChunkIter[T]([]T, int) [][]T`: Invoke callback for all chunks in a slice.
- `typ.Chunk[T]([]T, int) [][]T`: Divide up a slice.
- `typ.Clamp01[T](T) T`: Clamp a value between `0` and `1`.
- `typ.Clamp[T](T, T, T) T`: Clamp a value inside a range.
- `typ.Coal[T](...T) T`: Coalesce operator, returns first non-zero value.
- `typ.ContainsFunc[T]([]T, T, func(T, T) bool) bool`: Checks if value exists in slice with custom equals.
- `typ.ContainsValue[K, V](map[K]V, V) bool`: Does map contain value?
- `typ.Contains[T]([]T, T) bool`: Does slice contain value?
- `typ.CountBy[K, V]([]V, func(V) K) []Counting[K]`: Count elements by key.
- `typ.DistinctFunc[T]([]T, func(T, T) bool) []T`: Returns new slice of unique elements with custom equals.
- `typ.Distinct[T]([]T, func(T, T) bool) []T`: Returns new slice of unique elements.
- `typ.ExceptSet[T]([]T, Set[T]) []T`: Exclude values from other set.
- `typ.Except[T]([]T, []T) []T`: Exclude values from other slice.
- `typ.Fill[T]([]T, T)`: Fill a slice with a value.
- `typ.Filter[T](slice []T, func(T) bool) []T`: Returns filtered slice.
- `typ.FoldReverse[TState, T]([]T, TState, func(TState, T) TState) TState`: Accumulate values from slice in reverse order.
- `typ.Fold[TState, T]([]T, TState, func(TState, T) TState) TState`: Accumulate values from slice.
- `typ.GroupBy[K, V]([]V, func(V) K) []Grouping[K, V]`: Group elements by key.
- `typ.IndexFunc[T]([]T, func(T) bool) int`: Returns index of a value, or -1 if not found.
- `typ.Index[T]([]T, T) int`: Returns index of a value, or -1 if not found.
- `typ.IsNil[T](T) bool`: Returns true if the generic value is nil.
- `typ.Last[T]([]T) T`: Returns the last item in a slice.
- `typ.MakeChanOfChan[T](chan T, ...int) chan T`: Returns the result of `make(chan T)`, useful for anonymous types.
- `typ.MakeChanOf[T](T, ...int) chan T`: Returns the result of `make(chan T)`, useful for anonymous types.
- `typ.MakeMapOfMap[K,V](map[K]V, ...int) map[K]V`: Returns the result of `make(map[K]V)`, useful for anonymous types.
- `typ.MakeMapOf[K,V](K, V, ...int) map[K]V`: Returns the result of `make(map[K]V)`, useful for anonymous types.
- `typ.MakeSliceOfKey[K,V](map[K]V, ...int) []K`: Returns the result of `make([]K)`, useful for anonymous types.
- `typ.MakeSliceOfSlice[T]([]T, ...int) []T`: Returns the result of `make([]T)`, useful for anonymous types.
- `typ.MakeSliceOfValue[K,V](map[K]V, ...int) []V`: Returns the result of `make([]V)`, useful for anonymous types.
- `typ.MakeSliceOf[T](T, ...int) []T`: Returns the result of `make([]T)`, useful for anonymous types.
- `typ.MapErr[TA, TB](slice []TA, func(TA) (TB, error)) ([]TB, error)`: Returns converted slice, or first error.
- `typ.Map[TA, TB](slice []TA, func(TA) TB) []TB`: Returns converted slice.
- `typ.Max[T](...T) T`: Return the largest value.
- `typ.Min[T](...T) T`: Return the smallest value.
- `typ.NewOf[T](*T) *T`: Returns the result of `new(T)`, useful for anonymous types.
- `typ.PairsIter[T]([]T, func(T, T))`: Invoke callback for all pairs in a slice.
- `typ.Pairs[T]([]T) [][2]T`: Returns all pairs from a slice.
- `typ.Product[T](...T) T`: Multiplies together numbers.
- `typ.RecvTimeout[T](chan<- T, time.Duration)`: Receive from channel with timeout.
- `typ.Reverse[T]([]T)`: Reverse the order of a slice.
- `typ.SafeGetOr[T]([]T, int, T) T`: Index a slice, or return fallback value if index is out of bounds.
- `typ.SafeGet[T]([]T, int) T`: Index a slice, or return zero if index is out of bounds.
- `typ.Search[T]([]T, T)`: Searches for element index or insertion index in slice.
- `typ.SendTimeout[T](<-chan T, T, time.Duration)`: Send to channel with timeout.
- `typ.ShuffleRand[T]([]T, *rand.Rand)`: Randomizes the order of a slice.
- `typ.Shuffle[T]([]T)`: Randomizes the order of a slice.
- `typ.SortDesc[T]([]T)`: Sort ordered slices in descending order.
- `typ.Sort[T]([]T)`: Sort ordered slices in ascending order.
- `typ.Sum[T](...T) T`: Sums up numbers (addition).
- `typ.TernCast[T](bool, any, T) T`: Conditionally cast a value.
- `typ.Tern[T](bool, T, T) T`: Ternary operator, return based on conditional.
- `typ.TrimFunc[T]([]T, func(T) bool) []T`: Trim away unwanted matches from start and end.
- `typ.TrimLeftFunc[T]([]T, func(T) bool) []T`: Trim away unwanted matches from start.
- `typ.TrimLeft[T]([]T, []T)`: Trim away unwanted elements from start.
- `typ.TrimRightFunc[T]([]T, func(T) bool) []T`: Trim away unwanted matches from end.
- `typ.TrimRight[T]([]T, []T)`: Trim away unwanted elements from end.
- `typ.Trim[T]([]T, []T)`: Trim away unwanted elements from start and end.
- `typ.TryGet[T]([]T, int) (T, bool)`: Index a slice, or return false if index is out of bounds.
- `typ.WindowedIter[T]([]T, int, func([]T))`: Invoke callback for all windows in a slice.
- `typ.Windowed[T]([]T, int) [][]T`: Returns all windows from a slice.
- `typ.ZeroOf[T](T) T`: Returns the zero value for an anonymous type.
- `typ.Zero[T]() T`: Returns the zero value for a type.

<!-- lint enable maximum-line-length -->

## Development

Please read the [CONTRIBUTING.md](CONTRIBUTING.md) for information about
development environment and guidelines.

## Similar projects

All the below include multiple data structure implementations each, all with
Go 1.18 generics support.

- <https://github.com/zyedidia/generic>: An experimental collection of generic
  data structures written in Go.

- <https://github.com/marstr/collection>: Generic Golang implementation of a few
  basic data structures.

- <https://github.com/Kytabyte/go2-generic-containers>: Some container data
  structures written in the next generation of Golang with generics.

- <https://github.com/tomyl/collection>: Generic data structures for Go.

- <https://github.com/gabstv/container>: Generic data structures now that
  Go 1.18+ supports generics.

- <https://github.com/go-generics/collections>: Go generic collections

## License

This project is primarily licensed under the MIT license:

- My Go code in this project is licensed under the MIT license:
  [LICENSES/MIT.txt](LICENSES/MIT.txt)

- Some Go code in this project is forked from Go's source code, which is
  licensed under the 3-Clause BSD license: [LICENSES/BSD-3-Clause.txt](LICENSES/BSD-3-Clause.txt)

- Some Go code in this project is forked from Volatile Tech's source code
  (<https://github.com/volatiletech/null>), which is licensed under the
  2-Clause BSD license: [LICENSES/BSD-2-Clause.txt](LICENSES/BSD-2-Clause.txt)

- Documentation is licensed under the Creative Commons Attribution 4.0
  International (CC-BY-4.0) license: [LICENSES](LICENSES/CC-BY-4.0.txt)

- Miscellanious files are licensed under the Creative Commons Zero Universal
  license (CC0-1.0): [LICENSES](LICENSES/CC0-1.0.txt)

- GitHub Action for REUSE linting (and not any of go-typ's code) is licensed
  under GNU General Public License 3.0 or later (GPL-3.0-or-later):
  [LICENSES/GPL-3.0-or-later.txt](LICENSES/GPL-3.0-or-later.txt)

Copyright &copy; Kalle Fagerberg
