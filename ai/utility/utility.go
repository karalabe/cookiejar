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

type utility interface {
	Dependency(util utility)
	Evaluate() float64
}

// Data-source based utility, normalizing and transforming an input stream by an
// assigned curve.
type sourceUtility struct {
	curve  Curve   // Data transformation curve
	lo, hi float64 // Normalization limits

	children *bag.Bag // Derived utilities based on the current one

	value float64 // Generated output utility value
}

// Creates a new data source utility and associated a transformation curve.
func newSourceUtility(curve Curve) utility {
	return &sourceUtility{
		curve:    curve,
		children: bag.New(),
	}
}

// Sets the data limits used during normalization.
func (u *sourceUtility) limit(lo, hi float64) {
	u.lo, u.hi = lo, hi
}

// Updates the utility to a new data value.
func (u *sourceUtility) update(input float64) {
	// Normalize the input and calculate the output
	input = (input - u.lo) / (u.hi - u.lo)
	u.value = math.Min(1, math.Max(0, u.curve(input)))

	// Reset all derived utilities
	u.children.Do(func(util interface{}) {
		util.(*derivedUtility).Reset()
	})
}

// Adds a new dependency to the utility hierarchy.
func (u *sourceUtility) Dependency(util utility) {
	u.children.Insert(util)
}

// Returns the utility value for the set data point.
func (u *sourceUtility) Evaluate() float64 {
	return u.value
}

// Combination utility derived from two other utilities.
type derivedUtility struct {
	combinator Combinator // Curve transformation combinator
	srcA, srcB utility    // Base utilities from which to derive this one

	children *bag.Bag // Derived utilities based on the current one

	reset bool    // Flag whether the utility is reset
	value float64 // Generated output utility value
}

// Creates a new derived utility based on two existing ones.
func newDerivedUtility(combinator Combinator, srcA, srcB utility) utility {
	// Create the derived utility
	util := &derivedUtility{
		combinator: combinator,
		srcA:       srcA,
		srcB:       srcB,
		children:   bag.New(),
	}
	// Register the dependencies and return
	srcA.Dependency(util)
	srcB.Dependency(util)

	return util
}

// Resets the utility value, forcing it to recalculate when needed again.
func (u *derivedUtility) Reset() {
	// Skip if already reset
	if u.reset {
		return
	}
	// Otherwise reset all derived utilities
	u.reset = true
	u.children.Do(func(util interface{}) {
		util.(*derivedUtility).Reset()
	})
}

// Adds a new dependency to the utility hierarchy.
func (u *derivedUtility) Dependency(util utility) {
	u.children.Insert(util)
}

// Evaluates the utility chain and returns the current value.
func (u *derivedUtility) Evaluate() float64 {
	// If the utility was reset, reevaluate it
	if u.reset {
		u.value = math.Min(1, math.Max(0, u.combinator(u.srcA.Evaluate(), u.srcB.Evaluate())))
		u.reset = false
	}
	// Return the currently set value
	return u.value
}
