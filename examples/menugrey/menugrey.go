// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */
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

	gc.StartColor()
	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Keypad(true)
	gc.InitPair(1, gc.C_RED, gc.C_BLACK)
	gc.InitPair(2, gc.C_GREEN, gc.C_BLACK)
	gc.InitPair(3, gc.C_MAGENTA, gc.C_BLACK)

	// build the menu items
	menu_items := []string{
		"Choice 1",
		"Choice 2",
		"Choice 3",
		"Choice 4",
		"Choice 5",
		"Exit"}
	items := make([]*gc.MenuItem, len(menu_items))
	for i, val := range menu_items {
		items[i], _ = gc.NewItem(val, "")
		defer items[i].Free()

		if i == 2 || i == 4 {
			items[i].Selectable(false)
		}
	}

	// create the menu
	menu, _ := gc.NewMenu(items)
	defer menu.Free()

	y, _ := stdscr.MaxYX()
	stdscr.MovePrint(y-3, 0, "Use up/down arrows to move; 'q' to exit")
	stdscr.Refresh()

	menu.SetForeground(gc.ColorPair(1) | gc.A_REVERSE)
	menu.SetBackground(gc.ColorPair(2) | gc.A_BOLD)
	menu.Grey(gc.ColorPair(3) | gc.A_BOLD)

	menu.Post()
	defer menu.UnPost()

	for {
		gc.Update()
		ch := stdscr.GetChar()
		switch ch {
		case ' ':
			menu.Driver(gc.REQ_TOGGLE)
		case 'q':
			return
		case gc.KEY_RETURN:
			stdscr.Move(20, 0)
			stdscr.ClearToEOL()
			stdscr.Printf("Item selected is: %s", menu.Current(nil).Name())
			menu.PositionCursor()
		default:
			menu.Driver(gc.DriverActions[ch])
		}
	}
}
