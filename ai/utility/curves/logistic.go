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

package curves

import (
	"math"

	"gopkg.in/karalabe/cookiejar.v2/ai/utility"
)

// Sigmoid curve builder, Defined as y = 1 / (slope*e) ^ 10(infl-x).
type Logistic struct {
	Infl  float64 // Point of inflection of the rate of change
	Slope float64 // Multiplier changing the slope of the curve
	Inc   bool    // Switches between increasing or decreasing curve
}

// Creates the curve mapping function.
func (l Logistic) Make() utility.Curve {
	infl, slope := l.Infl, l.Slope
	if l.Inc {
		return func(x float64) float64 {
			return 1 / (1 + math.Pow(slope*math.E, 10*(infl-x)))
		}
	} else {
		return func(x float64) float64 {
			return 1 - 1/(1+math.Pow(slope*math.E, 10*(infl-x)))
		}
	}
}
