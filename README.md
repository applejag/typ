<!--
SPDX-FileCopyrightText: 2022 Kalle Fagerberg

SPDX-License-Identifier: CC-BY-4.0
-->

# go-typ

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/6b0289f204c044c2911a53c67a4833d9)](https://app.codacy.com/gh/go-typ/typ?utm_source=github.com&utm_medium=referral&utm_content=go-typ/typ&utm_campaign=Badge_Grade_Settings)
[![REUSE status](https://api.reuse.software/badge/github.com/go-typ/typ)](https://api.reuse.software/info/github.com/go-typ/typ)
[![Go Reference](https://pkg.go.dev/badge/gopkg.in/typ.v4.svg)](https://pkg.go.dev/gopkg.in/typ.v4)

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

Requires Go v1.18rc1 or later as the code makes heavy use of generics.

## Installation and usage

```sh
go get -u gopkg.in/typ.v4
```

```go
import (
	"fmt"

	"gopkg.in/typ.v4/avl"
	"gopkg.in/typ.v4/sets"
)

func UsingSets() {
	set1 := make(sets.Set[string])
	set1.Add("A")
	set1.Add("B")
	set1.Add("C")
	fmt.Println("set1:", set1) // {A B C}

	set2 := make(sets.Set[string])
	set2.Add("B")
	set2.Add("C")
	set2.Add("D")
	fmt.Println("set2:", set2) // {B C D}

	fmt.Println("union:", set1.Union(set2))         // {A B C D}
	fmt.Println("intersect:", set1.Intersect(set2)) // {B C}
	fmt.Println("set diff:", set1.SetDiff(set2))    // {A}
	fmt.Println("sym diff:", set1.SymDiff(set2))    // {A D}
}

func UsingAVLTree() {
	tree := avl.NewOrdered[string]()

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

<!-- lint disable maximum-line-length -->

- `gopkg.in/typ.v4/arrays`:

  - `arrays.Array2D[T]`: 2-dimensional array.

- `gopkg.in/typ.v4/sync2`:

  - `sync2.AtomicValue[T]`: Atomic value store, wrapper around [`sync/atomic.Value`](https://pkg.go.dev/sync/atomic#Value).
  - `sync2.KeyedMutex[T]`: Mutual exclusive lock on a per-key basis.
  - `sync2.KeyedRWMutex[T]`: Mutual exclusive reader/writer lock on a per-key basis.
  - `sync2.Map[K,V]`: Concurrent map, forked from [`sync.Map`](https://pkg.go.dev/sync#Map).
  - `sync2.Once1[R1]`: Run action once, and tracks return values, wrapper around [`sync.Once`](https://pkg.go.dev/sync#Once).
  - `sync2.Once2[R1,R2]`: Run action once, and tracks return values, wrapper around [`sync.Once`](https://pkg.go.dev/sync#Once).
  - `sync2.Once3[R1,R2,R3]`: Run action once, and tracks return values, wrapper around [`sync.Once`](https://pkg.go.dev/sync#Once).
  - `sync2.Pool[T]`: Object pool, wrapper around [`sync.Pool`](https://pkg.go.dev/sync#Pool).

- `gopkg.in/typ.v4/lists`:

  - `lists.List[T]`: Linked list, forked from [`container/list`](https://pkg.go.dev/container/list).
  - `lists.Queue[T]`: First-in-first-out collection.
  - `lists.Ring[T]`: Circular list, forked from [`container/ring`](https://pkg.go.dev/container/ring).
  - `lists.Stack[T]`: First-in-last-out collection.

- `gopkg.in/typ.v4/avl`:

  - `avl.Tree[T]`: AVL-tree (auto-balancing binary search tree) implementation.

- `gopkg.in/typ.v4/chans`:

  - `chans.PubSub[T]`: Publish-subscribe pattern using channels.

- `gopkg.in/typ.v4/maps`:

  - `maps.Bimap[K,V]`: Bi-directional map.

- `gopkg.in/typ.v4/sets`:

  - `sets.Set[T]`: Set, based on set theory.

- `gopkg.in/typ.v4/slices`:

  - `slices.Sorted[T]`: Always-sorted slice. Requires custom `less` function.

<!-- lint enable maximum-line-length -->

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

- <https://github.com/golang-design/go2generics>: 🧪 A chunk of experiments and
  demos about Go 2 generics design (type parameter & type set)

Official Go packages:

- [`golang.org/x/exp/constraints`](https://pkg.go.dev/golang.org/x/exp/constraints):
  Constraints that are useful for generic code, such as constraints.Ordered.

- [`golang.org/x/exp/maps`](https://pkg.go.dev/golang.org/x/exp/maps):
  A collection of generic functions that operate on slices of any element type.

- [`golang.org/x/exp/slices`](https://pkg.go.dev/golang.org/x/exp/slices):
  A collection of generic functions that operate on maps of any key or element
  type.

## License

This project is primarily licensed under the MIT license:

- My Go code in this project is licensed under the MIT license:
  [LICENSES/MIT.txt](LICENSES/MIT.txt)

- Some Go code in this project is forked from Go's source code, which is
  licensed under the 3-Clause BSD license: [LICENSES/BSD-3-Clause.txt](LICENSES/BSD-3-Clause.txt)

- Documentation is licensed under the Creative Commons Attribution 4.0
  International (CC-BY-4.0) license: [LICENSES](LICENSES/CC-BY-4.0.txt)

- Miscellanious files are licensed under the Creative Commons Zero Universal
  license (CC0-1.0): [LICENSES](LICENSES/CC0-1.0.txt)

- GitHub Action for REUSE linting (and not any of go-typ's code) is licensed
  under GNU General Public License 3.0 or later (GPL-3.0-or-later):
  [LICENSES/GPL-3.0-or-later.txt](LICENSES/GPL-3.0-or-later.txt)

Copyright &copy; Kalle Fagerberg
