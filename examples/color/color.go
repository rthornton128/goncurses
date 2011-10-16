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
    InitPair(1, "blue", "white")
    InitPair(2, "black", "cyan")
    
    stdscr.Background(ColorPair(2))
    stdscr.Keypad(true)
    stdscr.Print(12, 30, "Hello, World!!!")
    stdscr.ColorOn(1)
    stdscr.Print(13, 30, "Hello, World in Color!!!")
    stdscr.ColorOff(1)
    stdscr.Refresh()
    stdscr.GetChar()
}
