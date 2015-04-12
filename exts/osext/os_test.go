// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2015 Peter Szilagyi. All rights reserved.
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

package osext

import (
	"os"
	"testing"
)

var testFile = "test.txt"

func TestOpen(t *testing.T) {
	// Create a file and make sure it's removed after the test
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("failed to create test file: %v.", err)
	}
	file.Close()
	defer os.Remove(testFile)

	// Try and read the file
	file = MustOpen(testFile)
	file.Close()
}

func TestCreate(t *testing.T) {
	// Check that a test file is non-existent
	if stats, err := os.Stat(testFile); err == nil {
		t.Errorf("file already exists: %v.", stats)
	}
	// Create an empty file and make sure it's dumped after the test
	file := MustCreate(testFile)
	file.Close()
	defer os.Remove(testFile)

	// Verify that the file has been created
	if stats, err := os.Stat(testFile); err != nil {
		t.Errorf("file doesn't exist: %v.", stats)
	}
}
