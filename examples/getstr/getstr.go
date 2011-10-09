/* This example demonstrates reading a string from input, rather than a 
 * single character. Note that only the 'n' versions of getstr have been
 * implemented in goncurses to ensure buffer overflows won't exist */

package main

import . "goncurses"

func main() {
    stdscr, _ := Init();
    defer End()
    
    row, col := stdscr.Maxyx()
    msg := "Enter a string: "
    stdscr.Print(row/2, (col-len(msg))/2, msg)
    
    str, _ := stdscr.GetString(10)
    stdscr.Print(row-2, 0, "You entered: %s", str)

    stdscr.Refresh()
    stdscr.GetChar()
}
