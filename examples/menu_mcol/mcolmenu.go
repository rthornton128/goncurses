// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example shows a basic multi-column menu similar to that found in the
 * ncurses examples from TLDP */
package main

import gc "github.com/rthornton128/goncurses"

const (
	HEIGHT = 10
	WIDTH  = 40
)

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()

	gc.StartColor()
	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Keypad(true)
	gc.InitPair(1, gc.C_RED, gc.C_BLACK)
	gc.InitPair(2, gc.C_CYAN, gc.C_BLACK)

	// build the menu items
	menu_items := []string{
		"Choice 1",
		"Choice 2",
		"Choice 3",
		"Choice 4",
		"Choice 5",
		"Choice 6",
		"Choice 7",
		"Choice 8",
		"Choice 9",
		"Choice 10",
		"Choice 11",
		"Choice 12",
		"Choice 13",
		"Choice 14",
		"Choice 15",
		"Choice 16",
		"Choice 17",
		"Choice 18",
		"Choice 19",
		"Choice 20",
		"Exit"}
	items := make([]*gc.MenuItem, len(menu_items))
	for i, val := range menu_items {
		items[i], _ = gc.NewItem(val, "")
		defer items[i].Free()
	}

	// create the menu
	menu, _ := gc.NewMenu(items)
	defer menu.Free()

	menuwin, _ := gc.NewWindow(HEIGHT, WIDTH, 4, 14)
	menuwin.Keypad(true)

	menu.SetWindow(menuwin)
	dwin := menuwin.Derived(6, 38, 3, 1)
	menu.SubWindow(dwin)
	menu.Option(gc.O_SHOWDESC, true)
	menu.Format(5, 3)
	menu.Mark(" * ")

	// MovePrint centered menu title
	title := "My Menu"
	menuwin.Box(0, 0)
	menuwin.ColorOn(1)
	menuwin.MovePrint(1, (WIDTH/2)-(len(title)/2), title)
	menuwin.ColorOff(1)
	menuwin.MoveAddChar(2, 0, gc.ACS_LTEE)
	menuwin.HLine(2, 1, gc.ACS_HLINE, WIDTH-2)
	menuwin.MoveAddChar(2, WIDTH-1, gc.ACS_RTEE)

	y, _ := stdscr.MaxYX()
	stdscr.ColorOn(2)
	stdscr.MovePrint(y-3, 1,
		"Use up/down arrows or page up/down to navigate. 'q' to exit")
	stdscr.ColorOff(2)
	stdscr.Refresh()

	menu.Post()
	defer menu.UnPost()
	menuwin.Refresh()

	for {
		gc.Update()

		if ch := menuwin.GetChar(); ch == 'q' {
			return
		} else {
			menu.Driver(gc.DriverActions[ch])
		}
	}
}
