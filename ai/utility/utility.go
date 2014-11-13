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

// Generic curve function type. Input is automatically normalized to [0, 1]
// prior to passing it to the curve. Outputs is clamped to [0, 1].
type Curve func(float64) float64

// Generic curve function combinator. Input is guaranteed to be in [0, 1].
// Outputs is clamped to [0, 1].
type Combinator func(x, y float64) float64

type utility interface {
	Dependency(util utility)
}

type singleUtility interface {
	utility
	Evaluate() float64
}

type multiUtility interface {
	utility
	Member(id int) singleUtility
	Evaluate(id int) float64
}
