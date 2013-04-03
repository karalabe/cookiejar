// CookieJar - A contestant's algorithm toolbox
// Copyright 2013 Peter Szilagyi. All rights reserved.
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
//
// Author: peterke@gmail.com (Peter Szilagyi)

// Package set implements simple present/not data structure supporting arbitrary
// types (even a mixture).
//
// Internally it uses a simple map assigning zero-byte struct{}s to keys.
package set

// Set data structure.
type Set struct {
	data map[interface{}]struct{}
}

// Creates a new empty set.
func New() *Set {
	return &Set{make(map[interface{}]struct{})}
}

// Inserts an element into the set.
func (s *Set) Insert(val interface{}) {
	s.data[val] = struct{}{}
}

// Removes an element from the set. If none was present, nothing is done.
func (s *Set) Remove(val interface{}) {
	delete(s.data, val)
}

// Checks whether an element is inside the set.
func (s *Set) Exists(val interface{}) bool {
	_, ok := s.data[val]
	return ok
}

// Executes a function for every element in the set.
func (s *Set) Do(f func(interface{})) {
	for val, _ := range s.data {
		f(val)
	}
}

// Clears the contents of a set.
func (s *Set) Reset() {
	s.data = make(map[interface{}]struct{})
}
