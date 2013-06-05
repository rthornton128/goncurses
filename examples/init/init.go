/* This example mirrors the second example in the TLDP ncurses howto,
   demonstrating some of the initilization options for ncurses;
   In gnome, the F1 key launches help, so F2 is tested for instead */

package main

import (
	"code.google.com/p/goncurses"
	"log"
)

func main() {
	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("init", err)
	}
	defer goncurses.End()

	goncurses.Raw(true)   // turn on raw "uncooked" input
	goncurses.Echo(false) // turn echoing of typed characters off
	stdscr.Keypad(true)   // allow keypad input

	stdscr.Print("Press a key...")
	stdscr.Refresh()

	if ch := stdscr.GetChar(); ch == goncurses.KEY_F2 {
		stdscr.Print("The F2 key was pressed.")
	} else {
		stdscr.Print("The key pressed is: ")
		stdscr.AttrOn(goncurses.A_BOLD)
		stdscr.AddChar(goncurses.Character(ch))
		stdscr.AttrOff(goncurses.A_BOLD)
	}
	stdscr.Refresh()
	stdscr.GetChar()
}
