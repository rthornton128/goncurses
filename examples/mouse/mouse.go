// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* Expanding on the basic menu example, the example demonstrates how you
 * could possibly utilize the mouse to navigate a menu and select options
 */
package main

import (
	"log"

	gc "github.com/rthornton128/goncurses"
)

const (
	HEIGHT = 10
	WIDTH  = 30
)

func main() {
	var active int
	menu := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4", "Exit"}

	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Keypad(true)

	y, x := 5, 2

	win, err := gc.NewWindow(HEIGHT, WIDTH, y, x)
	if err != nil {
		log.Fatal("new_window:", err)
	}
	stdscr.MovePrint(0, 0, "Use up/down arrow keys, enter to "+
		"select or 'q' to quit")
	stdscr.MovePrint(1, 0, "You may also left mouse click on an entry to select")
	stdscr.Refresh()

	printmenu(win, menu, active)

	// Check to see if the Mouse is available. This function was not available
	// in older versions of ncurses (5.7 and older) and may return false even
	// if the mouse is in fact avaiable for use.
	if gc.MouseOk() {
		stdscr.MovePrint(3, 0, "WARN: Mouse support not detected.")
	}

	// Adjust the default mouse-click sensitivity to make it more responsive
	gc.MouseInterval(50)

	// If, for example, you are temporarily disabling the mouse or are
	// otherwise altering mouse button detection temporarily, you could
	// pass a pointer to a MouseButton object as the 2nd argument to
	// record that information. Invocation may look something like:

	var prev gc.MouseButton
	gc.MouseMask(gc.M_B1_PRESSED, nil) // only detect left mouse clicks
	gc.MouseMask(gc.M_ALL, &prev)      // temporarily enable all mouse clicks
	gc.MouseMask(prev, nil)            // change it back

	var key gc.Key
	for key != 'q' {
		key = stdscr.GetChar()
		switch key {
		case gc.KEY_UP:
			if active == 0 {
				active = len(menu) - 1
			} else {
				active -= 1
			}
		case gc.KEY_DOWN:
			if active == len(menu)-1 {
				active = 0
			} else {
				active += 1
			}
		case gc.KEY_MOUSE:
			/* pull the mouse event off the queue */
			if md := gc.GetMouse(); md != nil {
				new := getactive(x, y, md.X, md.Y, menu)
				if new != -1 {
					active = new
				}
			}
			fallthrough
		case gc.KEY_RETURN, gc.KEY_ENTER, gc.Key('\r'):
			stdscr.MovePrintf(23, 0, "Choice #%d: %s selected", active+1,
				menu[active])
			stdscr.ClearToEOL()
			stdscr.Refresh()
		default:
			stdscr.MovePrintf(23, 0, "Character pressed = %3d/%c", key, key)
			stdscr.ClearToEOL()
			stdscr.Refresh()
		}

		printmenu(win, menu, active)
	}
}

func getactive(x, y, mx, my int, menu []string) int {
	row := my - y - 2
	col := mx - x - 2

	if row < 0 || row > len(menu)-1 {
		return -1
	}

	l := menu[row]

	if col >= 0 && col < len(l) {
		return row
	}
	return -1
}

func printmenu(w *gc.Window, menu []string, active int) {
	y, x := 2, 2
	w.Box(0, 0)
	for i, s := range menu {
		if i == active {
			w.AttrOn(gc.A_REVERSE)
			w.MovePrint(y+i, x, s)
			w.AttrOff(gc.A_REVERSE)
		} else {
			w.MovePrint(y+i, x, s)
		}
	}
	w.Refresh()
}
