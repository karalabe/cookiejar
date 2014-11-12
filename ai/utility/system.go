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

import "fmt"

// Utility theory based AI system configuration.
type Config struct {
	Input map[string]InputConf
	Combo map[string]ComboConf
}

// Configuration for input based utility curve(s).
type InputConf struct {
	Count int     // Number of curves in the set (0 defaults to singleton)
	Min   float64 // Interval start for normalization
	Max   float64 // Interval end for normalization
	Curve Curve   // Function mapping the data to a curve
}

// Configuration for combination based utility curve(s).
type ComboConf struct {
	Count int        // Number of curves in the set (0 defaults to singleton)
	SrcA  string     // First input source of the combinator
	SrcB  string     // Second input source of the combinator
	Comb  Combinator // Function combining the input sources
}

// Utility theory based decision making system.
type System struct {
	utils map[string]utility
}

// Creates a utility theory AI system.
func New(config *Config) *System {
	sys := &System{
		utils: make(map[string]utility),
	}
	for name, input := range config.Input {
		sys.addInput(name, &input)
	}
	for name, combo := range config.Combo {
		sys.addCombo(name, &combo)
	}
	return sys
}

// Injects a new input based utility curve set into the system.
func (s *System) addInput(name string, config *InputConf) {
	// Create the name set if multiple is needed
	var names []string
	if config.Count == 0 {
		names = []string{name}
	} else {
		names = make([]string, config.Count)
		for i := 0; i < config.Count; i++ {
			names[i] = fmt.Sprintf("%s:%d", name, i)
		}
	}
	// Create the input curve set
	for _, name := range names {
		util := newInputUtility(config.Curve)
		util.limit(config.Min, config.Max)
		s.utils[name] = util
	}
}

// Injects a new combinatorial utility curve set into the system.
func (s *System) addCombo(name string, config *ComboConf) {
	// Singleton combinations require separate handling
	if config.Count == 0 {
		srcA := s.utils[config.SrcA]
		srcB := s.utils[config.SrcB]

		s.utils[name] = newComboUtility(config.Comb, srcA, srcB)
		return
	}
	// Construct the utility set
	for i := 0; i < config.Count; i++ {
		name := fmt.Sprintf("%s:%d", name, i)

		var srcA utility
		if _, ok := s.utils[config.SrcA]; ok {
			srcA = s.utils[config.SrcA]
		} else {
			srcA = s.utils[fmt.Sprintf("%s:%d", config.SrcA, i)]
		}

		var srcB utility
		if _, ok := s.utils[config.SrcB]; ok {
			srcB = s.utils[config.SrcB]
		} else {
			srcB = s.utils[fmt.Sprintf("%s:%d", config.SrcB, i)]
		}
		s.utils[name] = newComboUtility(config.Comb, srcA, srcB)
	}
}

// Sets the normalization limits for data a utility.
func (s *System) Limit(name string, min, max float64) {
	s.utils[name].(*inputUtility).limit(min, max)
}

// Sets the normalization limits of a member of a data utility set.
func (s *System) LimitOne(name string, index int, min, max float64) {
	s.utils[fmt.Sprintf("%s:%d", name, index)].(*inputUtility).limit(min, max)
}

// Sets the normalization limits of a whole data utility set globally.
func (s *System) LimitGlobal(name string, min, max float64) {
	for i := 0; ; i++ {
		if util, ok := s.utils[fmt.Sprintf("%s:%d", name, i)]; ok {
			util.(*inputUtility).limit(min, max)
		} else {
			return
		}
	}
}

// Updates the input of a data utility.
func (s *System) Update(name string, input float64) {
	s.utils[name].(*inputUtility).update(input)
}

// Updates the input of a member of a data utility set.
func (s *System) UpdateOne(name string, index int, input float64) {
	s.utils[fmt.Sprintf("%s:%d", name, index)].(*inputUtility).update(input)
}

// Updates the input of a member of whole data utility set individually.
func (s *System) UpdateAll(name string, input []float64) {
	for i, in := range input {
		s.UpdateOne(name, i, in)
	}
}

// Evaluates a singleton utility.
func (s *System) Evaluate(name string) float64 {
	return s.utils[name].Evaluate()
}

// Evaluates a member of a utility set.
func (s *System) EvaluateOne(name string, index int) float64 {
	return s.utils[fmt.Sprintf("%s:%d", name, index)].Evaluate()
}
