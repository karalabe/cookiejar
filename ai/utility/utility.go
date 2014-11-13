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

// Generic curve function type. Input is automatically normalized to [0, 1]
// prior to passing it to the curve. Outputs is clamped to [0, 1].
type Curve func(float64) float64

// Generic curve function combinator. Input is guaranteed to be in [0, 1].
// Outputs is clamped to [0, 1].
type Combinator func(x, y float64) float64

type utility interface {
	Dependency(util utility)
	Evaluate() float64
}

// Data-source based utility, normalizing and transforming an input stream by an
// assigned curve.
type inputUtility struct {
	curve   Curve   // Data transformation curve
	lo, hi  float64 // Normalization limits
	nonZero bool    // Absolute zero allowed or not

	children *bag.Bag // Derived utilities based on the current one

	value float64 // Generated output utility value
}

// Creates a new data source utility and associated a transformation curve.
func newInputUtility(curve Curve, nonZero bool) *inputUtility {
	return &inputUtility{
		curve:    curve,
		nonZero:  nonZero,
		children: bag.New(),
	}
}

// Sets the data limits used during normalization.
func (u *inputUtility) limit(lo, hi float64) {
	u.lo, u.hi = lo, hi
}

// Updates the utility to a new data value.
func (u *inputUtility) update(input float64) {
	// Normalize the input and calculate the output
	if diff := u.hi - u.lo; diff != 0 {
		input = (input - u.lo) / diff
	}
	u.value = math.Min(1, math.Max(0, u.curve(input)))

	// Prevent absolute zero as it can completely ruin the outputs
	if u.nonZero && u.value == 0 {
		u.value = float64(1e-9)
	}
	// Reset all derived utilities
	u.children.Do(func(util interface{}) {
		util.(*comboUtility).Reset()
	})
}

// Adds a new dependency to the utility hierarchy.
func (u *inputUtility) Dependency(util utility) {
	u.children.Insert(util)
}

// Returns the utility value for the set data point.
func (u *inputUtility) Evaluate() float64 {
	return u.value
}

// Combination utility derived from two other utilities.
type comboUtility struct {
	combinator Combinator // Curve transformation combinator
	srcA, srcB utility    // Base utilities from which to derive this one

	children *bag.Bag // Derived utilities based on the current one

	reset bool    // Flag whether the utility is reset
	value float64 // Generated output utility value
}

// Creates a new derived utility based on two existing ones.
func newComboUtility(combinator Combinator, srcA, srcB utility) *comboUtility {
	// Create the derived utility
	util := &comboUtility{
		combinator: combinator,
		srcA:       srcA,
		srcB:       srcB,
		children:   bag.New(),
		reset:      false,
	}
	// Register the dependencies and return
	srcA.Dependency(util)
	srcB.Dependency(util)

	return util
}

// Resets the utility value, forcing it to recalculate when needed again.
func (u *comboUtility) Reset() {
	// Skip if already reset
	if u.reset {
		return
	}
	// Otherwise reset all derived utilities
	u.reset = true
	u.children.Do(func(util interface{}) {
		util.(*comboUtility).Reset()
	})
}

// Adds a new dependency to the utility hierarchy.
func (u *comboUtility) Dependency(util utility) {
	u.children.Insert(util)
}

// Evaluates the utility chain and returns the current value.
func (u *comboUtility) Evaluate() float64 {
	// If the utility was reset, reevaluate it
	if u.reset {
		u.value = math.Min(1, math.Max(0, u.combinator(u.srcA.Evaluate(), u.srcB.Evaluate())))
		u.reset = false
	}
	// Return the currently set value
	return u.value
}
