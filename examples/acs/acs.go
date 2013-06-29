// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* An example of using AddChar to show a non-standard character */

package main

import . "code.google.com/p/goncurses"

func main() {
	stdscr, _ := Init()
	defer End()

	stdscr.Print("A reversed color diamond: ")
	stdscr.AddChar(ACS_DIAMOND | A_REVERSE)
	stdscr.Refresh()
	stdscr.GetChar()
}
