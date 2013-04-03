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

// Package bag implements a multi-set data structure supporting arbitrary types
// (even a mixture).
//
// Internally it uses a simple map assigning counts to the different values
// present in the bag.
package bag

// Bag data structure (multiset).
type Bag struct {
	data map[interface{}]int
}

// Creates a new empty bag.
func New() *Bag {
	return &Bag{make(map[interface{}]int)}
}

// Inserts an element into the bag.
func (b *Bag) Insert(val interface{}) {
	b.data[val]++
}

// Removes an element from the bag. If none was present, nothing is done.
func (b *Bag) Remove(val interface{}) {
	old, ok := b.data[val]
	if ok {
		if old > 1 {
			b.data[val] = old - 1
		} else {
			delete(b.data, val)
		}
	}
}

// Counts the number of occurances of an element in the bag.
func (b *Bag) Count(val interface{}) int {
	return b.data[val]
}

// Executes a function for every element in the bag.
func (b *Bag) Do(f func(interface{})) {
	for val, cnt := range b.data {
		for ; cnt > 0; cnt-- {
			f(val)
		}
	}
}

// Clears the contents of a bag
func (b *Bag) Reset() {
	b.data = make(map[interface{}]int)
}
