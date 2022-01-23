<!--
SPDX-FileCopyrightText: 2022 Kalle Fagerberg

SPDX-License-Identifier: CC-BY-4.0
-->

# go-typ

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/6b0289f204c044c2911a53c67a4833d9)](https://app.codacy.com/gh/go-typ/typ?utm_source=github.com&utm_medium=referral&utm_content=go-typ/typ&utm_campaign=Badge_Grade_Settings)
[![REUSE status](https://api.reuse.software/badge/github.com/go-typ/typ)](https://api.reuse.software/info/github.com/go-typ/typ)

Generic types and functions that are missing from Go, including sets, trees,
linked lists, etc.

All code is implemented with 0 dependencies and in pure Go code (no CGo).

## Compatibility

Requires Go v1.18beta1 or later as the code makes use of generics.

## Installation and usage

```sh
go get -u gopkg.in/typ.v0
```

## Features

### Types

- `typ.List[T]`: Linked list, forked from [`container/list`](https://pkg.go.dev/container/list)
- `typ.Number`: Type constraint for any number: integers, floats, & complex.
- `typ.Real`: Type constraint for real numbers: integers & floats.
- `typ.Ring[T]`: Circular list, forked from [`container/ring`](https://pkg.go.dev/container/ring)
- `typ.Set[T]`: Set, based on set theory
- `typ.Tree[T]`: AVL-tree (auto-balancing binary search tree) implementation

### Utility functions

- `typ.Clamp01[T](v T) T`: Clamp a value between `0` and `1`.
- `typ.Clamp[T](v, min, max T) T`: Clamp a value inside a range.
- `typ.ContainsValue[K, V](m map[K]V, value V) bool`: Does map contain value?
- `typ.Contains[T](slice []T, value T) bool`: Does slice contain value?
- `typ.Max[T](v ...T) T`: Return the largest value.
- `typ.Min[T](v ...T) T`: Return the smallest value.
- `typ.Product[T](v ...T) T`: Multiplies together numbers.
- `typ.Reverse[T](slice []T)`: Reverse the order of a slice.
- `typ.SortDesc[T](slice []T)`: Sort ordered slices in descending order.
- `typ.Sort[T](slice []T)`: Sort ordered slices in ascending order.
- `typ.Sum[T](v ...T) T`: Sums up numbers (addition).

## Development

Please read the [CONTRIBUTING.md](CONTRIBUTING.md) for information about
development environment and guidelines.

## License

This project is primarily licensed under the MIT license:

- My Go code in this project is licensed under the MIT license:
  [LICENSES/MIT.txt](LICENSES/MIT.txt)

- Some Go code in this project is forked from Go's source code, which is
  licensed under the BSD license: [LICENSES/LicenseRef-Go-BSD.txt](LICENSES/LicenseRef-Go-BSD.txt)

- Documentation is licensed under the Creative Commons Attribution 4.0
  International (CC-BY-4.0) license: [LICENSES](LICENSES/CC-BY-4.0.txt)

- Miscellanious files are licensed under the Creative Commons Zero Universal
  license (CC0-1.0): [LICENSES](LICENSES/CC0-1.0.txt)

- GitHub Action for REUSE linting (and not any of go-typ's code) is licensed
  under GNU General Public License 3.0 or later (GPL-3.0-or-later):
  [LICENSES/GPL-3.0-or-later.txt](LICENSES/GPL-3.0-or-later.txt)

Copyright &copy; Kalle Fagerberg
