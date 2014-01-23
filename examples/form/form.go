// goncurses - ncurses library for Go.
// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* This simple example demonstrates how to implement a form */
package main

import gc "code.google.com/p/goncurses"

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()

	gc.Echo(false)
	gc.CBreak(true)
	gc.StartColor()
	stdscr.Keypad(true)

	gc.InitPair(1, gc.C_WHITE, gc.C_BLUE)
	gc.InitPair(2, gc.C_YELLOW, gc.C_BLUE)

	fields := make([]*gc.Field, 2)
	fields[0], _ = gc.NewField(1, 10, 4, 18, 0, 0)
	defer fields[0].Free()
	fields[0].SetForeground(gc.ColorPair(1))
	fields[0].SetBackground(gc.ColorPair(2) | gc.A_UNDERLINE | gc.A_BOLD)
	fields[0].SetOptionsOff(gc.FO_AUTOSKIP)

	fields[1], _ = gc.NewField(1, 10, 6, 18, 0, 0)
	defer fields[1].Free()
	fields[1].SetForeground(gc.ColorPair(1))
	fields[1].SetBackground(gc.A_UNDERLINE)
	fields[1].SetOptionsOff(gc.FO_AUTOSKIP)
	fields[1].SetPad('*')

	form, _ := gc.NewForm(fields)
	form.Post()
	defer form.UnPost()
	defer form.Free()
	stdscr.Refresh()

	stdscr.AttrOn(gc.ColorPair(2) | gc.A_BOLD)
	stdscr.MovePrint(4, 10, "Value 1:")
	stdscr.AttrOff(gc.ColorPair(2) | gc.A_BOLD)
	stdscr.MovePrint(6, 10, "Value 2:")
	stdscr.Refresh()

	ch := stdscr.GetChar()
	for ch != 'q' {
		switch ch {
		case gc.KEY_DOWN:
			form.Driver(gc.REQ_NEXT_FIELD)
			form.Driver(gc.REQ_END_LINE)
		case gc.KEY_UP:
			form.Driver(gc.REQ_PREV_FIELD)
			form.Driver(gc.REQ_END_LINE)
		default:
			form.Driver(ch)
		}
		ch = stdscr.GetChar()
	}
}
