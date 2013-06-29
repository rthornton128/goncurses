// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This simple example mirrors the "hello world" TLDP ncurses howto */

package main

import (
	"code.google.com/p/goncurses"
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

	stdscr.Print("Hello, World!!!")
	stdscr.MovePrint(3, 0, "Press any key to continue")
	stdscr.Refresh()
	stdscr.GetChar()
}
