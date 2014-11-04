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
package mathext

import (
	"math/big"
	"testing"
)

type wrapperTest struct {
	name string
	res  int
	good int
}

func TestInt(t *testing.T) {
	if m := MaxInt(-10, 10); m != 10 {
		t.Errorf("max mismatch: have %v, want %v.", m, 10)
	}
	if m := MaxInt(10, -10); m != 10 {
		t.Errorf("max mismatch: have %v, want %v.", m, 10)
	}
	if m := MinInt(-10, 10); m != -10 {
		t.Errorf("min mismatch: have %v, want %v.", m, -10)
	}
	if m := MinInt(-10, 10); m != -10 {
		t.Errorf("min mismatch: have %v, want %v.", m, -10)
	}
}

func TestBigInt(t *testing.T) {
	pos := big.NewInt(10)
	neg := big.NewInt(10)

	if m := MaxBigInt(neg, pos); m.Cmp(pos) != 0 {
		t.Errorf("max mismatch: have %v, want %v.", m, pos)
	}
	if m := MaxBigInt(pos, neg); m.Cmp(pos) != 0 {
		t.Errorf("max mismatch: have %v, want %v.", m, pos)
	}
	if m := MinBigInt(neg, pos); m.Cmp(neg) != 0 {
		t.Errorf("min mismatch: have %v, want %v.", m, neg)
	}
	if m := MinBigInt(neg, pos); m.Cmp(neg) != 0 {
		t.Errorf("min mismatch: have %v, want %v.", m, neg)
	}
}

func TestBigRat(t *testing.T) {
	pos := big.NewRat(10, 314)
	neg := big.NewRat(10, 314)

	if m := MaxBigRat(neg, pos); m.Cmp(pos) != 0 {
		t.Errorf("max mismatch: have %v, want %v.", m, pos)
	}
	if m := MaxBigRat(pos, neg); m.Cmp(pos) != 0 {
		t.Errorf("max mismatch: have %v, want %v.", m, pos)
	}
	if m := MinBigRat(neg, pos); m.Cmp(neg) != 0 {
		t.Errorf("min mismatch: have %v, want %v.", m, neg)
	}
	if m := MinBigRat(neg, pos); m.Cmp(neg) != 0 {
		t.Errorf("min mismatch: have %v, want %v.", m, neg)
	}
}
