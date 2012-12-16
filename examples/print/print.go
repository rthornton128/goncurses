/* This example demonstrates the use of the print function. */

package main

import gc "code.google.com/p/goncurses"

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()

	row, col := stdscr.Maxyx()
	msg := "Just a string "
	stdscr.MovePrint(row/2, (col-len(msg))/2, msg)

	stdscr.MovePrint(row-3, 0, "This screen has %d rows and %d columns. ",
		row, col)
	stdscr.MovePrint(row-2, 0, "Try resizing your window and then run this "+
		"program again.")

	stdscr.Refresh()
	stdscr.GetChar()
}
