/* An example of using AddChar to show a non-standard character */

package main

import . "goncurses.googlecode.com/hg/goncurses"

func main() {
    stdscr, _ := Init()
    defer End()
    
    stdscr.Print("A reversed color diamond: ")
    stdscr.AddChar(ACS_DIAMOND | A_REVERSE)
    stdscr.Refresh()
    stdscr.GetChar()
}
