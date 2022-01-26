// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import "constraints"

// Real is a type constraint for any real numbers. That being integers or floats.
type Real interface {
	constraints.Integer | constraints.Float
}

// Number is a type constraint for any Go numbers, including complex numbers.
type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}
