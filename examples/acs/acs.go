// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* An example of using AddChar to show a non-standard character. Some common
 * VT100 symbols do not work on the windows command line (like ACS_DIAMOND) or
 * may require the use of chcp to change the codepage to 437 or 850 */
package main

import (
	gc "github.com/rthornton128/goncurses"
	"log"
)

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	gc.Cursor(0)
	gc.Echo(false)

	stdscr.Print("A reversed plus-minus symbol: ")
	stdscr.AddChar(gc.ACS_PLMINUS | gc.A_REVERSE)
	stdscr.Refresh()
	stdscr.GetChar()
}
