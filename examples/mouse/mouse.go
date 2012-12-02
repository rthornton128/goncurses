/* This example show a basic menu similar to that found in the ncurses
 * examples from TLDP */

package main

import gc "code.google.com/p/goncurses"

const (
	HEIGHT = 10
	WIDTH  = 30
)

func main() {
	var active int
	menu := []string{"Choice 1", "Choice 2", "Choice 3", "Choice 4", "Exit"}

	stdscr, _ := gc.Init()
	defer gc.End()

	gc.Raw(true)
	gc.Echo(false)
	gc.Cursor(0)
	stdscr.Clear()
	stdscr.Keypad(true)

	rows, cols := stdscr.Maxyx()
	y, x := (rows-HEIGHT)/2, (cols-WIDTH)/2

	win, _ := gc.NewWindow(HEIGHT, WIDTH, y, x)
	win.Keypad(true)
	stdscr.Print(0, 0,
		"Use arrow keys to go up and down, Press enter to select")
	stdscr.Refresh()

	printmenu(win, menu, active)
	gc.MouseMask(gc.M_B1_CLICKED, nil)

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
			md, _ := gc.GetMouse()
			new := getactive(x, y, md[0], md[1], menu)
			if new != -1 {
				active = new
			}
			stdscr.Print(23, 0, "Choice #%d: %s selected", active+1,
				menu[active])
			stdscr.ClearToEOL()
			stdscr.Refresh()
		case gc.KEY_RETURN:
			stdscr.Print(23, 0, "Choice #%d: %s selected", active+1,
				menu[active])
			stdscr.ClearToEOL()
			stdscr.Refresh()
		default:
			stdscr.Print(23, 0, "Character pressed = %3d/%c", ch, ch)
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
			w.Print(y+i, x, s)
			w.AttrOff(gc.A_REVERSE)
		} else {
			w.Print(y+i, x, s)
		}
	}
	w.Refresh()
}
