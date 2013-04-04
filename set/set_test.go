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
package set

import (
	"math/rand"
	"testing"
)

func TestSet(t *testing.T) {
	// Create some initial data
	size := 1048576
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	// Fill the set with the data and verify that they're all set
	set := New()
	for val := range data {
		set.Insert(val)
	}
	for val := range data {
		if !set.Exists(val) {
			t.Errorf("failed to locate element in set: %v in %v", val, set)
		}
	}
	// Remove a few elements and ensure they're out
	rems := data[:1024]
	for val := range rems {
		set.Remove(val)
	}
	for val := range rems {
		if set.Exists(val) {
			t.Errorf("element exists after remove: %v in %v", val, set)
		}
	}
	// Calcualte the sum of the remainder and verify
	sumSet := int64(0)
	set.Do(func(val interface{}) {
		sumSet += int64(val.(int))
	})
	sumDat := int64(0)
	for val := range data {
		sumDat += int64(val)
	}
	for val := range rems {
		sumDat -= int64(val)
	}
	if sumSet != sumDat {
		t.Errorf("iteration sum mismatch: have %v, want %v", sumSet, sumDat)
	}
	// Clear the set and ensure nothing's left
	set.Reset()
	for val := range data {
		if set.Exists(val) {
			t.Errorf("element exists after reset: %v in %v", val, set)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	// Execute the benchmark
	b.ResetTimer()
	set := New()
	for val := range data {
		set.Insert(val)
	}
}

func BenchmarkRemove(b *testing.B) {
	// Create some initial data and fill the set
	data := rand.Perm(b.N)
	set := New()
	for val := range data {
		set.Insert(val)
	}
	// Execute the benchmark (different order)
	rems := rand.Perm(b.N)
	b.ResetTimer()
	for val := range rems {
		set.Remove(val)
	}
}

func BenchmarkDo(b *testing.B) {
	// Create some initial data
	data := make([]int, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Int()
	}
	// Fill the set with it
	set := New()
	for val := range data {
		set.Insert(val)
	}
	// Execute the benchmark
	b.ResetTimer()
	set.Do(func(val interface{}) {
		// Do nothing
	})
}
