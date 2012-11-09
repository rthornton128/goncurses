/* This simple example mirrors the "hello world" TLDP ncurses howto */

package main

import "code.google.com/p/goncurses"

func main() {
	stdscr, _ := goncurses.Init()
	defer goncurses.End()

	stdscr.Print("Hello, World!!!")
	stdscr.Refresh()
	stdscr.GetChar()
}
