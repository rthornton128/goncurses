// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */
package main

import gc "github.com/rthornton128/goncurses"

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()

	gc.StartColor()
	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Keypad(true)

	// build the menu items
	menu_items := []string{
		"Choice 1",
		"Choice 2",
		"Choice 3",
		"Choice 4",
		"Choice 5",
		"Choice 6",
		"Choice 7",
		"Exit"}
	items := make([]*gc.MenuItem, len(menu_items))
	for i, val := range menu_items {
		items[i], _ = gc.NewItem(val, "")
		defer items[i].Free()
	}

	// create the menu
	menu, _ := gc.NewMenu(items)
	defer menu.Free()

	menu.Option(gc.O_ONEVALUE, false)

	y, _ := stdscr.MaxYX()
	stdscr.MovePrint(y-3, 0, "Use up/down arrows to move, spacebar to "+
		"toggle and enter to print. 'q' to exit")
	stdscr.Refresh()

	menu.Post()
	defer menu.UnPost()

	for {
		gc.Update()
		ch := stdscr.GetChar()

		switch ch {
		case 'q':
			return
		case ' ':
			menu.Driver(gc.REQ_TOGGLE)
		case gc.KEY_RETURN, gc.KEY_ENTER:
			var list string
			for _, item := range menu.Items() {
				if item.Value() {
					list += "\"" + item.Name() + "\" "
				}
			}
			stdscr.Move(20, 0)
			stdscr.ClearToEOL()
			stdscr.MovePrint(20, 0, list)
			stdscr.Refresh()
		default:
			menu.Driver(gc.DriverActions[ch])
		}
	}
}
