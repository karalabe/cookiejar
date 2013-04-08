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
package prque_test

import (
	"fmt"
	"github.com/karalabe/cookiejar/prque"
)

// Insert some data into a priority queue and pop them out in prioritized order.
func Example_usage() {
	// Define some data to push into the priority queue
	prio := []float32{77.7, 22.2, 44.4, 55.5, 11.1, 88.8, 33.3, 99.9, 0.0, 66.6}
	data := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// Create the priority queue and insert the prioritized data
	pq := prque.New()
	for i := 0; i < len(data); i++ {
		pq.Push(data[i], prio[i])
	}
	// Pop out the data and print them
	for !pq.Empty() {
		val, prio := pq.Pop()
		fmt.Printf("%.1f:%s ", prio, val)
	}
	// Output:
	// 99.9:seven 88.8:five 77.7:zero 66.6:nine 55.5:three 44.4:two 33.3:six 22.2:one 11.1:four 0.0:eight
}
