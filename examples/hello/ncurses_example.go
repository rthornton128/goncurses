// ncurses_example.go
package goncurses

import "code.google.com/p/goncurses"

func InitExample() {
	stdscr, _ := goncurses.Init()
	defer goncurses.End()
}
