/* This simnple example mirrors one in the TLDP ncurses howto. It demonstrates
 * how to move a window around the screen. It is advisable not to  */

package main

import . "goncurses.googlecode.com/hg/goncurses"

func main() {    
    stdscr, _ := Init()
    defer End()
    
    Echo(false)
    CBreak(true)
    Cursor(0)
    stdscr.Keypad(true)
    stdscr.Print("Press 'q' to exit")
    stdscr.Refresh()
    
    rows, cols := stdscr.Maxyx()
    height, width := 3, 10
    y, x := (rows-height)/2, (cols-width)/2
    win, _ := NewWindow(height, width, y, x)
    win.Refresh()
    
    for {
        switch stdscr.GetChar() {
        case 'q':
            return
        case KEY_LEFT:
            x -= 1
        case KEY_RIGHT:
            x += 1
        case KEY_UP:
            y -= 1
        case KEY_DOWN:
            y += 1
        }
        destroy(win)
        win = createWin(height, width, y, x)
    }
}

func createWin(h, w, y, x int) *Window {
    new, _ := NewWindow(h, w, y, x)
    new.Box(0, 0)
    new.Refresh()
    return new
}

func destroy(w *Window) {
    w.Border(' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ')
    w.Refresh()
    w.Delete()
    return
}
