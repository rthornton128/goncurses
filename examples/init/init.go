/* This example mirrors the second example in the TLDP ncurses howto,
    demonstrating some of the initilization options for ncurses;
    In gnome, the F1 key launches help, so F2 is tested for instead */

package main

import . "goncurses"

func main() {
    stdscr, _ := Initscr();
    defer Endwin()
    
    Raw()
    Echo()
    stdscr.Keypad(true)
    
    stdscr.Refresh()
    ch, _ := stdscr.Getch()
    if key := Key(ch); key == "F2" {
        stdscr.Print("The F2 key was pressed.")
    } else {
        stdscr.Print("The key pressed is: ")
        stdscr.Attron("bold")
        stdscr.Addch(Chtype(ch))
        stdscr.Attroff("bold")
    }
    stdscr.Refresh()
    stdscr.Getch()
}
