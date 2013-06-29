// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example demonstrates reading a string from input, rather than a 
 * single character. Note that only the 'n' versions of getstr have been
 * implemented in goncurses to ensure buffer overflows won't exist */

package main

import (
	gc "code.google.com/p/goncurses"
	"log"
)

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

	msg := "Enter a string: "
	row, col := stdscr.Maxyx()
	row, col = (row/2)-1, (col-len(msg))/2
	stdscr.MovePrint(row, col, msg)

	var str string
	str, err = stdscr.GetString(10)
	if err != nil {
		stdscr.MovePrint(row+1, col, "GetString Error:", err)
	} else {
		stdscr.MovePrintf(row+1, col, "You entered: %s", str)
	}

	stdscr.Refresh()
	stdscr.GetChar()
}
