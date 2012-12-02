/* This example demonstrates reading a string from input, rather than a 
 * single character. Note that only the 'n' versions of getstr have been
 * implemented in goncurses to ensure buffer overflows won't exist */

package main

import gc "code.google.com/p/goncurses"

func main() {
	stdscr, _ := gc.Init()
	defer gc.End()

	row, col := stdscr.Maxyx()
	msg := "Enter a string: "
	stdscr.Print(row/2, (col-len(msg)-8)/2, msg)

	str, _ := stdscr.GetString(10)
	stdscr.Print(row-2, 0, "You entered: %s", str)

	stdscr.Refresh()
	stdscr.GetChar()
}
