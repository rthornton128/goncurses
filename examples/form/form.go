/* This simnple example mirrors the "hello world" TLDP ncurses howto */

package main

/* Note that is not considered idiomatic Go to import curses this way */
import . "goncurses"//".googlecode.com/hg/goncurses"

func main() {
    stdscr, _ := Init()
    defer End()
    
    _, err := NewForm(nil)
    
    stdscr.Print("Error %s", err)
    stdscr.Refresh()
    stdscr.GetChar()
}
