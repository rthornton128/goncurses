// ncurses_example.go
package goncurses

import (
	"code.google.com/p/goncurses"
	"os"
)

func ExampleInit() {
	goncurses.Init()
	defer goncurses.End()
}

func ExampleInit_withPrint() {
	// A full example can be found in examples/hello
	stdscr, err := goncurses.Init()
	defer goncurses.End()

	if err != nil {
		os.Exit(1)
	}

	stdscr.Print("Hello!")
	stdscr.GetChar()
}
