// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This example demonstrates the use of the menu library */
package main

import (
	"github.com/rthornton128/goncurses"
	"log"
)

const (
	HEIGHT = 10
	WIDTH  = 30
)

func main() {
	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer goncurses.End()

	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

	menu_items := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4",
		"Exit"}
	items := make([]*goncurses.MenuItem, len(menu_items))
	for i, val := range menu_items {
		items[i], _ = goncurses.NewItem(val, "")
		defer items[i].Free()
	}

	menu, err := goncurses.NewMenu(items)
	if err != nil {
		stdscr.Print(err)
		return
	}
	defer menu.Free()

	menu.Post()

	stdscr.MovePrint(20, 0, "'q' to exit")
	stdscr.Refresh()

	for {
		goncurses.Update()
		ch := stdscr.GetChar()

		switch goncurses.KeyString(ch) {
		case "q":
			return
		case "down":
			menu.Driver(goncurses.REQ_DOWN)
		case "up":
			menu.Driver(goncurses.REQ_UP)
		}
	}
}
