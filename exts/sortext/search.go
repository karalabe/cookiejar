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
	"math/big"
	"sort"
)

// SearchBigInts searches for x in a sorted slice of *big.Ints and returns the
// index as specified by Search. The return value is the index to insert x if x
// is not present (it could be len(a)).
// The slice must be sorted in ascending order.
func SearchBigInts(a []*big.Int, x *big.Int) int {
	return sort.Search(len(a), func(i int) bool { return a[i].Cmp(x) >= 0 })
}

// SearchBigRats searches for x in a sorted slice of *big.Rats and returns the
// index as specified by Search. The return value is the index to insert x if x
// is not present (it could be len(a)).
// The slice must be sorted in ascending order.
func SearchBigRats(a []*big.Rat, x *big.Rat) int {
	return sort.Search(len(a), func(i int) bool { return a[i].Cmp(x) >= 0 })
}

// Search returns the result of applying SearchBigInts to the receiver and x.
func (p BigIntSlice) Search(x *big.Int) int { return SearchBigInts(p, x) }

// Search returns the result of applying SearchBigRats to the receiver and x.
func (p BigRatSlice) Search(x *big.Rat) int { return SearchBigRats(p, x) }
