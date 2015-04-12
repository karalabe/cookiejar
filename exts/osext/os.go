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

// Package osext contains extensions to the base Go os package.
package osext

import "os"

// MustOpen opens the named file for reading. If successful, methods on the
// returned file can be used for reading; the associated file descriptor has
// mode O_RDONLY. If there is an error, a panic will occur.
func MustOpen(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return file
}

// MustCreate creates the named file mode 0666 (before umask), truncating it if
// it already exists. If successful, methods on the returned File can be used
// for I/O; the associated file descriptor has mode O_RDWR. If there is an error,
// a panic will occur.
func MustCreate(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return file
}
