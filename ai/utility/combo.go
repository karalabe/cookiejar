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

import (
	"math"

	"gopkg.in/karalabe/cookiejar.v2/collections/bag"
)

// Combination utility derived from two other utilities.
type comboUtility struct {
	combinator Combinator    // Curve transformation combinator
	srcA, srcB singleUtility // Base utilities from which to derive this one

	children *bag.Bag // Derived utilities based on the current one

	reset  bool    // Flag whether the output is not yet calculated
	output float64 // Cached output utility value
}

// Creates a new derived utility based on two existing ones.
func newComboUtility(combinator Combinator) *comboUtility {
	return &comboUtility{
		combinator: combinator,
		children:   bag.New(),
		reset:      false,
	}
}

// Finishes initialization with the two source components.
func (u *comboUtility) Init(srcA, srcB utility) {
	u.srcA = srcA.(singleUtility)
	u.srcB = srcB.(singleUtility)

	srcA.Dependency(u)
	srcB.Dependency(u)
}

// Resets the utility value, forcing it to recalculate when needed again.
func (u *comboUtility) Reset() {
	if !u.reset {
		u.reset = true
		u.children.Do(func(util interface{}) {
			util.(*comboUtility).Reset()
		})
	}
}

// Adds a new dependency to the utility hierarchy.
func (u *comboUtility) Dependency(util utility) {
	u.children.Insert(util)
}

// Evaluates the utility chain and returns the current value.
func (u *comboUtility) Evaluate() float64 {
	// If the utility was reset, reevaluate it
	if u.reset {
		u.output = math.Min(1, math.Max(0, u.combinator(u.srcA.Evaluate(), u.srcB.Evaluate())))
		u.reset = false
	}
	// Return the currently set value
	return u.output
}
