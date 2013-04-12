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

// Package prque implements a priority queue data structure supporting arbitrary
// value types and float priorities.
//
// The reasoning behind using floats for the priorities vs. ints or interfaces
// was larger flexibility without sacrificing too much performance or code
// complexity.
//
// If you would like to use a min-priority queue, simply negate the priorities.
//
// Internally the queue is based on the standard heap package working on a
// sortable version of the block based stack.
package prque

import (
	"container/heap"
)

// Priority queue data structure.
type Prque struct {
	cont *sstack
}

// Creates a new priority queue.
func New() *Prque {
	return &Prque{newSstack()}
}

// Pushes a value with a given priority into the queue, expanding if necessary.
func (p *Prque) Push(data interface{}, priority float32) {
	heap.Push(p.cont, &item{data, priority})
}

// Pops the value with the greates priority off the stack and returns it.
// Currently no shrinking is done.
func (p *Prque) Pop() (interface{}, float32) {
	item := heap.Pop(p.cont).(*item)
	return item.value, item.priority
}

// Checks whether the priority queue is empty.
func (p *Prque) Empty() bool {
	return p.cont.Len() == 0
}

// Returns the number of element in the priority queue.
func (p *Prque) Size() int {
	return p.cont.Len()
}

// Clears the contents of the priority queue.
func (p *Prque) Reset() {
	p.cont.Reset()
}
