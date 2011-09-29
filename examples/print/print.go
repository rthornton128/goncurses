/* This example mirrors the second example in the TLDP ncurses howto,
    demonstrating some of the initilization options for ncurses;
    In gnome, the F1 key launches help, so F2 is tested for instead */

package main

import . "goncurses"

func main() {
    stdscr, _ := Initscr();
    defer Endwin()
    
    row, col := stdscr.Getmaxyx()
    msg := "Just a string"
    stdscr.Mvprint(row/2, (col-len(msg))/2, msg)
    
    stdscr.Mvprint(row-2, 0, "This screen has %d rows and %d columns. ",
        row, col)
    stdscr.Print("Try resizing your window and then run this program again.")

    stdscr.Refresh()
    stdscr.Getch()
}
