// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2014 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// The toolbox is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// Alternatively, the CookieJar toolbox may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).

package utility

import "math"

// A generic curve type for the utility library. Input will be normalized to
// [0, 1] prior to passing it to the curve. Outputs will be clamped to [0, 1].
type Curve func(float64) float64

// Creates a linear function defined as y = ax + b.
func Linear(a, b float64) Curve {
	return func(x float64) float64 {
		return a*x + b
	}
}

// Creates an exponential curve, specializing whether it's convex.
func Exponential(center, exp float64, convex bool) Curve {
	if convex {
		return func(x float64) float64 {
			return 1 - math.Pow(math.Abs(x-center), exp)
		}
	} else {
		return func(x float64) float64 {
			return math.Pow(math.Abs(x-center), exp)
		}
	}
}

// Creates a sigmoid threshold curve, specializing whether it's increasing.
func Logistic(center float64, inc bool) Curve {
	return func(x float64) float64 {
		return 1 / (1 + math.Exp(10*(center-x)))
	}
}
