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
package graph_test

import (
	"fmt"
	"gopkg.in/karalabe/cookiejar.v1/graph"
)

// Creates a simple graph, and prints the degree of each vertex.
func Example_usage() {
	// Create a star shaped 5 vertex graph
	g := graph.New(5)
	g.Connect(0, 2)
	g.Connect(0, 3)
	g.Connect(1, 3)
	g.Connect(1, 4)
	g.Connect(2, 4)

	// For each vertex, count the outgoing edges
	for v := 0; v < g.Vertices(); v++ {
		degree := 0
		g.Do(v, func(peer interface{}) {
			degree++
		})
		fmt.Printf("%v: %v\n", v, degree)
	}
	// Output:
	// 0: 2
	// 1: 2
	// 2: 2
	// 3: 2
	// 4: 2
}
