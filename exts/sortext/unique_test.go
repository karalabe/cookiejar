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
package sortext

import (
	"sort"
	"testing"
)

type uniqueTest struct {
	data []int
	num  int
}

var uniqueTests = []uniqueTest{
	{[]int{}, 0},
	{[]int{1}, 1},
	{
		[]int{
			1,
			2, 2,
			3, 3, 3,
			4, 4, 4, 4,
			5, 5, 5, 5, 5,
			6, 6, 6, 6, 6, 6,
		},
		6,
	},
}

func TestUnique(t *testing.T) {
	for i, tt := range uniqueTests {
		n := Unique(sort.IntSlice(tt.data))
		if n != tt.num {
			t.Errorf("test %d: unique count mismatch: have %v, want %v.", i, n, tt.num)
		}
		for j := 0; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if tt.data[j] >= tt.data[k] {
					t.Errorf("test %d: uniqueness violation: (%d, %d) in %v.", i, j, k, tt.data[:n])
				}
			}
		}
	}
}
