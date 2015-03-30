// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* The classic "Hello, World!" program in Goncurses! */
package main

import (
	"github.com/rthornton128/goncurses"
	"log"
)

func main() {
	// Initialize goncurses. It's essential End() is called to ensure the
	// terminal isn't altered after the program ends
	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("init", err)
	}
	defer goncurses.End()

	// (Go)ncurses draws by cursor position and uses reverse cartisian
	// coordinate system (y,x). Initially, the cursor is positioned at the
	// coordinates 0,0 so the first call to Print will output the text at the
	// top, left position of the screen since stdscr is a window which
	// represents the terminal's screen size.
	stdscr.Print("Hello, World!")
	stdscr.MovePrint(3, 0, "Press any key to continue")

	// Refresh() flushes output to the screen. Internally, it is the same as
	// calling NoutRefresh() on the window followed by a call to Update()
	stdscr.Refresh()

	// GetChar will block execution until an input event, like typing a
	// character on your keyboard, is received
	stdscr.GetChar()
}
