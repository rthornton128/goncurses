// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP. Errors are intentionally discarded */

package main

import . "code.google.com/p/goncurses"

const (
	HEIGHT = 10
	WIDTH  = 30
)

func main() {
	var active int
	menu := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4", "Exit"}

	stdscr, _ := Init()
	defer End()

	Raw(true)
	Echo(false)
	Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

	rows, cols := stdscr.MaxYX()
	y, x := (rows-HEIGHT)/2, (cols-WIDTH)/2

	win, _ := NewWindow(HEIGHT, WIDTH, y, x)
	win.Keypad(true)

	stdscr.Print("Use arrow keys to go up and down, Press enter to select")
	stdscr.Refresh()

	printmenu(&win, menu, active)

	for {
		ch := stdscr.GetChar()
		switch KeyString(ch) {
		case "q":
			return
		case "up":
			if active == 0 {
				active = len(menu) - 1
			} else {
				active -= 1
			}
		case "down":
			if active == len(menu)-1 {
				active = 0
			} else {
				active += 1
			}
		case "enter":
			stdscr.MovePrintf(23, 0, "Choice #%d: %s selected",
				active,
				menu[active])
			stdscr.ClearToEOL()
			stdscr.Refresh()
		default:
			stdscr.MovePrintf(23, 0, "Character pressed = %3d/%c",
				ch, ch)
			stdscr.ClearToEOL()
			stdscr.Refresh()
		}

		printmenu(&win, menu, active)
	}
}

func printmenu(w *Window, menu []string, active int) {
	y, x := 2, 2
	w.Box(0, 0)
	for i, s := range menu {
		if i == active {
			w.AttrOn(A_REVERSE)
			w.MovePrint(y+i, x, s)
			w.AttrOff(A_REVERSE)
		} else {
			w.MovePrint(y+i, x, s)
		}
	}
	w.Refresh()
}
