// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example demonstrates the use of the print function. */
package main

import gc "github.com/rthornton128/goncurses"

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()

	row, col := stdscr.MaxYX()
	msg := "Just a string "
	stdscr.MovePrint(row/2, (col-len(msg))/2, msg)

	stdscr.MovePrintf(row-3, 0, "This screen has %d rows and %d columns. ",
		row, col)
	stdscr.Println()
	stdscr.Print("Try resizing your terminal window and then " +
		"run this program again.")
	stdscr.Refresh()
	stdscr.GetChar()
}
