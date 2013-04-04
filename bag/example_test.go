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
package bag_test

import (
	"fmt"
	"github.com/karalabe/cookiejar/bag"
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
