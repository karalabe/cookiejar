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

// Clears the contents of a bag.
func (b *Bag) Reset() {
	b.data = make(map[interface{}]int)
}
