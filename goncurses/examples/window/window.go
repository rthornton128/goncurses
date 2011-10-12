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
    
    ch, _ := stdscr.GetChar()
    
    for {
        switch(Key(ch)) {
        case "q":
            return
        case "left":
            x -= 1
        case "right":
            x += 1
        case "up":
            y -= 1
        case "down":
            y += 1
        }
        destroy(win)
        win = createWin(height, width, y, x)
        
        ch, _ = stdscr.GetChar()
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
