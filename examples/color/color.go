// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* Demonstrates some of the color facilities of ncurses */
package main

import (
	gc "github.com/rthornton128/goncurses"
	"log"
)

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

	// HasColors can be used to determine whether the current terminal
	// has the capability of using colours. You could then chose whether or
	// not to use some other mechanism, like using A_REVERSE, instead
	if !gc.HasColors() {
		log.Fatal("Example requires a colour capable terminal")
	}

	// Must be called after Init but before using any colour related functions
	if err := gc.StartColor(); err != nil {
		log.Fatal(err)
	}

	gc.Echo(false)

	// Initialize a colour pair. Should only fail if an improper pair value
	// is given
	if err := gc.InitPair(1, gc.C_RED, gc.C_WHITE); err != nil {
		log.Fatal("InitPair failed: ", err)
	}
	gc.InitPair(2, gc.C_BLACK, gc.C_CYAN)

	stdscr.Println("Type any key to proceed and again to exit")

	// An example of trying to set an invalid color pair
	err = gc.InitPair(-1, gc.C_BLACK, gc.C_CYAN)
	stdscr.Println("An intentional error:", err)

	stdscr.Keypad(true)
	stdscr.MovePrint(12, 30, "Hello, World!!!")
	stdscr.Refresh()
	stdscr.GetChar()
	// Note that background doesn't just accept colours but will fill
	// any blank positions with the supplied character too. Note that newly
	// added text with spaces in it will have the blanks converted to the fill
	// character, if given
	stdscr.SetBackground(gc.Char('*') | gc.ColorPair(2))
	// ColorOn/Off is a shortcut to calling AttrOn/Off(gc.ColorPair(pair))
	stdscr.ColorOn(1)
	stdscr.MovePrint(13, 30, "Hello, World in Color!!!")
	stdscr.ColorOff(1)
	stdscr.Refresh()
	stdscr.GetChar()
}
