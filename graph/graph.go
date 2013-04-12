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

// Package graph implements a simple graph data structure and supporting API to
// allow implementing graph alogirthms on top.
package graph

import (
	"github.com/karalabe/cookiejar/bag"
)

// Data structure for representing a graph.
type Graph struct {
	nodes int
	edges []*bag.Bag
}

// Creates a new undirected graph.
func New(vertices int) *Graph {
	g := new(Graph)
	g.nodes = vertices
	g.edges = make([]*bag.Bag, vertices)
	for i := 0; i < vertices; i++ {
		g.edges[i] = bag.New()
	}
	return g
}

// Returns the number of vertices in the graph.
func (g *Graph) Vertices() int {
	return g.nodes
}

// Connects two vertices of a graph (may be a loopback).
func (g *Graph) Connect(a, b int) {
	g.edges[a].Insert(b)
	if a != b {
		g.edges[b].Insert(a)
	}
}

// Executes a function for every neighbor of a vertex.
func (g *Graph) Do(v int, f func(interface{})) {
	g.edges[v].Do(f)
}
