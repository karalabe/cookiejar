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
package bag_test

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v1/collections/bag"
)

// Small demo of the common functions in the bag package.
func Example_usage() {
	// Create a new bag with some integers in it
	b := bag.New()
	for i := 0; i < 10; i++ {
		b.Insert(i)
	}
	b.Insert(8)
	// Remove every odd integer
	for i := 1; i < 10; i += 2 {
		b.Remove(i)
	}
	// Print the element count of all numbers
	for i := 0; i < 10; i++ {
		fmt.Printf("#%d: %d\n", i, b.Count(i))
	}
	// Calculate the sum with a Do iteration
	sum := 0
	b.Do(func(val interface{}) {
		sum += val.(int)
	})
	fmt.Println("Sum:", sum)
	// Output:
	// #0: 1
	// #1: 0
	// #2: 1
	// #3: 0
	// #4: 1
	// #5: 0
	// #6: 1
	// #7: 0
	// #8: 2
	// #9: 0
	// Sum: 28
}
