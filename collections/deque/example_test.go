// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
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

package deque_test

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

// Simple usage example that inserts the numbers 0, 1, 2 into a deque and then
// removes them one by one, varying the removal side.
func Example_usage() {
	// Create a deque an push some data in
	d := deque.New()
	for i := 0; i < 3; i++ {
		d.PushLeft(i)
	}
	// Pop out the deque contents and display them
	fmt.Println(d.PopLeft())
	fmt.Println(d.PopRight())
	fmt.Println(d.PopLeft())
	// Output:
	// 2
	// 0
	// 1
}
