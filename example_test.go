// ncurses_example.go
package goncurses

import "code.google.com/p/goncurses"

func ExampleInit() {
	goncurses.Init()
	defer goncurses.End()
}
