/* This simnple example mirrors one in the TLDP ncurses howto. It demonstrates
 * how to move a window around the screen. It is advisable not to  */

package main

import "code.google.com/p/goncurses"

func main() {
	stdscr, _ := goncurses.Init()
	defer goncurses.End()

	goncurses.Echo(false)
	goncurses.CBreak(true)
	goncurses.Cursor(0)
	stdscr.Keypad(true)
	stdscr.Print("Use arrow keys to move window. Press 'q' to exit")
	stdscr.Refresh()

	rows, cols := stdscr.Maxyx()
	height, width := 3, 10
	y, x := (rows-height)/2, (cols-width)/2
	win := createWin(height, width, y, x)

	for {
		switch stdscr.GetChar() {
		case 'q':
			return
		case goncurses.KEY_LEFT:
			x -= 1
		case goncurses.KEY_RIGHT:
			x += 1
		case goncurses.KEY_UP:
			y -= 1
		case goncurses.KEY_DOWN:
			y += 1
		}
		destroy(win)
		win = createWin(height, width, y, x)
	}
}

func createWin(h, w, y, x int) goncurses.Window {
	new, _ := goncurses.NewWindow(h, w, y, x)
	new.Box(0, 0)
	new.Refresh()
	return new
}

func destroy(w goncurses.Window) {
	w.Border(' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ')
	w.Refresh()
	w.Delete()
	return
}
