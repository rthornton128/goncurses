/* This example demonstrates reading a string from input, rather than a 
 * single character. Note that only the 'n' versions of getstr have been
 * implemented in goncurses to ensure buffer overflows won't exist */

package main

import . "goncurses"

func main() {
    stdscr, _ := Initscr();
    defer Endwin()
    
    row, col := stdscr.Getmaxyx()
    msg := "Enter a string"
    stdscr.Mvprint(row/2, (col-len(msg))/2, msg)
    
    str, _ := stdscr.Getnstr(10)
    stdscr.Mvprint(row-2, 0, "You entered: %s", str)

    stdscr.Refresh()
    stdscr.Getch()
}
