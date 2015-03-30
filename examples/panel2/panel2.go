// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* A slightly more advanced example of how to use the panel routines */
package main

import gc "github.com/rthornton128/goncurses"

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()

	gc.StartColor()
	gc.CBreak(true)
	gc.Echo(true)
	stdscr.Keypad(true)
	stdscr.Print("Hit 'tab' to cycle through windows, 'q' to quit")

	gc.InitPair(1, gc.C_RED, gc.C_BLACK)
	gc.InitPair(2, gc.C_GREEN, gc.C_BLACK)
	gc.InitPair(3, gc.C_BLUE, gc.C_BLACK)
	gc.InitPair(4, gc.C_CYAN, gc.C_BLACK)

	var panels [3]*gc.Panel
	y, x := 4, 10

	for i := 0; i < 3; i++ {
		h, w := 10, 40
		title := "Window Number %d"

		window, _ := gc.NewWindow(h, w, y+(i*4), x+(i*10))
		window.Box(0, 0)
		window.MoveAddChar(2, 0, gc.ACS_LTEE)
		window.HLine(2, 1, gc.ACS_HLINE, w-2)
		window.MoveAddChar(2, w-1, gc.ACS_RTEE)
		window.ColorOn(int16(i + 1))
		window.MovePrintf(1, (w/2)-(len(title)/2), title, i+1)
		window.ColorOff(int16(i + 1))
		panels[i] = gc.NewPanel(window)

	}

	active := 2

	for {
		gc.UpdatePanels()
		gc.Update()

		switch stdscr.GetChar() {
		case 'q':
			return
		case gc.KEY_TAB:
			active += 1
			if active > 2 {
				active = 0
			}
			panels[active].Top()
		}
	}
}
