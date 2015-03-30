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
	gc.InitPair(1, gc.C_RED, gc.C_BLACK)

	// build the menu items
	menu_items := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4",
		"Exit"}
	items := make([]*gc.MenuItem, len(menu_items))
	for i, val := range menu_items {
		items[i], _ = gc.NewItem(val, "")
		defer items[i].Free()
	}

	// create the menu
	menu, _ := gc.NewMenu(items)
	defer menu.Free()

	menuwin, _ := gc.NewWindow(10, 40, 4, 14)
	menuwin.Keypad(true)

	menu.SetWindow(menuwin)
	dwin := menuwin.Derived(6, 38, 3, 1)
	menu.SubWindow(dwin)
	menu.Mark(" * ")

	// Print centered menu title
	y, x := menuwin.MaxYX()
	title := "My Menu"
	menuwin.Box(0, 0)
	menuwin.ColorOn(1)
	menuwin.MovePrint(1, (x/2)-(len(title)/2), title)
	menuwin.ColorOff(1)
	menuwin.MoveAddChar(2, 0, gc.ACS_LTEE)
	menuwin.HLine(2, 1, gc.ACS_HLINE, x-3)
	menuwin.MoveAddChar(2, x-2, gc.ACS_RTEE)

	y, x = stdscr.MaxYX()
	stdscr.MovePrint(y-2, 1, "'q' to exit")
	stdscr.Refresh()

	menu.Post()
	defer menu.UnPost()
	menuwin.Refresh()

	for {
		gc.Update()
		ch := menuwin.GetChar()

		switch ch {
		case 'q':
			return
		case gc.KEY_DOWN:
			menu.Driver(gc.REQ_DOWN)
		case gc.KEY_UP:
			menu.Driver(gc.REQ_UP)
		}
	}
}
