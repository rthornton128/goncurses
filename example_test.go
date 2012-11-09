// ncurses_example.go
package goncurses

import (
	"code.google.com/p/goncurses"
	"fmt"
	"os"
)

func ExampleEnd() {
	// You should always test to make sure ncurses has initialized properly.
	// In order for your error messages to be visible on the terminal you will
	// need to either log error messages or output them to to stderr.
	stdscr, err := goncurses.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer goncurses.End()
}

func ExampleInit() {
	// You should always test to make sure ncurses has initialized properly.
	// In order for your error messages to be visible on the terminal you will
	// need to either log error messages or output them to to stderr.
	stdscr, err := goncurses.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer goncurses.End()
}
