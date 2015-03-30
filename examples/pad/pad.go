// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A basic example of how to create and display a pad. Unlike a window, a
// pad can exceed the size of the physical screen. In order to display a
// pad you must select the portion you wish to view and where it should be
// located on the screen.
package main

import (
	gc "github.com/rthornton128/goncurses"
	"log"
)

func main() {
	_, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	// create a new pad of 50 rows and 200 columns...
	var pad *gc.Pad
	pad, err = gc.NewPad(50, 200)
	if err != nil {
		log.Fatal(err)
	}
	// ...and fill it with some characters
	for x := 0; x < 50; x++ {
		pad.MovePrint(x, x, "This is a pad.")
	}
	// Refresh the pad to show only a portion of the pad. Understanding
	// what these coordinates mean can be a bit tricky. The first two
	// coordinates are the position in the pad, in this case 0,5 (remember
	// the coordinates in ncurses are y,x). The second set of numbers are the
	// coordinates on the screen on which to display the content, so row 5,
	// column 10. The last set of numbers tell how high and how wide the
	// rectangle to displayed should be, in this case 15 rows long and 25
	// columns wide.
	pad.Refresh(0, 5, 5, 10, 15, 25)
	pad.GetChar()
}
