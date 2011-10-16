/* This simnple example demonstrates some of the color facilities of ncurses */

package main

/* Note that is not considered idiomatic Go to import curses this way */
import . "goncurses.googlecode.com/hg/goncurses"

func main() {
    stdscr, _ := Init()
    defer End()
    StartColor()
    
    Raw(true)
    Echo(true)
    InitPair(1, C_BLUE, C_WHITE)
    InitPair(2, C_BLACK, C_CYAN)
    
    // An example of trying to set an invalid color pair
    err := InitPair(255, C_BLACK, C_CYAN)
    stdscr.Print(err.String())
    
    stdscr.Background(ColorPair(2))
    stdscr.Keypad(true)
    stdscr.Print(12, 30, "Hello, World!!!")
    stdscr.ColorOn(1)
    stdscr.Print(13, 30, "Hello, World in Color!!!")
    stdscr.ColorOff(1)
    stdscr.Refresh()
    stdscr.GetChar()
}
