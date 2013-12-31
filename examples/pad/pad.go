// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A basic example of how to create and display a pad
package main

import gc "code.google.com/p/goncurses"

func main() {
	gc.Init()
	defer gc.End()

	// create a new pad of 30 rows and 100 columns and fill it
	pad := gc.NewPad(30, 100)
	for x := 1; x < 30; x++ {
		pad.MovePrint(x, x, "This is a pad.")
	}
	// show a 10x15 portion of the pad located at the coordinates
	// 5, 10 on the stdscr. The portion of the pad shown is a box
	// starting at 0, 5 to 10, 20
	pad.Refresh(0, 5, 5, 10, 15, 25)
	pad.GetChar()
}
