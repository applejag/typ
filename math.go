// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"constraints"
)

// Min returns the smallest value.
func Min[T constraints.Ordered](v ...T) T {
	switch len(v) {
	case 0:
		panic("typ.Min: at least one argument is required")
	case 1:
		return v[0]
	default:
		min := v[0]
		for _, v := range v[1:] {
			if v < min {
				min = v
			}
		}
		return min
	}
}

// Max returns the largest value.
func Max[T constraints.Ordered](v ...T) T {
	switch len(v) {
	case 0:
		panic("typ.Max: at least one argument is required")
	case 1:
		return v[0]
	default:
		max := v[0]
		for _, v := range v[1:] {
			if v > max {
				max = v
			}
		}
		return max
	}
}

// Clamp returns the value clamped between the minimum and maximum values.
func Clamp[T constraints.Ordered](v, min, max T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Clamp01 returns the value clamped between 0 (zero) and 1 (one).
func Clamp01[T Real](v T) T {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

// Sum adds upp all numbers from the arguments. Returns 0 if no arguments.
func Sum[T Number](v ...T) T {
	var sum T
	for _, num := range v {
		sum += num
	}
	return sum
}

// Product multiplies together all numbers from the arguments. Returns 1 if no
// arguments.
func Product[T Number](v ...T) T {
	var product T = 1
	for _, num := range v {
		product *= num
	}
	return product
}

// Abs returns the absolute value of a number, in other words removing the sign,
// in other words (again) changing negative numbers to positive and leaving
// positive numbers as-is.
// 	Abs(0)   // => 0
// 	Abs(15)  // => 15
// 	Abs(-15) // => 15
func Abs[T Real](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

// DigitsSign10 returns the number of digits in the number as if it would be
// converted to a string in base 10, plus 1 if the number is negative to account
// for the negative sign. This is computed by comparing its value to all orders
// of 10, making it increadibly faster than calculating logaritms or by
// performing divisions.
func DigitsSign10[T constraints.Integer](v T) int {
	if v < 0 {
		return Digits10(-v) + 1
	}
	return Digits10(v)
}

// Digits10 returns the number of digits in the number as if it would be
// converted to a string in base 10. This is computed by comparing its value
// to all orders of 10, making it increadibly faster than calculating logaritms
// or by performing divisions.
func Digits10[T constraints.Integer](v T) int {
	if v < 0 {
		v = -v
	}
	n := uint64(v)
	switch {
	case n < 10:
		return 1
	case n < 1e2:
		return 2
	case n < 1e3:
		return 3
	case n < 1e4:
		return 4
	case n < 1e5:
		return 5
	case n < 1e6:
		return 6
	case n < 1e7:
		return 7
	case n < 1e8:
		return 8
	case n < 1e9:
		return 9
	case n < 1e10:
		return 10
	case n < 1e11:
		return 11
	case n < 1e12:
		return 12
	case n < 1e13:
		return 13
	case n < 1e14:
		return 14
	case n < 1e15:
		return 15
	case n < 1e16:
		return 16
	case n < 1e17:
		return 17
	case n < 1e18:
		return 18
	case n < 1e19:
		return 19
	default:
		// largest uint64 is 20 digits long.
		// 18446744073709551615 <- max uint64
		// 01234567890123456789
		return 20
	}
}
