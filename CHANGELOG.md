<!--
SPDX-FileCopyrightText: 2022 Kalle Fagerberg

SPDX-License-Identifier: CC-BY-4.0
-->

<!-- lint disable maximum-line-length -->

# go-typ changelog

This project tries to follow [SemVer 2.0.0](https://semver.org/).

## v2.0.0 (2022-02-19)

- Added constraints from [`golang.org/x/exp/constraints`](https://pkg.go.dev/golang.org/x/exp/constraints)
  as they were removed from the stdlib. (729e08d)

- Added utility functions:

  - `typ.ClearMap[K, V](map[K]V)`: Removes all elements from a map. (ce84398)
  - `typ.CloneMap[K, V](map[K]V) map[K]V`: Returns a shallow copy of a map. (a4f42ca)
  - `typ.CloneSlice[T]([]T) []T`: Returns a shallow copy of a slice. (a4f42ca)
  - `typ.DerefZero[T](*T) T`: Returns a dereferenced pointer value, or zero if nil. (10de9a6)
  - `typ.GrowSlice[T]([]T, int) []T`: Adds `n` number of zero elements to a slice. (814016c)
  - `typ.SortDescFunc[T]([]T, func(T, T) bool)`: Sorts a slice with a given sort function in descending order. (b55962b)
  - `typ.SortFunc[T]([]T, func(T, T) bool)`: Sorts a slice with a given sort function. (b55962b)
  - `typ.SortStableDescFunc[T]([]T, func(T, T) bool)`: Sorts a slice with a given sort function in descending order, while keeping original order of equal elements. (b55962b)
  - `typ.SortStableFunc[T]([]T, func(T, T) bool)`: Sorts a slice with a given sort function, while keeping original order of equal elements. (b55962b)

- Renamed `typ.Ptr` to `typ.Ref`. (10de9a6)

- Renamed `typ.OrderedTree` to `typ.AVLTree`. (d96902f)

- Changed `typ.AVLTree` to not be constrainted on only ordered types. Now you
  can use `typ.NewAVLTree()` to create one with a custom comparator. (d96902f)

- Changed most functions to use generic constraints as `[S ~[]E, E any]` instead
  of just `[S []any]`. (b973cb7)

## v1.3.0 (2022-02-03)

- Added types:

  - `typ.KeyedMutex[T]`: Mutual exclusive lock on a per-key basis. (4f99f8e)
  - `typ.KeyedRWMutex[T]`: Mutual exclusive reader/writer lock on a per-key basis. (4f99f8e)

- Added utility functions:

  - `typ.Abs[T](T) T`: Absolute value of a number. (f6f0cdf)
  - `typ.Digits10Sign[T](T) int`: Number of base 10 digits (including sign) in integer. (36fbfef)
  - `typ.Digits10[T](T) int`: Number of base 10 digits (excluding sign) in integer. (36fbfef)
  - `typ.RecvContext[T](context.Context, <-chan T) (T, bool)`: Receive from a channel, or cancel with context. (1bfa4b7)
  - `typ.RecvQueuedFull[T](<-chan T, []T)`: Receive all queued values from a channel's buffer. (a56b0e5)
  - `typ.RecvQueued[T](<-chan T, int) []T`: Receive all queued values from a channel's buffer. (a56b0e5)
  - `typ.SendContext[T](context.Context, chan<- T, T) bool`: Send to a channel, or cancel with context. (1bfa4b7)

## v1.2.0 (2022-01-29)

- Added types:

  - `typ.Array2D[T]`: 2-dimensional array. (74289fe)
  - `typ.OrderedSlice[T]`: Always-sorted slice for ordered types. (08b8720)
  - `typ.SortedSlice[T]`: Always-sorted slice. Requires custom `less` function. (08b8720)

- Added utility functions:

  - `typ.All[T]([]T, func(T) bool) bool`: Does condition match all values? (024361a)
  - `typ.Any[T]([]T, func(T) bool) bool`: Does condition match any value? (024361a)
  - `typ.ChunkIter[T]([]T, int) [][]T`: Invoke callback for all chunks in a slice. (e269dec)
  - `typ.Chunk[T]([]T, int) [][]T`: Divide up a slice. (e269dec)
  - `typ.Concat[T]([]T, []T) []T`: Returns two concatenated slices. (024361a)
  - `typ.ContainsFunc[T]([]T, T, func(T, T) bool) bool`: Checks if value exists in slice with custom equals. (e94faf7)
  - `typ.CountBy[K, V]([]V, func(V) K) []Counting[K]`: Count elements by key. (2105841)
  - `typ.DistinctFunc[T]([]T, func(T, T) bool) []T`: Returns new slice of unique elements with custom equals. (e94faf7)
  - `typ.ExceptSet[T]([]T, Set[T]) []T`: Exclude values from other set. (6c21e5d)
  - `typ.Except[T]([]T, []T) []T`: Exclude values from other slice. (6c21e5d)
  - `typ.Fill[T]([]T, T)`: Fill a slice with a value. (d202eea)
  - `typ.Filter[T](slice []T, func(T) bool) []T`: Returns filtered slice. (024361a)
  - `typ.FoldReverse[TState, T]([]T, TState, func(TState, T) TState) TState`: Accumulate values from slice in reverse order. (0871a38)
  - `typ.Fold[TState, T]([]T, TState, func(TState, T) TState) TState`: Accumulate values from slice. (024361a)
  - `typ.GroupBy[K, V]([]V, func(V) K) []Grouping[K, V]`: Group elements by key. (8468938)
  - `typ.IndexFunc[T]([]T, func(T) bool) int`: Returns index of a value, or -1 if not found. (a69bf35)
  - `typ.InsertSlice[T](*[]T, int, []T)`: Inserts a slice of values at index. (cb51458, 21bf45f, f9a6c42)
  - `typ.Insert[T](*[]T, int, T)`: Inserts a value at index. (cb51458, 21bf45f, f9a6c42)
  - `typ.Last[T]([]T) T`: Returns the last item in a slice. (9bc54ab)
  - `typ.MapErr[TA, TB](slice []TA, func(TA) (TB, error)) ([]TB, error)`: Returns converted slice, or first error. (024361a)
  - `typ.Map[TA, TB](slice []TA, func(TA) TB) []TB`: Returns converted slice. (024361a)
  - `typ.PairsIter[T]([]T, func(T, T))`: Invoke callback for all pairs in a slice. (92498c7)
  - `typ.Pairs[T]([]T) [][2]T`: Returns all pairs from a slice. (92498c7)
  - `typ.Ptr[T](T) *T`: Return a pointer of the value, such as a literal. (d48dd7f)
  - `typ.RemoveSlice[T](*[]T, int, int)`: Removes a slice of values at index. (cb51458, 21bf45f, f9a6c42)
  - `typ.Remove[T](*[]T, int)`: Removes a value at index. (cb51458, 21bf45f, f9a6c42)
  - `typ.TrimFunc[T]([]T, func(T) bool) []T`: Trim away unwanted matches from start and end. (e94faf7)
  - `typ.TrimLeftFunc[T]([]T, func(T) bool) []T`: Trim away unwanted matches from start. (e94faf7)
  - `typ.TrimRightFunc[T]([]T, func(T) bool) []T`: Trim away unwanted matches from end. (e94faf7)
  - `typ.WindowedIter[T]([]T, int, func([]T))`: Invoke callback for all windows in a slice. (91a701e)
  - `typ.Windowed[T]([]T, int) [][]T`: Returns all windows from a slice. (91a701e)

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
