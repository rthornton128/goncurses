/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */

package main

import (
	gc "code.google.com/p/goncurses"
	"log"
)

const (
	HEIGHT = 10
	WIDTH  = 30
)

func main() {
	var active int
	menu := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4", "Exit"}

	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()

	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

	rows, cols := stdscr.Maxyx()
	y, x := (rows-HEIGHT)/2, (cols-WIDTH)/2

	win, err := gc.NewWindow(HEIGHT, WIDTH, y, x)
	if err != nil {
		log.Fatal("new_window:", err)
	}
	win.Keypad(true)
	stdscr.MovePrint(0, 0,
		"Use arrow keys to go up and down, Press enter to select")
	stdscr.Refresh()

	printmenu(win, menu, active)
	if gc.Mouse() {
		stdscr.MovePrint(3, 0, "WARN: Mouse support not detected.")
	}
	// If, for example, you are temporarily disabling the mouse or are
	// otherwise altering mouse button detection temporarily, you could
	// pass a pointer to a MouseButton object as the 2nd argument to
	// record that information. Invocation may look something like:
	// var old gc.MouseButton
	// gc.MouseMask(gc.M_ALL, &old) /* temporarily enable all mouse clicks */
	// gc.MouseMask(old, nil)		/* change it back */
	gc.MouseMask(gc.M_B1_PRESSED, nil)

	for {
		ch := stdscr.GetChar()
		switch ch {
		case 'q':
			return
		case gc.KEY_UP:
			if active == 0 {
				active = len(menu) - 1
			} else {
				active -= 1
			}
		case gc.KEY_DOWN:
			if active == len(menu)-1 {
				active = 0
			} else {
				active += 1
			}
		case gc.KEY_MOUSE:
			md, err := gc.GetMouse()
			if err != nil {
				stdscr.MovePrint(20, 0, "%s", err)
			}
			new := getactive(x, y, md[0], md[1], menu)
			if new != -1 {
				active = new
			}
			stdscr.MovePrintf(23, 0, "Choice #%d: %s selected", active+1,
				menu[active])
			stdscr.ClearToEOL()
			stdscr.Refresh()
		case gc.KEY_RETURN:
			stdscr.MovePrintf(23, 0, "Choice #%d: %s selected", active+1,
				menu[active])
			stdscr.ClearToEOL()
			stdscr.Refresh()
		default:
			stdscr.MovePrintf(23, 0, "Character pressed = %3d/%c", ch, ch)
			stdscr.ClearToEOL()
			stdscr.Refresh()
		}

		printmenu(win, menu, active)
	}
}

func getactive(x, y, mx, my int, menu []string) int {
	row := my - y - 2
	col := mx - x - 2

	if row < 0 || row > len(menu)-1 {
		return -1
	}

	l := menu[row]

	if col >= 0 && col < len(l) {
		return row
	}
	return -1
}

func printmenu(w gc.Window, menu []string, active int) {
	y, x := 2, 2
	w.Box(0, 0)
	for i, s := range menu {
		if i == active {
			w.AttrOn(gc.A_REVERSE)
			w.MovePrint(y+i, x, s)
			w.AttrOff(gc.A_REVERSE)
		} else {
			w.MovePrint(y+i, x, s)
		}
	}
	w.Refresh()
}
