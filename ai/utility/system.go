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

// Package utility implements a reasoner AI based on utility theory.
package utility

import (
	"fmt"
	"strings"
)

// Utility theory based decision making system.
type System struct {
	utils map[string]utility
}

// Creates a utility theory AI system.
func New() *System {
	return &System{
		utils: make(map[string]utility),
	}
}

// Adds a new utility to the reasoner system.
func (s *System) Add(name string, curve Curve) {
	// Transform to lowercase and do some sanity checks
	name = strings.ToLower(name)
	if _, ok := s.utils[name]; ok {
		panic(fmt.Sprintf("Utility already registered: %s", name))
	}
	// Create the new utility
	s.utils[name] = newSourceUtility(curve)
}

// Combines two existing utilities in the system.
func (s *System) Combine(name string, utilA, utilB string, combinator Combinator) {
	// Transform to lowercase and do some sanity checks
	name = strings.ToLower(name)
	if _, ok := s.utils[name]; ok {
		panic(fmt.Sprintf("Utility already registered: %s", name))
	}
	// Look up the dependent utilities
	srcA, ok := s.utils[utilA]
	if !ok {
		panic(fmt.Sprintf("Utility A not registered: %s", utilA))
	}
	srcB, ok := s.utils[utilB]
	if !ok {
		panic(fmt.Sprintf("Utility B not registered: %s", utilB))
	}
	// Create the new utility
	s.utils[name] = newDerivedUtility(combinator, srcA, srcB)
}

// Sets the normalization limits for a utility curve.
func (s *System) Limit(name string, lo, hi float64) {
	// Transform to lowercase and do some sanity checks
	name = strings.ToLower(name)
	if util, ok := s.utils[name]; !ok {
		panic(fmt.Sprintf("Utility not registered: %s", name))
	} else {
		util.(*sourceUtility).limit(lo, hi)
	}
}

// Updates the utility to a new data value.
func (s *System) Update(name string, input float64) {
	// Transform to lowercase and do some sanity checks
	name = strings.ToLower(name)
	if util, ok := s.utils[name]; !ok {
		panic(fmt.Sprintf("Utility not registered: %s", name))
	} else {
		util.(*sourceUtility).update(input)
	}
}

// Evaluates a batch of utilities.
func (s *System) Evaluate(names []string) []float64 {
	values := make([]float64, len(names))
	for i := 0; i < len(names); i++ {
		values[i] = s.utils[names[i]].Evaluate()
	}
	return values
}
