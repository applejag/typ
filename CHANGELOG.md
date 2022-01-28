<!--
SPDX-FileCopyrightText: 2022 Kalle Fagerberg

SPDX-License-Identifier: CC-BY-4.0
-->

<!-- lint disable maximum-line-length -->

# go-typ changelog

This project tries to follow [SemVer 2.0.0](https://semver.org/).

## v1.2.0 (WIP)

- Added utility functions:

  - `typ.ContainsFunc[T]([]T, T, func(T, T) bool) bool`: Checks if value exists in slice with custom equals. (e94faf7)
  - `typ.DistinctFunc[T]([]T, func(T, T) bool) []T`: Returns new slice of unique elements with custom equals. (e94faf7)
  - `typ.TrimFunc[T]([]T, func(T) bool) []T`: Trim away unwanted matches from start and end. (e94faf7)
  - `typ.TrimLeftFunc[T]([]T, func(T) bool) []T`: Trim away unwanted matches from start. (e94faf7)
  - `typ.TrimRightFunc[T]([]T, func(T) bool) []T`: Trim away unwanted matches from end. (e94faf7)
  - `typ.IndexFunc[T]([]T, func(T) bool) int`: Returns index of a value, or -1 if not found. (a69bf35)
  - `typ.Last[T]([]T) T`: Returns the last item in a slice. (9bc54ab)
  - `typ.Any[T]([]T, func(T) bool) bool`: Does condition match any value? (024361a)
  - `typ.All[T]([]T, func(T) bool) bool`: Does condition match all values? (024361a)
  - `typ.Map[TA, TB](slice []TA, func(TA) TB) []TB`: Returns converted slice. (024361a)
  - `typ.MapErr[TA, TB](slice []TA, func(TA) (TB, error)) ([]TB, error)`: Returns converted slice, or first error. (024361a)
  - `typ.Filter[T](slice []T, func(T) bool) []T`: Returns filtered slice. (024361a)
  - `typ.Fold[TState, T]([]T, TState, func(TState, T) TState) TState`: Accumulate values from slice. (024361a)
  - `typ.Concat[T]([]T, []T) []T`: Returns two concatenated slices. (024361a)

- Changed `typ.Set.Set()` and `.Unset()` to `.Add()` and `.Remove()`, respectively. (bfabb2d)

## v1.1.1 (2022-01-26)

- Added example tests for `typ.Set` and `typ.OrderedTree`. (e15e311)

## v1.1.0 (2022-01-26)

- Added utility functions:

  - `typ.Index[T]([]T, T) int`: Returns index of a value, or -1 if not found. (f104746)
  - `typ.MakeChanOfChan[T](chan T, ...int) chan T`: Returns the result of `make(chan T)`, useful for anonymous types. (f104746)
  - `typ.MakeChanOf[T](T, ...int) chan T`: Returns the result of `make(chan T)`, useful for anonymous types. (f104746)
  - `typ.MakeMapOfMap[K,V](map[K]V, ...int) map[K]V`: Returns the result of `make(map[K]V)`, useful for anonymous types. (f104746)
  - `typ.MakeMapOf[K,V](K, V, ...int) map[K]V`: Returns the result of `make(map[K]V)`, useful for anonymous types. (f104746)
  - `typ.MakeSliceOfKey[K,V](map[K]V, ...int) []K`: Returns the result of `make([]K)`, useful for anonymous types. (f104746)
  - `typ.MakeSliceOfSlice[T]([]T, ...int) []T`: Returns the result of `make([]T)`, useful for anonymous types. (f104746)
  - `typ.MakeSliceOfValue[K,V](map[K]V, ...int) []V`: Returns the result of `make([]V)`, useful for anonymous types. (f104746)
  - `typ.MakeSliceOf[T](T, ...int) []T`: Returns the result of `make([]T)`, useful for anonymous types. (f104746)
  - `typ.NewOf[T](*T) *T`: Returns the result of `new(T)`, useful for anonymous types. (f104746)
  - `typ.SafeGetOr[T]([]T, int, T) T`: Index a slice, or return fallback value if index is out of bounds. (f104746)
  - `typ.SafeGet[T]([]T, int) T`: Index a slice, or return zero if index is out of bounds. (f104746)
  - `typ.TryGet[T]([]T, int) (T, bool)`: Index a slice, or return false if index is out of bounds. (f104746)
  - `typ.ZeroOf[T](T) T`: Returns the zero value for an anonymous type. (f104746)

## v1.0.1 (2022-01-25)

- Fixed package reference in docs and tests. (2e1eb32)

- Added utility functions:

  - `typ.TrimLeft[T]([]T, []T)`: Trim away unwanted elements from start. (2286b5c)
  - `typ.TrimRight[T]([]T, []T)`: Trim away unwanted elements from end. (2286b5c)
  - `typ.Trim[T]([]T, []T)`: Trim away unwanted elements from start and end. (2286b5c)

## v1.0.0 (2022-01-25)

- Added types:

  - `typ.Null[T]`: Nullable type without needing pointers, forked from [`github.com/volatiletech/null/v9`](https://github.com/volatiletech/null) (#22)

- Added utility functions:

  - `typ.Coal[T](...T) T`: Coalesce operator, returns first non-zero value. (#20)
  - `typ.IsNil[T](T) bool`: Returns true if the generic value is nil. (#22)
  - `typ.TernCast[T](bool, any, T) T`: Conditionally cast a value. (#22)
  - `typ.Tern[T](bool, T, T) T`: Ternary operator, return based on conditional. (#20)

## v0.1.0 (2022-01-23)

- Added types:

  - `typ.AtomicValue[T]`: Atomic value store, wrapper around [`sync/atomic.Value`](https://pkg.go.dev/sync/atomic#Value).
  - `typ.List[T]`: Linked list, forked from [`container/list`](https://pkg.go.dev/container/list).
  - `typ.Number`: Type constraint for any number: integers, floats, & complex.
  - `typ.Pool[T]`: Object pool, wrapper around [`sync.Pool`](https://pkg.go.dev/sync#Pool).
  - `typ.Publisher[T]`: Publish-subscribe pattern (pubsub) using channels.
  - `typ.Real`: Type constraint for real numbers: integers & floats.
  - `typ.Ring[T]`: Circular list, forked from [`container/ring`](https://pkg.go.dev/container/ring).
  - `typ.Set[T]`: Set, based on set theory.
  - `typ.Stack[T]`: First-in-last-out collection.
  - `typ.SyncMap[K,V]`: Concurrent map, forked from [`sync.Map`](https://pkg.go.dev/sync#Map).
  - `typ.Tree[T]`: AVL-tree (auto-balancing binary search tree) implementation.

- Added utility functions:

  - `typ.Clamp01[T](T) T`: Clamp a value between `0` and `1`.
  - `typ.Clamp[T](T, T, T) T`: Clamp a value inside a range.
  - `typ.ContainsValue[K, V](map[K]V, V) bool`: Does map contain value?
  - `typ.Contains[T]([]T, T) bool`: Does slice contain value?
  - `typ.Max[T](...T) T`: Return the largest value.
  - `typ.Min[T](...T) T`: Return the smallest value.
  - `typ.Product[T](...T) T`: Multiplies together numbers.
  - `typ.RecvTimeout[T](chan<- T, time.Duration)`: Receive from channel with timeout.
  - `typ.Reverse[T]([]T)`: Reverse the order of a slice.
  - `typ.Search[T]([]T, T)`: Searches for element index or insertion index in slice.
  - `typ.SendTimeout[T](<-chan T, T, time.Duration)`: Send to channel with timeout.
  - `typ.ShuffleRand[T]([]T, *rand.Rand)`: Randomizes the order of a slice.
  - `typ.Shuffle[T]([]T)`: Randomizes the order of a slice.
  - `typ.SortDesc[T]([]T)`: Sort ordered slices in descending order.
  - `typ.Sort[T]([]T)`: Sort ordered slices in ascending order.
  - `typ.Sum[T](...T) T`: Sums up numbers (addition).
  - `typ.Zero[T]()`: Returns the zero value for a type.
