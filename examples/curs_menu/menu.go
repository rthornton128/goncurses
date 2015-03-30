// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */
package main

import (
	. "github.com/rthornton128/goncurses"
	"log"
)

const (
	HEIGHT = 10
	WIDTH  = 30
)

func main() {
	var active int
	menu := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4", "Exit"}

	stdscr, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	defer End()

	Raw(true)
	Echo(false)
	Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

	my, mx := stdscr.MaxYX()
	y, x := 2, (mx/2)-(WIDTH/2)

	win, _ := NewWindow(HEIGHT, WIDTH, y, x)
	win.Keypad(true)

	stdscr.Print("Use arrow keys to go up and down, Press enter to select")
	stdscr.Refresh()

	printmenu(win, menu, active)

	for {
		ch := stdscr.GetChar()
		switch Key(ch) {
		case 'q':
			return
		case KEY_UP:
			if active == 0 {
				active = len(menu) - 1
			} else {
				active -= 1
			}
		case KEY_DOWN:
			if active == len(menu)-1 {
				active = 0
			} else {
				active += 1
			}
		case KEY_RETURN, KEY_ENTER, Key('\r'):
			stdscr.MovePrintf(my-2, 0, "Choice #%d: %s selected",
				active,
				menu[active])
			stdscr.ClearToEOL()
			stdscr.Refresh()
		default:
			stdscr.MovePrintf(my-2, 0, "Character pressed = %3d/%c",
				ch, ch)
			stdscr.ClearToEOL()
			stdscr.Refresh()
		}

		printmenu(win, menu, active)
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
