/* This simnple example mirrors the "hello world" TLDP ncurses howto */

package main

/* Note that is not considered idiomatic Go to import curses this way */
import . "goncurses"

func main() {
    stdscr, _ := Initscr()
    defer Endwin()
    
    stdscr.Print("Hello, World!!!")
    stdscr.Refresh()
    stdscr.Getch()
}
