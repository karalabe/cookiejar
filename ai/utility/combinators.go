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

// A generic curve combinator for the utility library. Input will be normalized
// to [0, 1] prior to passing it to the combinator. Outputs will be clamped to
// [0, 1].
type Combinator func(x, y float64) float64

// Creates an additive curve combinator z = a*x + b*y + c.
func Additive(a, b, c float64) Combinator {
	return func(x, y float64) float64 {
		return a*x + b*y + c
	}
}

// Creates a multiplicative curve combinator z = a*x*y + b.
func Multiplicative(a, b float64) Combinator {
	return func(x, y float64) float64 {
		return a*x*y + b
	}
}
