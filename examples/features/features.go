package main

import (
	gc "github.com/rthornton128/goncurses"
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
	var attributes = []struct {
		attr gc.Char
		text string
	}{
		{gc.A_NORMAL, "normal"},
		{gc.A_STANDOUT, "standout"},
		{gc.A_UNDERLINE | gc.A_BOLD, "underline"},
		{gc.A_REVERSE, "reverse"},
		{gc.A_BLINK, "blink"},
		{gc.A_DIM, "dim"},
		{gc.A_BOLD, "bold"},
	}
	stdscr.MovePrint(0, 0, "Normal terminal colors: ")
	for i, c := range colours {
		gc.InitPair(int16(i), c, c)
		stdscr.ColorOn(int16(i))
		stdscr.AddChar(' ')
		stdscr.ColorOff(int16(i))
	}
	stdscr.MovePrint(1, 0, "Bold terminal colors:   ")
	stdscr.AttrOn(gc.A_BLINK)
	for i, _ := range colours {
		stdscr.ColorOn(int16(i))
		stdscr.AddChar(' ')
		stdscr.ColorOff(int16(i))
	}
	stdscr.AttrOff(gc.A_BLINK)
	stdscr.Move(2, 0)
	for _, a := range attributes {
		stdscr.AttrOn(a.attr)
		stdscr.Println(a.text)
		stdscr.AttrOff(a.attr)
	}
	stdscr.GetChar()
}
