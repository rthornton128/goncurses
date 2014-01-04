package main

import (
	gc "code.google.com/p/goncurses"

	"log"
)

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	gc.StartColor()
	gc.Echo(false)
	gc.Raw(true)

	var colours = []int16{gc.C_BLACK, gc.C_BLUE, gc.C_CYAN, gc.C_GREEN,
		gc.C_MAGENTA, gc.C_RED, gc.C_WHITE, gc.C_YELLOW}
	var attrs = []gc.Char{gc.A_NORMAL, gc.A_STANDOUT, gc.A_UNDERLINE,
		gc.A_REVERSE, gc.A_BLINK, gc.A_DIM, gc.A_BOLD}
	for j, a := range attrs {
		stdscr.Move(j, 0)
		stdscr.AttrOn(a)
		for i, c := range colours {
			gc.InitPair(int16(i), gc.C_WHITE, c)
			stdscr.ColorOn(int16(i))
			stdscr.AddChar('+')
			stdscr.ColorOff(int16(i))
		}
		stdscr.AttrOff(a)
	}

	for i := 0; i < 16; i++ {
		for j := 0; j < 8; j++ {
			x := 127 + ((i + 1) * (j + 1))
			stdscr.MovePrintf(i+10, j*6, "%x %c", x, x)
		}
	}
	stdscr.GetChar()
}
