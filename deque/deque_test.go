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
package deque

import (
	"math/rand"
	"testing"
)

func TestDeque(t *testing.T) {
	// Create some initial data
	size := 1048576
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	deque := New()
	for rep := 0; rep < 2; rep++ {
		// Push all the data into the deque, each one on a different side
		outs := []int{}
		for i := 0; i < size; i++ {
			if i%2 == 0 {
				deque.PushLeft(data[i])
			} else {
				deque.PushRight(data[i])
			}
			// Pop out every third and fourth (inversely than inserted)
			if i%4 == 2 {
				outs = append(outs, deque.PopRight().(int))
			} else if i%4 == 3 {
				outs = append(outs, deque.PopLeft().(int))
			}
		}
		rest := []int{}
		for !deque.Empty() {
			if len(rest)%2 == 0 {
				rest = append(rest, deque.PopRight().(int))
			} else {
				rest = append(rest, deque.PopLeft().(int))
			}
		}
		// Make sure the contents of the resulting slices are ok
		for i := 1; i < size; i += 4 {
			if data[i] != outs[i/2] {
				t.Errorf("push/pop mismatch: have %v, want %v.", outs[i/2], data[i])
			}
			if data[i+1] != outs[i/2+1] {
				t.Errorf("push/pop mismatch: have %v, want %v.", outs[i/2+1], data[i+1])
			}
		}
		for i := 0; i < size; i += 4 {
			if data[i] != rest[len(rest)-1-i/2] {
				t.Errorf("push/pop mismatch: have %v, want %v.", rest[len(rest)-1-i/2], data[i])
			}
			if data[i+3] != rest[len(rest)-1-i/2-1] {
				t.Errorf("push/pop mismatch: have %v, want %v.", rest[len(rest)-1-i/2-1], data[i+1])
			}
		}
	}
}

func TestQueue(t *testing.T) {
	// Create some initial data
	size := 1048576
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	// Simulate a queue in both directions
	deque := New()
	for rep := 0; rep < 2; rep++ {
		for _, val := range data {
			deque.PushLeft(val)
		}
		outs := []int{}
		for !deque.Empty() {
			outs = append(outs, deque.PopRight().(int))
		}
		for i := 0; i < len(data); i++ {
			if data[i] != outs[i] {
				t.Errorf("push/pop mismatch: have %v, want %v.", outs[i], data[i])
			}
		}
		for _, val := range data {
			deque.PushRight(val)
		}
		outs = []int{}
		for !deque.Empty() {
			outs = append(outs, deque.PopLeft().(int))
		}
		for i := 0; i < len(data); i++ {
			if data[i] != outs[i] {
				t.Errorf("push/pop mismatch: have %v, want %v.", outs[i], data[i])
			}
		}
	}
}

func TestStack(t *testing.T) {
	// Create some initial data
	size := 1048576
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	// Simulate a stack in both directions
	deque := New()
	for rep := 0; rep < 2; rep++ {
		for _, val := range data {
			deque.PushLeft(val)
		}
		outs := []int{}
		for !deque.Empty() {
			outs = append(outs, deque.PopLeft().(int))
		}
		for i := 0; i < len(data); i++ {
			if data[i] != outs[len(outs)-i-1] {
				t.Errorf("push/pop mismatch: have %v, want %v.", outs[len(outs)-i-1], data[i])
			}
		}
		for _, val := range data {
			deque.PushRight(val)
		}
		outs = []int{}
		for !deque.Empty() {
			outs = append(outs, deque.PopRight().(int))
		}
		for i := 0; i < len(data); i++ {
			if data[i] != outs[len(outs)-i-1] {
				t.Errorf("push/pop mismatch: have %v, want %v.", outs[len(outs)-i-1], data[i])
			}
		}
	}
}

func TestReset(t *testing.T) {
	// Push some stuff into the deque
	size := 1048576
	deque := New()
	for i := 0; i < size; i++ {
		deque.PushLeft(i)
	}
	// Clear and verify
	deque.Reset()
	if !deque.Empty() {
		t.Errorf("deque not empty after reset: %v", deque)
	}
}

func BenchmarkPush(b *testing.B) {
	deque := New()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			deque.PushLeft(i)
		} else {
			deque.PushRight(i)
		}
	}
}

func BenchmarkPop(b *testing.B) {
	deque := New()
	for i := 0; i < b.N; i++ {
		deque.PushLeft(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			deque.PopLeft()
		} else {
			deque.PopRight()
		}
	}
}
