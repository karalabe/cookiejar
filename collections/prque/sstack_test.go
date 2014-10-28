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
package prque

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSstack(t *testing.T) {
	// Create some initial data
	size := 16 * blockSize
	data := make([]*item, size)
	for i := 0; i < size; i++ {
		data[i] = &item{rand.Int(), rand.Float32()}
	}
	stack := newSstack()
	for rep := 0; rep < 2; rep++ {
		// Push all the data into the stack, pop out every second
		secs := []*item{}
		for i := 0; i < size; i++ {
			stack.Push(data[i])
			if i%2 == 0 {
				secs = append(secs, stack.Pop().(*item))
			}
		}
		rest := []*item{}
		for stack.Len() > 0 {
			rest = append(rest, stack.Pop().(*item))
		}
		// Make sure the contents of the resulting slices are ok
		for i := 0; i < size; i++ {
			if i%2 == 0 && data[i] != secs[i/2] {
				t.Errorf("push/pop mismatch: have %v, want %v.", secs[i/2], data[i])
			}
			if i%2 == 1 && data[i] != rest[len(rest)-i/2-1] {
				t.Errorf("push/pop mismatch: have %v, want %v.", rest[len(rest)-i/2-1], data[i])
			}
		}
	}
}

func TestSstackSort(t *testing.T) {
	// Create some initial data
	size := 16 * blockSize
	data := make([]*item, size)
	for i := 0; i < size; i++ {
		data[i] = &item{rand.Int(), float32(i)}
	}
	// Push all the data into the stack
	stack := newSstack()
	for _, val := range data {
		stack.Push(val)
	}
	// Sort and pop the stack contents (should reverse the order)
	sort.Sort(stack)
	for _, val := range data {
		out := stack.Pop()
		if out != val {
			t.Errorf("push/pop mismatch after sort: have %v, want %v.", out, val)
		}
	}
}

func TestSstackReset(t *testing.T) {
	// Push some stuff onto the stack
	size := 16 * blockSize
	stack := newSstack()
	for i := 0; i < size; i++ {
		stack.Push(&item{i, float32(i)})
	}
	// Clear and verify
	stack.Reset()
	if stack.Len() != 0 {
		t.Errorf("stack not empty after reset: %v", stack)
	}
}
