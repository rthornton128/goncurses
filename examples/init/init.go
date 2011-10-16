/* This example mirrors the second example in the TLDP ncurses howto,
    demonstrating some of the initilization options for ncurses;
    In gnome, the F1 key launches help, so F2 is tested for instead */

package main

import . "goncurses.googlecode.com/hg/goncurses"

func main() {
    stdscr, _ := Init()
    defer End()
    
    Raw(true)
    Echo(false)
    stdscr.Keypad(true)
    
    stdscr.Print("Press a key...")
    stdscr.Refresh()
    
    ch, _ := stdscr.GetChar()
    if key := Key(ch); key == "F2" {
        stdscr.Print("The F2 key was pressed.")
    } else {
        stdscr.Print("The key pressed is: ")
        stdscr.AttrOn(A_BOLD)
        stdscr.AddChar(ch)
        stdscr.AttrOff(A_BOLD)
    }
    stdscr.Refresh()
    stdscr.GetChar()
}
