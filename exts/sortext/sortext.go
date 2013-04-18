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

// Package sortext contains extensions to the base Go sort package.
package sortext

import (
	"math/big"
	"sort"
)

// BigIntSlice attaches the methods of Interface to []*big.Int, sorting in increasing order.
type BigIntSlice []*big.Int

func (b BigIntSlice) Len() int           { return len(b) }
func (b BigIntSlice) Less(i, j int) bool { return b[i].Cmp(b[j]) < 0 }
func (b BigIntSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// Sort is a convenience method.
func (b BigIntSlice) Sort() { sort.Sort(b) }

// BigInts sorts a slice of *big.Ints in increasing order.
func BigInts(a []*big.Int) { sort.Sort(BigIntSlice(a)) }

// BigIntsAreSorted tests whether a slice of *big.Ints is sorted in increasing order.
func BigIntsAreSorted(a []*big.Int) bool { return sort.IsSorted(BigIntSlice(a)) }

// BigRatSlice attaches the methods of Interface to []*big.Rat, sorting in increasing order.
type BigRatSlice []*big.Rat

func (b BigRatSlice) Len() int           { return len(b) }
func (b BigRatSlice) Less(i, j int) bool { return b[i].Cmp(b[j]) < 0 }
func (b BigRatSlice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// Sort is a convenience method.
func (b BigRatSlice) Sort() { sort.Sort(b) }

// BigRats sorts a slice of *big.Rats in increasing order.
func BigRats(a []*big.Rat) { sort.Sort(BigRatSlice(a)) }

// BigRatsAreSorted tests whether a slice of *big.Rats is sorted in increasing order.
func BigRatsAreSorted(a []*big.Rat) bool { return sort.IsSorted(BigRatSlice(a)) }
