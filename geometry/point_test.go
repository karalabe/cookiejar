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

package geom

import (
	"math"
	"testing"
)

var origin2 = &Point2{0, 0}
var unitX2 = &Point2{1, 0}
var unitY2 = &Point2{0, 1}

var origin3 = &Point3{0, 0, 0}
var unitX3 = &Point3{1, 0, 0}
var unitY3 = &Point3{0, 1, 0}
var unitZ3 = &Point3{0, 0, 1}
var diag3 = &Point3{1, 1, 1}

func TestDist2D(t *testing.T) {
	if d := origin2.Dist(unitX2); math.Abs(d-1) > eps {
		t.Errorf("distance mismatch: have %v, want %v.", d, 1)
	}
	if d := origin2.Dist(unitY2); math.Abs(d-1) > eps {
		t.Errorf("distance mismatch: have %v, want %v.", d, 1)
	}
	if d := unitX2.Dist(unitY2); math.Abs(d-math.Sqrt2) > eps {
		t.Errorf("distance mismatch: have %v, want %v.", d, math.Sqrt2)
	}
}

func TestDistSqr2D(t *testing.T) {
	if d := origin2.DistSqr(unitX2); math.Abs(d-1) > eps {
		t.Errorf("squared distance mismatch: have %v, want %v.", d, 1)
	}
	if d := origin2.DistSqr(unitY2); math.Abs(d-1) > eps {
		t.Errorf("squared distance mismatch: have %v, want %v.", d, 1)
	}
	if d := unitX2.DistSqr(unitY2); math.Abs(d-2) > eps {
		t.Errorf("squared distance mismatch: have %v, want %v.", d, 2)
	}
}

func TestEqual2D(t *testing.T) {
	// Check X coordinate
	a, b, c := &Point2{0, 0}, &Point2{0.999999 * eps, 0}, &Point2{eps, 0}
	if !a.Equal(b) {
		t.Errorf("equality should hold: %v == %v (given eps)", a, b)
	}
	if a.Equal(c) {
		t.Errorf("equality should not hold: %v == %v (given eps)", a, c)
	}
	// Check Y coordinate
	a, b, c = &Point2{0, 0}, &Point2{0, 0.999999 * eps}, &Point2{0, eps}
	if !a.Equal(b) {
		t.Errorf("equality should hold: %v == %v (given eps)", a, b)
	}
	if a.Equal(c) {
		t.Errorf("equality should not hold: %v == %v (given eps)", a, c)
	}
}

func TestDist3D(t *testing.T) {
	if d := origin3.Dist(unitX3); math.Abs(d-1) > eps {
		t.Errorf("distance mismatch: have %v, want %v.", d, 1)
	}
	if d := origin3.Dist(unitY3); math.Abs(d-1) > eps {
		t.Errorf("distance mismatch: have %v, want %v.", d, 1)
	}
	if d := origin3.Dist(unitZ3); math.Abs(d-1) > eps {
		t.Errorf("distance mismatch: have %v, want %v.", d, 1)
	}
	if d := origin3.Dist(diag3); math.Abs(d-math.Sqrt(3)) > eps {
		t.Errorf("distance mismatch: have %v, want %v.", d, math.Sqrt(3))
	}
}

func TestDistSqr3D(t *testing.T) {
	if d := origin3.DistSqr(unitX3); math.Abs(d-1) > eps {
		t.Errorf("squared distance mismatch: have %v, want %v.", d, 1)
	}
	if d := origin3.DistSqr(unitY3); math.Abs(d-1) > eps {
		t.Errorf("squared distance mismatch: have %v, want %v.", d, 1)
	}
	if d := origin3.DistSqr(unitZ3); math.Abs(d-1) > eps {
		t.Errorf("squared distance mismatch: have %v, want %v.", d, 1)
	}
	if d := origin3.DistSqr(diag3); math.Abs(d-3) > eps {
		t.Errorf("squared distance mismatch: have %v, want %v.", d, 3)
	}
}

func TestEqual3D(t *testing.T) {
	// Check X coordinate
	a, b, c := &Point3{0, 0, 0}, &Point3{0.999999 * eps, 0, 0}, &Point3{eps, 0, 0}
	if !a.Equal(b) {
		t.Errorf("equality should hold: %v == %v (given eps)", a, b)
	}
	if a.Equal(c) {
		t.Errorf("equality should not hold: %v == %v (given eps)", a, c)
	}
	// Check Y coordinate
	a, b, c = &Point3{0, 0, 0}, &Point3{0, 0.999999 * eps, 0}, &Point3{0, eps, 0}
	if !a.Equal(b) {
		t.Errorf("equality should hold: %v == %v (given eps)", a, b)
	}
	if a.Equal(c) {
		t.Errorf("equality should not hold: %v == %v (given eps)", a, c)
	}
	// Check Z coordinate
	a, b, c = &Point3{0, 0, 0}, &Point3{0, 0, 0.999999 * eps}, &Point3{0, 0, eps}
	if !a.Equal(b) {
		t.Errorf("equality should hold: %v == %v (given eps)", a, b)
	}
	if a.Equal(c) {
		t.Errorf("equality should not hold: %v == %v (given eps)", a, c)
	}
}
