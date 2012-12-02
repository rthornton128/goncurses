/* This example mirrors the second example in the TLDP ncurses howto,
   demonstrating some of the initilization options for ncurses;
   In gnome, the F1 key launches help, so F2 is tested for instead */

package main

import "code.google.com/p/goncurses"

func main() {
	stdscr, _ := goncurses.Init()
	defer goncurses.End()

	goncurses.Raw(true)
	goncurses.Echo(false)
	stdscr.Keypad(true)

	stdscr.Print("Press a key...")
	stdscr.Refresh()

	if ch := stdscr.GetChar(); ch == goncurses.KEY_F2 {
		stdscr.Print("The F2 key was pressed.")
	} else {
		stdscr.Print("The key pressed is: ")
		stdscr.AttrOn(goncurses.A_BOLD)
		stdscr.AddChar(goncurses.AsciiCharacter(ch))
		stdscr.AttrOff(goncurses.A_BOLD)
	}
	stdscr.Refresh()
	stdscr.GetChar()
}
