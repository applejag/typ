// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import "constraints"

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func compare[T constraints.Ordered](a, b T) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}
