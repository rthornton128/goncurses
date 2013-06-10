/* This simnple example demonstrates some of the color facilities of ncurses */

package main

/* Note that is not considered idiomatic Go to import curses this way */
import (
	. "code.google.com/p/goncurses"
	"log"
)

func main() {
	stdscr, err := Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer End()
	StartColor()

	Raw(true)
	Echo(false)
	InitPair(1, C_BLUE, C_WHITE)
	InitPair(2, C_BLACK, C_CYAN)

	stdscr.Println("Type 'q' to proceed and again to exit")

	// An example of trying to set an invalid color pair
	err = InitPair(255, C_BLACK, C_CYAN)
	stdscr.Println("An intentional error:", err)

	stdscr.Keypad(true)
	stdscr.MovePrint(12, 30, "Hello, World!!!")
	stdscr.Refresh()
	stdscr.GetChar()
	// Note that background doesn't just accept colours but will fill
	// any blank positions with the supplied character too
	stdscr.SetBackground(Character('-' | ColorPair(2)))
	stdscr.ColorOn(1)
	stdscr.MovePrint(13, 30, "Hello, World in Color!!!")
	stdscr.ColorOff(1)
	stdscr.Refresh()
	stdscr.GetChar()
}
