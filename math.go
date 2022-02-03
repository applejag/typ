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
