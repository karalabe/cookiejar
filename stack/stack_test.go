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
package stack

import (
	"math/rand"
	"testing"
)

func TestStack(t *testing.T) {
	// Create some initial data
	size := 1048576
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	stack := New()
	for rep := 0; rep < 2; rep++ {
		// Push all the data into the stack, pop out every second
		secs := []int{}
		for i := 0; i < size; i++ {
			stack.Push(data[i])
			if i%2 == 0 {
				secs = append(secs, stack.Pop().(int))
			}
		}
		rest := []int{}
		for !stack.Empty() {
			rest = append(rest, stack.Pop().(int))
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

func TestReset(t *testing.T) {
	// Push some stuff onto the stack
	size := 1048576
	stack := New()
	for i := 0; i < size; i++ {
		stack.Push(i)
	}
	// Clear and verify
	stack.Reset()
	if !stack.Empty() {
		t.Errorf("stack not empty after reset: %v", stack)
	}
}

func BenchmarkPush(b *testing.B) {
	stack := New()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	stack := New()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for !stack.Empty() {
		stack.Pop()
	}
}
