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

package geometry

import (
	"math"
)

// Two dimensional point.
type Point2 struct {
	X, Y float64
}

// Three dimensional point.
type Point3 struct {
	X, Y, Z float64
}

// Allocates and returns a new 2D point.
func NewPoint2(x, y float64) *Point2 {
	return &Point2{x, y}
}

// Allocates and returns a new 3D point.
func NewPoint3(x, y, z float64) *Point3 {
	return &Point3{x, y, z}
}

// Calculates the distance between x and y.
func (x *Point2) Dist(y *Point2) float64 {
	return math.Sqrt(x.DistSqr(y))
}

// Calculates the distance between x and y.
func (x *Point3) Dist(y *Point3) float64 {
	return math.Sqrt(x.DistSqr(y))
}

// Calculates the squared distance between x and y.
func (x *Point2) DistSqr(y *Point2) float64 {
	dx := x.X - y.X
	dy := x.Y - y.Y
	return dx*dx + dy*dy
}

// Calculates the squared distance between x and y.
func (x *Point3) DistSqr(y *Point3) float64 {
	dx := x.X - y.X
	dy := x.Y - y.Y
	dz := x.Z - y.Z
	return dx*dx + dy*dy + dz*dz
}

// Returns whether two points are equal.
func (x *Point2) Equal(y *Point2) bool {
	return math.Abs(x.X-y.X) < eps && math.Abs(x.Y-y.Y) < eps
}

// Returns whether two points are equal.
func (x *Point3) Equal(y *Point3) bool {
	return math.Abs(x.X-y.X) < eps && math.Abs(x.Y-y.Y) < eps && math.Abs(x.Z-y.Z) < eps
}
