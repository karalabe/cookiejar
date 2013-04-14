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
package bfs_test

import (
	"fmt"
	"github.com/karalabe/cookiejar/graph"
	"github.com/karalabe/cookiejar/graph/bfs"
)

// Small API demo based on a trie graph and a few disconnected vertices.
func Example_usage() {
	// Create the graph
	g := graph.New(7)
	g.Connect(0, 1)
	g.Connect(1, 2)
	g.Connect(1, 4)
	g.Connect(2, 3)
	g.Connect(4, 5)

	// Create the breadth first search algo structure for g and source node #2
	b := bfs.New(g, 0)

	// Get the path between #0 (source) and #2
	fmt.Println("Path 0->5:", b.Path(5))
	fmt.Println("Order:", b.Order())
	fmt.Println("Reachable #4 #6:", b.Reachable(4), b.Reachable(6))

	// Output:
	// Path 0->5: [0 1 4 5]
	// Order: [0 1 2 4 3 5]
	// Reachable #4 #6: true false
}
