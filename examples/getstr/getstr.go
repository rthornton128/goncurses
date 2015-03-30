// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example demonstrates reading a string from input, rather than a
 * single character */
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

	msg := "Enter a string: "
	row, col := stdscr.MaxYX()
	row, col = (row/2)-1, (col-len(msg))/2
	stdscr.MovePrint(row, col, msg)

	/* GetString will only retieve the specified number of characters. Any
	attempts by the user to enter more characters will elicit an audiable
	beep */
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
