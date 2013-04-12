// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice,
//       this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright
//       notice, this list of conditions and the following disclaimer in the
//       documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.
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

// Returns the number of elements in the set.
func (s *Set) Size() int {
	return len(s.data)
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
